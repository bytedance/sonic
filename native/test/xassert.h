/*
 * Copyright 2022 ByteDance Inc.
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

#ifndef XASSERT_H
#define XASSERT_H


#ifndef DEBUG
    #define xassert(expr)     ((void)0)
#else
    #include "xprintf.h"
    #define xassert(expr) \
    ((expr)	\
    ? ((void)0)						\
    : _xassert(#expr, __FILE__, __LINE__, __PRETTY_FUNCTION__))

static void* raise = 0;
static void xabort() {
    *(int*)(raise) = 1;
}
static void _xassert(const char *assertion, const char *file, 
    const unsigned line, const char *func) {
    xprintf("%s:%u: %s Assertion `%s' failed.\n",
		      file, line, func ? func : "?", assertion);
    xabort();
}
#endif // DEBUG

#endif // XASSERT_H