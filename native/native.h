/*
 * Copyright 2021 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#ifndef NATIVE_H
#define NATIVE_H

#include <stdint.h>
#include <sys/types.h>
#include <immintrin.h>
#include <stdbool.h>

#include "types.h"

#define likely(v)       (__builtin_expect((v), 1))
#define unlikely(v)     (__builtin_expect((v), 0))
#define always_inline   inline __attribute__((always_inline)) 

#define as_m128p(v)     ((__m128i *)(v))
#define as_m128c(v)     ((const __m128i *)(v))
#define as_m256c(v)     ((const __m256i *)(v))
#define as_m128v(v)     (*(const __m128i *)(v))
#define as_uint64v(p)   (*(uint64_t *)(p))
#define is_infinity(v)  ((as_uint64v(&v) << 1) == 0xFFE0000000000000)

typedef struct {
    void * buf;
    size_t len;
    size_t cap;
} GoSlice;

static const uint8_t GO_KIND_MASK = (1 << 5) - 1;
typedef enum {
	Invalid = 0,
	Bool,
	Int,
	Int8,
	Int16,
	Int32,
	Int64,
	Uint,
	Uint8,
	Uint16,
	Uint32,
	Uint64,
	Uintptr,
	Float32,
	Float64,
	Complex64,
	Complex128,
	Array,
	Chan,
	Func,
	Interface,
	Map,
	Pointer,
	Slice,
	String,
	Struct,
	UnsafePointer,
} GoKind;

typedef struct {
    uint64_t size;
    uint64_t ptr_data;
    uint32_t hash;
    uint8_t  flags;
    uint8_t  align;
    uint8_t  filed_align;
    uint8_t  kind_flags;
    uint64_t traits;
    void*    gc_data;
    int32_t  str;
    int32_t  ptr_to_self;
} GoType;

typedef struct {
    GoType * type;
    void   * value;
} GoIface;

typedef struct {
    const char * buf;
    size_t       len;
} GoString;

typedef struct {
    long    t;
    double  d;
    int64_t i;
} JsonNumber;

typedef struct {
    long    vt;
    double  dv;
    int64_t iv;
    int64_t ep;
    char*   dbuf;
    ssize_t dcap;
} JsonState;

typedef struct {
    int64_t sp;
    int64_t vt[MAX_RECURSE];
} StateMachine;

int f64toa(char *out, double val);
int i64toa(char *out, int64_t val);
int u64toa(char *out, uint64_t val);

size_t lspace(const char *sp, size_t nb, size_t p);

ssize_t quote(const char *sp, ssize_t nb, char *dp, ssize_t *dn, uint64_t flags);
ssize_t unquote(const char *sp, ssize_t nb, char *dp, ssize_t *ep, uint64_t flags);
ssize_t html_escape(const char *sp, ssize_t nb, char *dp, ssize_t *dn);

long value(const char *s, size_t n, long p, JsonState *ret, uint64_t flags);
void vstring(const GoString *src, long *p, JsonState *ret, uint64_t flags);
void vnumber(const GoString *src, long *p, JsonState *ret);
void vsigned(const GoString *src, long *p, JsonState *ret);
void vunsigned(const GoString *src, long *p, JsonState *ret);

long skip_one(const GoString *src, long *p, StateMachine *m, uint64_t flags);
long skip_array(const GoString *src, long *p, StateMachine *m, uint64_t flags);
long skip_object(const GoString *src, long *p, StateMachine *m, uint64_t flags);

long skip_string(const GoString *src, long *p, uint64_t flags);
long skip_negative(const GoString *src, long *p);
long skip_positive(const GoString *src, long *p);
long skip_number(const GoString *src, long *p);

bool atof_eisel_lemire64(uint64_t mant, int exp10, int sgn, double *val);
double atof_native(const char *sp, ssize_t nb, char *dbuf, ssize_t cap);

long validate_string(const GoString *src, long *p);
long validate_one(const GoString *src, long *p, StateMachine *m);
long validate_utf8(const GoString *src, long *p, StateMachine *m);
long validate_utf8_fast(const GoString *src); 

long skip_one_fast(const GoString *src, long *p);
long get_by_path(const GoString *src, long *p, const GoSlice *path);
#endif
