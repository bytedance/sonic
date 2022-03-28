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

#define V_EOF           1
#define V_NULL          2
#define V_TRUE          3
#define V_FALSE         4
#define V_ARRAY         5
#define V_OBJECT        6
#define V_STRING        7
#define V_DOUBLE        8
#define V_INTEGER       9
#define V_KEY_SEP       10
#define V_ELEM_SEP      11
#define V_ARRAY_END     12
#define V_OBJECT_END    13
#define V_ATOF_NEED_FALLBACK 14

#define F_DBLUNQ        (1 << 0)
#define F_UNIREP        (1 << 1)

#define VS_NULL         0x6c6c756e      // 'null' in little endian
#define VS_TRUE         0x65757274      // 'true' in little endian
#define VS_ALSE         0x65736c61      // 'alse' in little endian ('false' without the 'f')

#define ERR_EOF         1
#define ERR_INVAL       2
#define ERR_ESCAPE      3
#define ERR_UNICODE     4
#define ERR_OVERFLOW    5
#define ERR_NUMBER_FMT  6
#define ERR_RECURSE_MAX 7
#define ERR_FLOAT_INF   8

#define MAX_RECURSE     65536

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
    char * buf;
    size_t len;
    size_t cap;
} GoSlice;

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
    int sp;
    int vt[MAX_RECURSE];
} StateMachine;

int f64toa(char *out, double val);
int i64toa(char *out, int64_t val);
int u64toa(char *out, uint64_t val);

size_t lzero(const char *sp, size_t nb);
size_t lspace(const char *sp, size_t nb, size_t p);

ssize_t quote(const char *sp, ssize_t nb, char *dp, ssize_t *dn, uint64_t flags);
ssize_t unquote(const char *sp, ssize_t nb, char *dp, ssize_t *ep, uint64_t flags);
ssize_t html_escape(const char *sp, ssize_t nb, char *dp, ssize_t *dn);

long value(const char *s, size_t n, long p, JsonState *ret, uint64_t flags);
void vstring(const GoString *src, long *p, JsonState *ret);
void vnumber(const GoString *src, long *p, JsonState *ret);
void vsigned(const GoString *src, long *p, JsonState *ret);
void vunsigned(const GoString *src, long *p, JsonState *ret);

long skip_one(const GoString *src, long *p, StateMachine *m);
long skip_array(const GoString *src, long *p, StateMachine *m);
long skip_object(const GoString *src, long *p, StateMachine *m);

long skip_string(const GoString *src, long *p);
long skip_negative(const GoString *src, long *p);
long skip_positive(const GoString *src, long *p);
long skip_number(const GoString *src, long *p);

bool atof_eisel_lemire64(uint64_t mant, int exp10, int sgn, double *val);
double atof_native(const char *sp, ssize_t nb, char* dbuf, ssize_t cap);

ssize_t utf8_validate(const char *sp, ssize_t nb);
long validate_string(const GoString *src, long *p);
long validate_one(const GoString *src, long *p, StateMachine *m);

#endif
