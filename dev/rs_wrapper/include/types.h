
#pragma once

#include <stdint.h>
#include <stddef.h>
#include <stdbool.h>

typedef struct  {
    uint64_t pos;
    uint64_t len;
} Object;  // 16 bytes

typedef struct  {
    uint64_t pos;
    uint64_t len;
} Array;  // 16 bytes

typedef struct  {
    uint64_t len;
    const char* p;
} String;  // 16 bytes

typedef struct  {
    uint8_t t;
    uint8_t _len[7];
    char _ptr[8];
} Type;  // 16 bytes

typedef struct  {
    uint64_t len;
    const char* p;
} RawNumber;  // 16 bytes

typedef struct {
    uint64_t t;
    union {
        int64_t i64;
        uint64_t u64;
        double f64;
    } num;  // lower 8 bytes
} Number;

typedef struct  {
    uint64_t _[2];
} Data;

typedef union {
    Type t;
    Number n;
    Object o;
    Array a;
    String sv;
    RawNumber num;
    Data data;
} Node;

typedef struct {
    uint32_t error_code;
    size_t offset;
} ParseResult;

typedef struct {
    void* dom;
    const char* str_buf; // string start address, used to calcaue offset
    uint64_t str_len;
    Node* node;

    int64_t error_offset;
    const char* error_msg;
    bool has_utf8_lossy;
    uint64_t error_msg_len;
    uint64_t error_msg_cap;
} Document;

