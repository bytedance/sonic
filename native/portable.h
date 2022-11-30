#ifndef PORTABLE_H
#define PORTABLE_H

#if IS_ARM64

#define SIMDE_ENABLE_NATIVE_ALIASES
#include <simde/x86/avx2.h>
#include <simde/x86/sse2.h>

#define  _mm256_zeroupper() (void)(0)

#define USE_AVX2 1
#define USE_AVX  1

#else
#include <immintrin.h>
#endif

#endif