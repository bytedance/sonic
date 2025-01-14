#!/bin/bash
#
# Copyright 2025 ByteDance Inc.
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


# This script is used to synchronize the code of cloudwego/iasm,
# and trim it for minimizing the size of the compiled binary file.

set -e

echo "Cloning cloudwego/iasm into 'tmp'..."

IASM_TAG="v0.2.0"
rm -rf tmp
if git clone --branch $IASM_TAG --depth 1 git@github.com:cloudwego/iasm.git tmp >/dev/null 2>&1; then
    echo "done"
else
    echo "git clone failed"
    exit 1
fi

rm -rf expr
cp -r tmp/expr .

rm -rf obj
cp -r tmp/obj .

rm -rf x86_64
cp -r tmp/x86_64 .

rm -rf tmp

# rm unused code
rm ./x86_64/assembler*

# replace import
sed -i.bak 's:github.com/cloudwego:github.com/bytedance/sonic/loader/internal:g' ./x86_64/*.go && rm ./x86_64/*.bak

# trim file methods
./trim.py

gofmt -w */*.go

echo "done"
