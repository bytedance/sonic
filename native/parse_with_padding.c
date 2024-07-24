#include "native.h"
#include "simd.h"
#include <stdint.h>

#include "parsing.h"
#include "scanning.h"
#include "test/xassert.h"

// Option and Error -----------------------------------------------------------

typedef int32_t     error_code;

static const error_code SONIC_OK                = 0;
static const error_code SONIC_CONTROL_CHAR      = 1;
static const error_code SONIC_INVALID_ESCAPED   = 2;
static const error_code SONIC_INVALID_NUM       = 3;
static const error_code SONIC_FLOAT_INF         = 4;
static const error_code SONIC_EOF               = 5;
static const error_code SONIC_INVALID_CHAR      = 6;
static const error_code SONIC_EXPECT_KEY        = 7;
static const error_code SONIC_EXPECT_COLON      = 8;
static const error_code SONIC_EXPECT_OBJ_COMMA_OR_END  = 9;
static const error_code SONIC_EXPECT_ARR_COMMA_OR_END  = 10;
static const error_code SONIC_VISIT_FAILED        = 11;
static const error_code SONIC_INVALID_ESCAPED_UTF = 12;
static const error_code SONIC_INVALID_LITERAL     = 13;
static const error_code SONIC_STACK_OVERFLOW      = 14;

typedef uint32_t    option;

static const option _F_USE_INT64 = 0;
static const option _F_DIS_URC = 2;
static const option _F_DIS_UNKNOWN = 3;
static const option _F_COPY_STR = 4;

static const option _F_USE_NUMBER = 1; // types.B_USE_NUMBER
static const option _F_VALIDATE_STR = 5; // types.B_VALIDATE_STRING
static const option _F_ALLOW_CTL = 31; // types.B_ALLOW_CONTROL

// NOTE: only enbale these flags in sonic-rs.
static const option F_USE_NUMBER        = 1 << _F_USE_NUMBER;
static const option F_USE_INT64         = 1 << _F_USE_INT64;
static const option F_VALIDATE_STRING   = 1 << _F_VALIDATE_STR;

static const uint32_t MAX_DEPTTH = 4096;

// Macro Helper ----------------------------------------------------------------

#define repeat_2(expr)  {{ expr; } { expr; }}
#define repeat_4(expr)  { repeat_2(expr) repeat_2(expr) }
#define repeat_8(expr)  { repeat_4(expr) repeat_4(expr) }
#define repeat_16(expr) { repeat_8(expr) repeat_8(expr) }

#define trailing_zeros(bits)        (__builtin_ctzll(bits))
#define clear_low_bits(bits, n)     ((bits) & (~((1ull << n) - 1)))
#define is_zero(bits)               ((bits) == 0)
#define at_first(first, second)     ((((second) - 1) & (first)) != 0)

// JSON Reader ----------------------------------------------------------------

#define SONIC_NONE '\0'

typedef struct {
    char        (*next)(void* ctx);
    uint8_t*    (*peek_n)(void* ctx, size_t n);
    uint8_t**   (*cur)(void* ctx);
    ssize_t     (*remain)(void* ctx);
    void        (*eat)(void* ctx, size_t n);
    void*   ctx;
} reader;

// Read the padded JSON
typedef struct {
    uint8_t* cur;
    const uint8_t* start;
    const uint8_t* end;
} padding_reader;


// the next will always ok when parsing because of the padding chars
char padding_reader_next(void* ctx) {
    padding_reader* rdr = (padding_reader*)ctx;
    return *rdr->cur++;
}


uint8_t* padding_reader_peek_n(void* ctx, size_t n) {
    padding_reader* rdr = (padding_reader*)ctx;
    return rdr->cur;
}

uint8_t** padding_reader_cur(void* ctx) {
    padding_reader* rdr = (padding_reader*)ctx;
    return (uint8_t**)&rdr->cur;
}

ssize_t padding_reader_remain(void* ctx) {
    padding_reader* rdr = (padding_reader*)ctx;
    return rdr->end - rdr->cur;
}

void padding_reader_eat(void* ctx, size_t n) {
    padding_reader* rdr = (padding_reader*)ctx;
    rdr->cur += n;
}

// Value Buffer ----------------------------------------------------------------

static const uint64_t KNULL     = 0;

static const uint64_t BOOL      = 2;
static const uint64_t FALSE     = BOOL; // 2
static const uint64_t TRUE      = (1 << 3) | BOOL; // 10

static const uint64_t NUMBER    = 3;
static const uint64_t UINT      = NUMBER; // 3
static const uint64_t SINT      = (1 << 3) | NUMBER; // 11
static const uint64_t REAL      = (2 << 3) | NUMBER; // 19
static const uint64_t RAWNUMBER = (3 << 3) | NUMBER; // 27

static const uint64_t STRING        = 4;
static const uint64_t STRING_COMMON = STRING; // 4
static const uint64_t STRING_HASESC = (1 << 3) | STRING; // 12

static const uint64_t OBJECT    = 6;
static const uint64_t ARRAY     = 7;

static const uint64_t POS_MASK      = (~(uint64_t)0) << 32;
static const uint64_t POS_BITS      = 32;
static const uint64_t TYPE_MASK     = 0xFF;
static const uint64_t TYPE_BITS     = 8;
static const uint64_t CON_LEN_MASK  = (~(uint64_t)0) >> 32;
static const uint64_t CON_LEN_BITS  = 32;

typedef union {
    uint64_t u64;
    int64_t  i64;
    double   f64;
} num;


typedef struct {
    /// (low)
    /// | type (8 bits) | fsm depth (24 bits) | pos (32 bits)    |
    uint64_t    typ;
    /// (low)
    /// |   number val: f64/i64/u64  (64 bits)                   |
    /// |   container val: next offset (32 bits) | len (32 bits) |
    /// |   string val   : len(64 bits)                          |
    union {
        num         num;
        uint64_t    slen;
        int64_t     parent;
        struct {
            uint32_t    len;
            uint32_t    next;
        }           con;
    } val;
} node;

// JSON Stat -------------------------------------------------------------------

typedef struct {
    uint32_t object;
    uint32_t array;
    uint32_t string;
    uint32_t number;
    uint32_t array_elems;
    uint32_t object_keys;
    uint32_t max_depth;
} json_stat;

// JSON Visitor ----------------------------------------------------------------

typedef struct {
    bool (*on_null)(void* ctx, size_t pos);
    bool (*on_bool)(void* ctx, size_t pos, bool val);
    bool (*on_int)(void* ctx, size_t pos, int64_t val);
    bool (*on_uint)(void* ctx, size_t pos, uint64_t val);
    bool (*on_float)(void* ctx, size_t pos, double val);
    bool (*on_string)(void* ctx, size_t pos, size_t len, bool has_esc);
    bool (*on_number)(void* ctx, size_t pos, size_t len);
    uint64_t* (*on_array_start)(void* ctx, size_t pos);
    uint64_t* (*on_array_end)(void* ctx, size_t len);
    uint64_t* (*on_object_start)(void* ctx, size_t pos);
    uint64_t* (*on_object_end)(void* ctx, size_t len);
    bool (*on_key)(void* ctx, size_t pos, size_t len, bool has_esc);
    void* ctx;
} visitor;

typedef struct {
    node*        cur;
    int64_t      parent;
    uint64_t     depth;
    node*        start;
    const node*  end;
    json_stat    stat;
} node_buf;

static always_inline uint64_t node_pack_type(uint64_t typ, size_t pos) {
    return typ | (((uint64_t)pos) << POS_BITS);
}

static always_inline void value_inc(uint64_t* typ) {
    *typ += (1 << TYPE_BITS);
}

static always_inline uint64_t get_type(node* node) {
    return node->typ & TYPE_MASK;
}

static always_inline uint64_t get_pos(node* node) {
    return node->typ >> POS_BITS;
}

static always_inline size_t get_count(uint64_t typ) {
    return (typ & ~POS_MASK) >> TYPE_BITS;
}

static always_inline bool has_parent(node_buf* buf) {
    return buf->parent != -1;
}

static always_inline node* node_buf_parent(node_buf* buf) {
    if (unlikely(!has_parent(buf))) {
        return NULL;
    }
    return &buf->start[buf->parent];
}

static always_inline node* node_buf_top(node_buf* buf) {
    return buf->cur - 1;
}

static always_inline bool top_is_key(node_buf* buf) {
    ssize_t offset = (buf->cur - 1) - node_buf_parent(buf);
    return (offset % 2) != 0;
}

#define node_buf_grow(buf, n) { \
    if (unlikely(buf->cur + n > buf->end)) \
        return false; \
    }

static always_inline bool node_on_null(void* ctx, size_t pos) {
    node_buf* buf = (node_buf*)ctx;
    buf->cur->typ = node_pack_type(KNULL, pos);
    buf->cur++;
    node_buf_grow(buf, 1);
    return true;
}

static always_inline bool node_on_num(void* ctx, size_t pos, uint64_t typ, num num) {
    node_buf* buf = (node_buf*)ctx;
    buf->cur->typ = node_pack_type(typ, pos);
    buf->cur->val.num = num;
    buf->stat.number++;
    buf->cur++;
    node_buf_grow(buf, 1);
    return true;
}

static always_inline bool node_on_bool(void* ctx, size_t pos, bool val) {
    node_buf* buf = (node_buf*)ctx;
    buf->cur->typ = node_pack_type(val ? TRUE : FALSE, pos);
    buf->cur++;
    node_buf_grow(buf, 1);
    return true;
}

static always_inline bool node_on_int(void* ctx, size_t pos, int64_t val) {
    return node_on_num(ctx, pos, SINT, (num){.i64 = val});
}

static always_inline bool node_on_uint(void* ctx, size_t pos, uint64_t val) {
    return node_on_num(ctx, pos, UINT, (num){.u64 = val});
}

static always_inline bool node_on_float(void* ctx, size_t pos, double val) {
    return node_on_num(ctx, pos, REAL, (num){.f64 = val});
}

static always_inline bool node_on_number(void* ctx, size_t pos, size_t len) {
    node_buf* buf = (node_buf*)ctx;
    buf->cur->typ = node_pack_type(RAWNUMBER, pos);
    buf->cur->val.slen = len;
    buf->cur++;
    buf->stat.number++;
    node_buf_grow(buf, 1);
    return true;
}

static always_inline bool node_on_string(void* ctx, size_t pos, size_t len, bool has_esc) {
    node_buf* buf = (node_buf*)ctx;
    buf->cur->typ = node_pack_type(has_esc ? STRING_HASESC : STRING_COMMON, pos);
    buf->cur->val.slen = len;
    buf->cur++;
    buf->stat.string++;
    node_buf_grow(buf, 1);
    return true;
}

static always_inline bool node_on_key(void* ctx, size_t pos, size_t len, bool has_esc) {
    node_buf* buf = (node_buf*)ctx;
    buf->cur->typ = node_pack_type(has_esc ? STRING_HASESC : STRING_COMMON, pos);
    buf->cur->val.slen = len;
    buf->cur++;
    node_buf_grow(buf, 1);
    return true;
}

static always_inline uint64_t* node_on_container_start(void* ctx, size_t pos, bool is_obj) {
    node_buf* buf = (node_buf*)ctx;
    buf->cur->typ = node_pack_type(is_obj ? OBJECT : ARRAY, pos);
    buf->cur->val.parent = buf->parent;
    buf->parent = buf->cur - buf->start;
    buf->cur++;
    buf->depth++;
    if (unlikely(buf->cur + 1 > buf->end)) {
        return NULL;
    }
    return &(node_buf_parent(buf)->typ);
}

static always_inline uint64_t* node_on_container_end(void* ctx, size_t len) {
    node_buf* buf = (node_buf*)ctx;
    node* p = node_buf_parent(buf);
    buf->parent = node_buf_parent(buf)->val.parent;
    p->val.con.next = buf->cur - p;
    p->val.con.len = len;
    p->typ &= POS_MASK | TYPE_MASK;     // remove recorded depth
    if (buf->depth > buf->stat.max_depth) {
        buf->stat.max_depth = buf->depth;
        if (buf->depth > MAX_DEPTTH) {
            return NULL;
        }
    }
    buf->depth--;
    if (!has_parent(buf)) {
        return NULL;
    }
    return &(node_buf_parent(buf)->typ);
}

static always_inline uint64_t* node_on_array_start(void* ctx, size_t pos) {
    return node_on_container_start(ctx, pos, false);
}

static always_inline uint64_t* node_on_object_start(void* ctx, size_t pos) {
    return node_on_container_start(ctx, pos, true);
}

static always_inline uint64_t* node_on_array_end(void* ctx, size_t len) {
    node_buf* buf = (node_buf*)ctx;
    buf->stat.array++;
    buf->stat.array_elems += (uint32_t)len;
    return node_on_container_end(ctx, len);
}

static always_inline uint64_t* node_on_object_end(void* ctx, size_t len) {
    node_buf* buf = (node_buf*)ctx;
    buf->stat.object++;
    buf->stat.object_keys += (uint32_t)len;
    return node_on_container_end(ctx, len);
}

// Parser ----------------------------------------------------------------

// Parser/SkipSpace ------------------------------------------------------

typedef struct {
    uint8_t*       cur;
    uint64_t    bits;
} nonspace_block;

static always_inline bool is_space(char c) {
    return c == ' ' || c == '\t' || c == '\n' || c == '\r';
}

static always_inline void nonspace_block_init(nonspace_block* slf) {
    slf->bits = 0;
    slf->cur = NULL;
}

static always_inline uint64_t get_nonspace_bits(const uint8_t* s) {
#if defined(__AVX2__)
    __m256i space_tab = _mm256_setr_epi8(
        '\x20', 0, 0, 0, 0, 0, 0, 0,
         0, '\x09', '\x0A', 0, 0, '\x0D', 0, 0,
        '\x20', 0, 0, 0, 0, 0, 0, 0,
         0, '\x09', '\x0A', 0, 0, '\x0D', 0, 0
    );

    __m256i lo = _mm256_loadu_si256((__m256i*)s);
    __m256i hi = _mm256_loadu_si256((__m256i*)(s + 32));
    __m256i shuf_lo = _mm256_shuffle_epi8(space_tab, lo);
    __m256i shuf_hi = _mm256_shuffle_epi8(space_tab, hi);
    uint32_t mask_lo = (uint32_t)_mm256_movemask_epi8(_mm256_cmpeq_epi8(lo, shuf_lo));
    uint32_t mask_hi = (uint32_t)_mm256_movemask_epi8(_mm256_cmpeq_epi8(hi, shuf_hi));
    return ~((uint64_t)mask_lo | ((uint64_t)(mask_hi) << 32));
#else
    __m128i space_tab = _mm_setr_epi8(
        '\x20', 0, 0, 0, 0, 0, 0, 0,
         0, '\x09', '\x0A', 0, 0, '\x0D', 0, 0
    );
    __m128i lo = _mm_loadu_si128((__m128i*)s);
    __m128i hi = _mm_loadu_si128((__m128i*)(s + 16));
    __m128i lo1 = _mm_loadu_si128((__m128i*)(s + 32));
    __m128i hi1 = _mm_loadu_si128((__m128i*)(s + 48));
    __m128i shuf_lo = _mm_shuffle_epi8(space_tab, lo);
    __m128i shuf_hi = _mm_shuffle_epi8(space_tab, hi);
    __m128i shuf_lo1 = _mm_shuffle_epi8(space_tab, lo1);
    __m128i shuf_hi1 = _mm_shuffle_epi8(space_tab, hi1);
    uint32_t mask_lo = (uint32_t)_mm_movemask_epi8(_mm_cmpeq_epi8(lo, shuf_lo));
    uint32_t mask_hi = (uint32_t)_mm_movemask_epi8(_mm_cmpeq_epi8(hi, shuf_hi));
    uint32_t mask_lo1 = (uint32_t)_mm_movemask_epi8(_mm_cmpeq_epi8(lo1, shuf_lo1));
    uint32_t mask_hi1 = (uint32_t)_mm_movemask_epi8(_mm_cmpeq_epi8(hi1, shuf_hi1));
    return ~((uint64_t)mask_lo | ((uint64_t)(mask_hi) << 16) | ((uint64_t)(mask_lo1) << 32) | ((uint64_t)(mask_hi1) << 48));
#endif
}

static always_inline char skip_space_impl(nonspace_block* slf, reader* rdr) {
    // fast path 1: check at most two spaces for compat JSON or 
    // some cases in pretty JSON(such as ` "name": "balabala" `)
    repeat_2(
        char c = rdr->next(rdr->ctx);
        if (likely(c != SONIC_NONE)) {
            if (likely(!is_space(c))) {
                return c;
            }
        }
    );

    // fast path 2: reuse the bitmap for short key or numbers
    // XXX: simplify the branch
    uint8_t** cur = rdr->cur(rdr->ctx);
    size_t offset = *cur - slf->cur;
    if (offset < 64) {
        uint64_t bits = clear_low_bits(slf->bits, offset);
        if (!is_zero(bits)) {
            size_t pos = trailing_zeros(bits);
            *cur =  slf->cur + pos + 1;
            return slf->cur[pos];
        } else {
            *cur = slf->cur + 64;
        }
    }

    // slow path: use simd to accelerate skipping space
    uint8_t* block = rdr->peek_n(rdr->ctx, 64);
    while (block != NULL) {
        uint64_t bits = get_nonspace_bits(block);
        if (!is_zero(bits)) {
            slf->bits = bits;
            slf->cur = block;
            size_t pos = trailing_zeros(bits);
            *cur = block + pos + 1;
            return block[pos];
        }
        rdr->eat(rdr->ctx, 64);
        block = rdr->peek_n(rdr->ctx, 64);
    }

    // remain bytes, do with scalar code
    char c = rdr->next(rdr->ctx);
    while (c != SONIC_NONE) {
        if (!is_space(c)) {
            return c;
        }
        c = rdr->next(rdr->ctx);
    }

    // not found non-space char until eof
    return SONIC_NONE;
}

static always_inline char skip_space(nonspace_block* slf, reader* rdr) {
    return skip_space_impl(slf, rdr);
}
// Parser/ParseString -------------------------------------------------------

// TODO: reuse block for short strings
typedef struct {
    uint32_t  bs;
    uint32_t  quote;
    uint32_t  esc;
} string_block;

static always_inline string_block string_block_new(uint8_t* s) {
    v256u v = v256_loadu((uint8_t*)s);
    return (string_block){
        .bs = mask256_tobitmask(v256_eq(v, v256_splat('\\'))),
        .quote = mask256_tobitmask(v256_eq(v, v256_splat('"'))),
        .esc = mask256_tobitmask(v256_le(v, v256_splat('\x1f')))
    };
}

static always_inline bool has_quote_first(string_block* block) {
    return at_first(block->quote, block->bs | block->esc);
}

static always_inline bool has_backslash(string_block* block) {
    return  at_first(block->bs, block->quote);
}

static always_inline bool has_unescaped(string_block* block) {
    return  at_first(block->esc, block->quote);
}

static const uint8_t ESCAPED_TAB[256] = {
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, '"', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, '/', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    '\\', 0, 0, 0, 0, 0, '\x08', /* \b */
    0, 0, 0, '\x0c', /* \f */
    0, 0, 0, 0, 0, 0, 0, '\n', 0, 0, 0, '\r', 0, '\t', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
    0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
};

static always_inline ssize_t handle_unicode(char** sp, char **dp) {
    uint32_t r0;
    uint32_t r1;

    /* check for hexadecimal characters */
    if (!unhex16_is(*sp + 2)) {
        return SONIC_INVALID_ESCAPED_UTF;
    }

    /* decode the code-point */
    r0 = unhex16_fast(*sp + 2);
    *sp += 6;

retry_decode:
    /* ASCII characters, unlikely */
    if (unlikely(r0 <= 0x7f)) {
        *(*dp)++ = (char)r0;
        return 0;
    }

    /* latin-1 characters, unlikely */
    if (unlikely(r0 <= 0x07ff)) {
         *(*dp)++     = (char)(0xc0 | (r0 >> 6));
         *(*dp)++ = (char)(0x80 | (r0 & 0x3f));
        return 0;
    }

    /* 3-byte characters, likely */
    if (likely(r0 < 0xd800 || r0 > 0xdfff)) {
        *(*dp)++     = (char)(0xe0 | ((r0 >> 12)       ));
        *(*dp)++     = (char)(0x80 | ((r0 >>  6) & 0x3f));
        *(*dp)++     = (char)(0x80 | ((r0      ) & 0x3f));
        return 0;
    }

    /* surrogate half, must follows by the other half */
    if (r0 > 0xdbff || **sp != '\\' || *(*sp + 1) != 'u') {
        unirep(dp);
        return 0;
    }

    /* check the hexadecimal escape */
    if (!unhex16_is((char*)*sp + 2)) {
        return SONIC_INVALID_ESCAPED_UTF;
    }

    /* decode the second code-point */
    r1 = unhex16_fast((char*)*sp + 2);
    *sp += 6;

    /* it must be the other half */
    if (r1 < 0xdc00 || r1 > 0xdfff) {
        r0 = r1;
        unirep(dp);
        goto retry_decode;
    }

    /* merge two surrogates */
    r0 = (r0 - 0xd800) << 10;
    r1 = (r1 - 0xdc00) + 0x010000;
    r0 += r1;

    /* encode the character */
    *(*dp)++ = (char)(0xf0 | ((r0 >> 18)       ));
    *(*dp)++ = (char)(0x80 | ((r0 >> 12) & 0x3f));
    *(*dp)++ = (char)(0x80 | ((r0 >>  6) & 0x3f));
    *(*dp)++ = (char)(0x80 | ((r0      ) & 0x3f));
    return 0;
}

// positive is length
// negative is - error_code
static always_inline long parse_string_inplace(uint8_t** cur, bool* has_esc, uint64_t opts) {
    string_block block;
    uint8_t* start = *cur;
    v256u v;

    // breakpoint();
    while (true) {
        block = string_block_new(*cur);
        if (has_quote_first(&block)) {
            *cur += trailing_zeros(block.quote) + 1; // skip the quote char
            *has_esc = false;
            return *cur - start - 1;
        }

        if (unlikely(has_backslash(&block))) {
            break;
        }

        if (unlikely((opts & F_VALIDATE_STRING) != 0 && has_unescaped(&block))) {
            *cur += trailing_zeros(block.esc);
            return -SONIC_CONTROL_CHAR;
        }

        *cur += 32;
    }

    // deal with the escaped string
    *has_esc = true;
    *cur += trailing_zeros(block.bs);
    uint8_t* dst = *cur;
    uint8_t esc;

escape:
    esc = *(*cur + 1);
    if (likely(esc) != 'u') {
        if (unlikely(ESCAPED_TAB[esc]) == 0) {
            return -SONIC_INVALID_ESCAPED;
        }
        *cur += 2;
        *dst++ = ESCAPED_TAB[esc];
    } else if (handle_unicode((char**)cur, (char**)&dst) != 0) {
        return -SONIC_INVALID_ESCAPED_UTF;
    }

    // check continous escaped char
    if (**cur == '\\') {
        goto escape;
    }

find_and_move:
    v = v256_loadu((uint8_t*)*cur);
    block =  (string_block){
        .bs = mask256_tobitmask(v256_eq(v, v256_splat('\\'))),
        .quote = mask256_tobitmask(v256_eq(v, v256_splat('"'))),
        .esc = mask256_tobitmask(v256_le(v, v256_splat('\x1f')))
    };

    if (has_quote_first(&block)) {
        // while **src != b'"' {
        //     *dst = **src;
        //     dst = dst.add(1);
        //     *src = src.add(1);
        // }
        while (true) {
            repeat_8( {
                if (**cur != '"') {
                    *dst++ = **cur;
                    *cur += 1;
                } else {
                    *cur += 1;
                    return dst - start;
                }
            });
        }
    }

    if (unlikely((opts & F_VALIDATE_STRING) != 0 && has_unescaped(&block))) {
        *cur += trailing_zeros(block.esc);
        return -SONIC_CONTROL_CHAR;
    }

    if (has_backslash(&block)) {
        // while **src != b'\\' {
        //     *dst = **src;
        //     dst = dst.add(1);
        //     *src = src.add(1);
        // }
        while (true) {
            repeat_8( {
                if (**cur != '\\') {
                    *dst++ = **cur;
                    *cur += 1;
                } else {
                    goto escape;
                }
            });
        }
    }

    v256_storeu(v, dst);
    *cur += 32;
    dst += 32;
    goto find_and_move;
}

// Parser/ParseNumber -------------------------------------------------------

// TODO: use SIMD to optimize it

typedef struct {
    uint64_t   typ;
    num        num;
} parse_num;

static always_inline bool is_dec_digit(char c) {
    return c >= '0' && c <= '9';
}

#define should_digit(c) { \
    if (unlikely(!is_dec_digit(c))) { \
        return SONIC_INVALID_NUM; \
    } \
}

static const size_t FLOATING_LONGEST_DIGITS = 17;

static always_inline error_code parse_number(uint8_t** cur, parse_num* num, bool neg, GoSlice* dbuf) {
    uint64_t sig = 0;
    int exp = 0;
    bool truc = false;
    int exp_neg = 1;
    int exp_dec = 0;
    double val = 0.0;
    uint8_t* start = *cur;
    ssize_t frac_cnt = 0;
    ssize_t needed = 0;
    ssize_t cnt = 0;

    if (**cur == '0') {
        *cur += 1;

        if (**cur != '.' && **cur != 'e' && **cur != 'E') {
            *num = (parse_num){ neg? SINT: UINT, {.i64 = 0}};
            return SONIC_OK;
        }

        if (**cur == '.') {
            // skip continous zeros
            *cur += 1;
            uint8_t* dot_pos = *cur;
            should_digit(**cur);
            
            // skip zeros until significant
            while (**cur == '0') {
                *cur += 1;
            }

            // special case: 0.000e....
            if (unlikely(**cur == 'e' || **cur == 'E')) {
                 goto parse_exponent;
            }

            exp -= (*cur) - dot_pos;
            goto parse_fraction;
        }

        // parse 0e...
        goto parse_exponent;
    } else {
        while (is_dec_digit(**cur)) {
            sig = sig * 10 + (**cur - '0');
            *cur += 1;
            cnt ++;
        }

        if (unlikely(cnt == 0)) {
            return SONIC_INVALID_NUM;
        }

        // fallback for long integer
        if (unlikely(cnt > 19)) {
            *cur -= cnt;
            sig = 0;
            cnt = 0;
            while (is_dec_digit(**cur) && cnt < 19) {
                sig = sig * 10 + (**cur - '0');
                *cur += 1;
                cnt ++;
            }

            // overflow for u64 sig, mark as truncated
            while (is_dec_digit(**cur)) {
                *cur += 1;
                truc = true;
                exp += 1;
            }
        }

        if (**cur == '.') { // parse fraction
            *cur += 1;
            should_digit(**cur);
            goto parse_fraction;
        }
        
        if (likely(**cur != 'e' && **cur != 'E')) { // parse integer
            if (likely(exp == 0)) {
                if (neg) {
                    if (sig > (1ull << 63)) {
                        *num = (parse_num){REAL, {.f64 = (double)sig * -1}};
                    } else {
                        *num = (parse_num){SINT, {.i64 = -sig}};
                    }
                } else {
                    *num = (parse_num){UINT, {.u64 = sig}};
                }
                return SONIC_OK;
            }

            // maybe not overflow for u64
            if (unlikely(exp == 1)) {
                uint64_t res = 0;
                if (__builtin_mul_overflow(sig, 10, &res) || __builtin_add_overflow(res, *(*cur - 1) - '0' , &res)) {
                    goto parse_float;
                }
                if (neg) {
                    *num = (parse_num){REAL, {.f64 = (double)res * -1}};
                } else {
                    *num = (parse_num){UINT, {.u64 = res}};
                }
                return SONIC_OK;
                
            }
            // overflow for u64
            goto parse_float;
        }
        goto parse_exponent;
    }

parse_fraction:
    frac_cnt = 0;
    needed = FLOATING_LONGEST_DIGITS - cnt;
    while (needed > 0 && is_dec_digit(**cur)) {
        sig = sig * 10 + (**cur - '0');
        *cur += 1;
        needed --;
        frac_cnt += 1;
    }
    exp -= frac_cnt;
    while (is_dec_digit(**cur)) {
        *cur += 1;
        truc = true;
    }

    if (likely(**cur != 'e' && **cur != 'E')) {
        goto parse_float;
    } else {
        goto parse_exponent;
    }

parse_exponent:
    *cur += 1;

    // parse exponent sign
    if (**cur == '-') {
        *cur += 1;
        exp_neg = -1;
    } else if (**cur == '+') {
        *cur += 1;
    }

    should_digit(**cur);
    cnt = 0;
    while (is_dec_digit(**cur)) {
        exp_dec = exp_dec * 10 + (**cur - '0');
        *cur += 1;
        cnt ++;
    }

    // exponent is overflow
    if (unlikely(cnt >= 10)) {
        exp = 0;
        exp_dec = 10000;
    }
    
    exp = exp + exp_dec * exp_neg;
parse_float:
    /* when fast algorithms failed, use slow fallback.*/
    if(!atof_fast(sig, exp, neg ? -1: 1, truc, &val)) {
        val = atof_native_1((const char *)start, *cur - start , dbuf->buf , dbuf->cap) * (neg? -1 : 1);
    }

    /* check parsed double val */
    if (is_infinity(val)) {
        return SONIC_FLOAT_INF;
    }

    /* update the result */
    *num = (parse_num){REAL, {.f64 = val}};
    return SONIC_OK;
}

static always_inline long skip_positive2(uint8_t** cur, size_t len) {
    // skip number
    long r = do_skip_number((const char*)*cur, len);
    if (r < 0) {
        *cur -= r + 1;
        return -SONIC_INVALID_NUM;
    }

    *cur += r;
    return r;
}

// Parser/ParseLiteral ------------------------------------------------------

static always_inline error_code parse_true(uint8_t** cur) {
    if (unlikely((*cur)[0] != 'r' || (*cur)[1] != 'u' || (*cur)[2] != 'e')) {
        if (likely((*cur)[0] != 'r')) {
             *cur += 1;
            return SONIC_INVALID_LITERAL;
        }
        if (likely((*cur)[1] != 'u')) {
             *cur += 2;
            return SONIC_INVALID_LITERAL;
        }
        if (likely((*cur)[2] != 'e')) {
             *cur += 3;
            return SONIC_INVALID_LITERAL;
        }
    }
    *cur += 3;
    return SONIC_OK;
}

static always_inline error_code parse_false(uint8_t** cur) {
    if (unlikely((*cur)[0] != 'a' || (*cur)[1] != 'l' || (*cur)[2] != 's' || (*cur)[3] != 'e')) {
        if (likely((*cur)[0] != 'a')) {
             *cur += 1;
            return SONIC_INVALID_LITERAL;
        }
        if (likely((*cur)[1] != 'l')) {
             *cur += 2;
            return SONIC_INVALID_LITERAL;
        }
        if (likely((*cur)[2] != 's')) {
             *cur += 3;
            return SONIC_INVALID_LITERAL;
        }
        if (likely((*cur)[3] != 'e')) {
             *cur += 4;
            return SONIC_INVALID_LITERAL;
        }
    }
    *cur += 4;
    return SONIC_OK;
}

static always_inline error_code parse_null(uint8_t** cur) {
    if (unlikely((*cur)[0] != 'u' || (*cur)[1] != 'l' || (*cur)[2] != 'l')) {
        if (likely((*cur)[0] != 'u')) {
             *cur += 1;
            return SONIC_INVALID_LITERAL;
        }
        if (likely((*cur)[1] != 'l')) {
             *cur += 2;
            return SONIC_INVALID_LITERAL;
        }
        if (likely((*cur)[2] != 'l')) {
             *cur += 3;
            return SONIC_INVALID_LITERAL;
        }
    }
    *cur += 3;
    return SONIC_OK;
}

// Parser ----------------------------------------------------------------

static always_inline bool visit_number(visitor* vis, size_t pos, parse_num* num) {
    switch (num->typ) {
        case UINT:
            return vis->on_uint(vis->ctx, pos, num->num.u64);
        case SINT:
            return vis->on_int(vis->ctx, pos, num->num.i64);
        case REAL:
            return vis->on_float(vis->ctx, pos, num->num.f64);
        default:
            return false;
    }
}

#define offset_from(ptr)  ((ptr) - start)

#define check_state(vis) { \
    state = (vis); \
    if (state == NULL) { \
        return SONIC_VISIT_FAILED; \
    } \
}

#define check_visit() { \
    if (unlikely(!visited)) { \
        return SONIC_VISIT_FAILED; \
    } \
}

#define check_error() { \
    if (unlikely(err != SONIC_OK)) { \
        return err; \
    } \
}

// API ----------------------------------------------------------------

typedef struct {
    GoString    json;
    GoSlice     padded;
    GoSlice     nodes;
    GoSlice     dbuf;
    GoSlice     backup;

    uint64_t    opt;
    // JSON buffer cursor
    uint64_t    start;
    uint64_t    cur;
    uint64_t    end;
    nonspace_block nbk;

    // node buffer cursor
    node_buf    nbuf;
    bool        uf8Inv;
    bool        is_eface;
} GoParser;


static uint64_t PADDING_SIZE = 64;

static always_inline error_code parse(GoParser* slf, reader* rdr, visitor* vis) {
    uint8_t** cur   = rdr->cur(rdr->ctx);
    uint8_t* start  = (uint8_t*)(slf->start);
    uint64_t* state = NULL;
    long slen       = 0;
    bool neg        = false;
    ssize_t remain  = 0;
    parse_num num   = {0};
    void* ctx       = vis->ctx;
    bool has_esc    = false;
    error_code err  = SONIC_OK;
    ssize_t pos     = 0;
    bool visited    = true;
    char c          = 0;

    // dispatch for current node status
    node_buf* nodes = (node_buf*)(vis->ctx);
    if (unlikely(has_parent(nodes))) {
        // restore the current state from the parent node
        state = &node_buf_parent(nodes)->typ;
        switch (node_buf_top(nodes)->typ & TYPE_MASK) {
            case OBJECT: {
                c = skip_space(&slf->nbk, rdr);
                if (c == '}') {
                    state = vis->on_object_end(ctx, 0);
                    goto scope_end;
                }
                goto obj_key;
            }
            case ARRAY: {
                c = skip_space(&slf->nbk, rdr);
                if (c == ']') {
                    state = vis->on_array_end(ctx, 0);
                    goto scope_end;
                }
                goto arr_val;
            }
            default: {
                if (get_type(node_buf_parent(nodes)) == OBJECT) {
                    if (top_is_key(nodes)) {
                        goto obj_val;
                    } else {
                        c = skip_space(&slf->nbk, rdr);
                        goto obj_cont;
                    }
                } else {
                    c = skip_space(&slf->nbk, rdr);
                    goto arr_cont;
                }
            }
        }
    }

    c = skip_space(&slf->nbk, rdr);
    pos = offset_from(*cur - 1);
    switch (c) {
        case '{':
            check_state(vis->on_object_start(ctx, pos));
            c = skip_space(&slf->nbk, rdr);
            if (c == '}') {
                state = vis->on_object_end(ctx, 0);
                goto scope_end;
            }
            goto obj_key;
        case '[':
            check_state(vis->on_array_start(ctx, pos));
            c = skip_space(&slf->nbk, rdr);
            if (c == ']') {
                state = vis->on_array_end(ctx, 0);
                goto scope_end;
            }
            goto arr_val;
        case '-': neg = true;
        case '0'...'9':
            *cur -= !neg;
            /* fallthrough */
            if (slf->opt & F_USE_NUMBER) {
                remain = rdr->remain(rdr->ctx) + !neg;
                slen = skip_positive2(cur, remain);
                if (slen < 0) {
                    err =  (error_code)(-slen);
                }
               visited = vis->on_number(ctx, pos, slen + neg);
            } else {
                err =  parse_number(cur, &num, neg, &slf->dbuf);
                visited = visit_number(vis, pos, &num);
            }
            neg = false;
            break;
        case '"':
            slen = parse_string_inplace(cur, &has_esc, slf->opt);
            if (slen < 0) {
                err =  (error_code)(-slen);
            }
            // check eof 
            if (rdr->remain(rdr->ctx) < 0) {
                err = SONIC_EOF;
            }
            visited = vis->on_string(ctx, pos + 1, slen, has_esc);
            break;
        case 't':
            err =  parse_true(cur);
            visited = vis->on_bool(ctx, pos, true);
            break;
        case 'f':
            err =  parse_false(cur);
            visited = vis->on_bool(ctx, pos, false);
            break;
        case 'n':
            err =  parse_null(cur);
            visited = vis->on_null(ctx, pos);
            break;
        default: err =  SONIC_INVALID_CHAR;
    }
        
    check_error();
    check_visit();
    return err;

obj_key:
    if (unlikely(c != '"')) {
        return SONIC_EXPECT_KEY;
    }

    // parse key
    pos = offset_from(*cur);

    slen = parse_string_inplace(cur, &has_esc, slf->opt);
    if (slen < 0) {
        err =  (error_code)(-slen);
        return err;
    }

    c = skip_space(&slf->nbk, rdr);
    if (unlikely(c != ':')) {
        return SONIC_EXPECT_COLON;
    }

    visited = vis->on_key(ctx, pos, slen, has_esc);
    check_visit();

obj_val:
    // parse value
    c = skip_space(&slf->nbk, rdr);
    pos = offset_from(*cur - 1);
    switch (c) {
        case '{':
            check_state(vis->on_object_start(ctx, pos));
            c = skip_space(&slf->nbk, rdr);
            if (c == '}') {
                state = vis->on_object_end(ctx, 0);
                goto scope_end;
            }
            goto obj_key;
        case '[':
            check_state(vis->on_array_start(ctx, pos));
            c = skip_space(&slf->nbk, rdr);
            if (c == ']') {
                state = vis->on_array_end(ctx, 0);
                goto scope_end;
            }
            goto arr_val;
        case '-': neg = true;
        case '0'...'9':
            *cur -= !neg;
            /* fallthrough */
            if (slf->opt & F_USE_NUMBER) {
                remain = rdr->remain(rdr->ctx) + !neg;
                slen = skip_positive2(cur, remain);
                if (slen < 0) {
                    err =  (error_code)(-slen);
                }
               visited = vis->on_number(ctx, pos, slen + neg);
            } else {
                err =  parse_number(cur, &num, neg, &slf->dbuf);
                visited = visit_number(vis, pos, &num);
            }
            neg = false;
            break;
        case '"':
            slen = parse_string_inplace(cur, &has_esc, slf->opt);
            if (slen < 0) {
                err =  (error_code)(-slen);
            }
            visited = vis->on_string(ctx, pos + 1, slen, has_esc);
            break;
        case 't':
            err =  parse_true(cur);
            visited = vis->on_bool(ctx, pos, true);
            break;
        case 'f':
            err =  parse_false(cur);
             visited = vis->on_bool(ctx, pos, false);
            break;
        case 'n':
            err =  parse_null(cur);
             visited = vis->on_null(ctx, pos);
            break;
        default: err =  SONIC_INVALID_CHAR;
    }

    check_error();
    check_visit();
    c = skip_space(&slf->nbk, rdr);

obj_cont:
    value_inc(state);
    if (likely(c == ',')) {
        c = skip_space(&slf->nbk, rdr);
        goto obj_key;
    }

    if (unlikely(c != '}')) {
        err =  SONIC_EXPECT_OBJ_COMMA_OR_END;
        return err;
    }

    state = vis->on_object_end(ctx, get_count(*state));

scope_end:
    if (unlikely(state == NULL)) {
        return err;
    }

    c = skip_space(&slf->nbk, rdr);
    if (((*state) & TYPE_MASK) == OBJECT) {
        goto obj_cont;
    } else {
        goto arr_cont;
    }

arr_val:
    pos = offset_from(*cur - 1);
    switch (c) {
        case '{':
            check_state(vis->on_object_start(ctx, pos));
            c = skip_space(&slf->nbk, rdr);
            if (c == '}') {
                state = vis->on_object_end(ctx, 0);
                goto scope_end;
            }
            goto obj_key;
        case '[':
            check_state(vis->on_array_start(ctx, pos));
            c = skip_space(&slf->nbk, rdr);
            if (c == ']') {
                state = vis->on_array_end(ctx, 0);
                goto scope_end;
            }
            goto arr_val;
        case '-': neg = true;
        case '0'...'9':
            *cur -= !neg;
            /* fallthrough */
            if (slf->opt & F_USE_NUMBER) {
                remain = rdr->remain(rdr->ctx) + !neg;
                slen = skip_positive2(cur, remain);
                if (slen < 0) {
                    err = (error_code)(-slen);
                }
               visited = vis->on_number(ctx, pos, slen + neg);
            } else {
                err =  parse_number(cur, &num, neg, &slf->dbuf);
                visited = visit_number(vis, pos, &num);
            }
            neg = false;
            break;
        case '"':
            slen = parse_string_inplace(cur, &has_esc, slf->opt);
            if (slen < 0) {
                err =  (error_code)(-slen);
            }
            visited = vis->on_string(ctx, pos + 1, slen, has_esc);
            break;
        case 't':
            err =  parse_true(cur);
            visited = vis->on_bool(ctx, pos, true);
            break;
        case 'f':
            err =  parse_false(cur);
            visited = vis->on_bool(ctx, pos, false);
            break;
        case 'n':
            err =  parse_null(cur);
            visited =  vis->on_null(ctx, pos);
            break;
        default: err = SONIC_INVALID_CHAR;
    }

    check_error();
    check_visit();
    c = skip_space(&slf->nbk, rdr);

arr_cont:
    value_inc(state);
    if (likely(c == ',')) {
        c = skip_space(&slf->nbk, rdr);
        goto arr_val;
    }

    if (unlikely(c != ']')) {
        err =  SONIC_EXPECT_ARR_COMMA_OR_END;
        return err;
    }

    state = vis->on_array_end(ctx, get_count(*state));
    goto scope_end;
}


long parse_with_padding(void* p) {
    GoParser* gp = (GoParser*)p;
    node_buf* n = &(gp->nbuf);

    visitor vis = {
        .on_null = node_on_null,
        .on_bool = node_on_bool,
        .on_int = node_on_int,
        .on_uint = node_on_uint,
        .on_float = node_on_float,
        .on_string = node_on_string,
        .on_number = node_on_number,
        .on_array_start = node_on_array_start,
        .on_array_end = node_on_array_end,
        .on_object_start = node_on_object_start,
        .on_object_end = node_on_object_end,
        .on_key = node_on_key,
        .ctx = n,
    };

    padding_reader pad = {
        .cur = (uint8_t*)gp->cur,
        .end = (uint8_t*)gp->end,
        .start = (uint8_t*)gp->start,
    };

    reader rdr = {
        .ctx = &pad,
        .cur = padding_reader_cur,
        .next = padding_reader_next,
        .eat = padding_reader_eat,
        .remain = padding_reader_remain,
        .peek_n = padding_reader_peek_n,
    };

    error_code err = parse(gp, &rdr, &vis);
    gp->cur = (uint64_t)pad.cur;

    // check depth
    if (unlikely(gp->nbuf.stat.max_depth > MAX_DEPTTH)) {
        return SONIC_STACK_OVERFLOW;
    }

    return err;
}