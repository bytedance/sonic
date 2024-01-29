
#pragma once

#include "types.h"

Document sonic_rs_ffi_parse(const char *input, size_t len, uint64_t config);
void sonic_rs_ffi_free(void* dom, const char* msg, uint64_t msg_cap);
