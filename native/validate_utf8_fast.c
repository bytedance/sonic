#include "scanning.h"
#if defined(__SVE__)
#include <arm_sve.h>
#endif
// validate_utf8_fast returns zero if valid, otherwise, the error position.
long validate_utf8_fast(const GoString *s) {
#if USE_AVX2
    /* fast path for valid utf8 */
    if (validate_utf8_avx2(s) == 0) {
        return 0;
    }
#elif defined(__SVE__)
    // Arm64 SVE
    if (validate_utf8_sve(s) == 0) {
        return 0;
    }
#endif
    return validate_utf8_errors(s);
}