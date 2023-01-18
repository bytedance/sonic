
#ifndef TYPES_H
#define TYPES_H

// NOTE: !NOT MODIFIED ONLY.
// This definitions are copied from internal/native/types/types.go.

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
#define ERR_MISMATCH    9
#define ERR_INVAL_UTF8  10

#define MAX_RECURSE     4096

#endif
