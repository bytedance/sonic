#
# Copyright 2021 ByteDance Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

.PHONY: all clean

CFLAGS := -mavx
CFLAGS += -mavx2
CFLAGS += -mbmi
CFLAGS += -mbmi2
CFLAGS += -mfma
CFLAGS += -msse
CFLAGS += -msse2
CFLAGS += -msse3
CFLAGS += -msse4
CFLAGS += -mssse3
CFLAGS += -mno-red-zone
CFLAGS += -ffast-math
CFLAGS += -fno-asynchronous-unwind-tables
CFLAGS += -fno-builtin
CFLAGS += -fno-exceptions
CFLAGS += -fno-rtti
CFLAGS += -fno-stack-protector
CFLAGS += -nostdlib
CFLAGS += -O3

NATIVE_ASM := $(wildcard native/*.S)
NATIVE_SRC := $(wildcard native/*.h)
NATIVE_SRC += $(wildcard native/*.c)

all: internal/native/native_amd64.s

clean:
	rm -vf internal/native/native_amd64.s output/*.s

internal/native/native_amd64.s: ${NATIVE_SRC} ${NATIVE_ASM} internal/native/native_amd64.go
	mkdir -p output
	clang ${CFLAGS} -S -o output/native.s native/native.c
	python3 tools/asm2asm/asm2asm.py internal/native/native_amd64.s output/native.s ${NATIVE_ASM}
	asmfmt -w internal/native/native_amd64.s
