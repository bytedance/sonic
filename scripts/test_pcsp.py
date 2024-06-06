#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import os
import glob
import re

MOCKED_FILES=[]

def make_output_file(stub_file: str) -> str:
    fpath = os.path.splitext(stub_file)[0]
    return fpath + '_text_amd64.go'

def run_cmd(cmd :str):
    print(cmd)
    if os.system(cmd):
        clear_files()
        print ("Failed to run cmd: %s"%(cmd))
        exit(1)

def make_mock_file(prefix :str) -> str:
    asm = []
    # read the asm binary
    with open(prefix + '_text_amd64.go') as fp:
         asm.extend(fp.read().splitlines())

    # dump the mocked function
    fpath = prefix + '_mock_text_amd64_test.go'
    with open(fpath, 'w') as fp:
        for line in asm:
            # replace stubr
            if "var _text_" in line:
                line = line.replace("var _text_", "var _mock_text_")
            
            # check all non-sp instructions
            subsp = "subq"  in line and ", %rsp" in line
            addsp = "addq"  in line and ", %rsp" in line
            movbp = "movq"  in line and ", %rbp" in line
            popq  = "popq"  in line
            pushq = "pushq" in line
            ret   = "retq"  in line and "0xc3,"  in line
            leaq  = "leaq"  in line and ", %rsp" in line
        
            # replace all non-sp instructions
            if "//" in line and not subsp and not addsp and not popq and not pushq and not ret and not leaq and not movbp:
                newline = re.sub(r"0x\w{2}", "0xcc", line.split("//")[0]) + "//" + line.split("//")[1]
            else:
                newline = line
            print(newline, file = fp)

    print(f"Mocked text file generated: {fpath}")
    return fpath

def clear_files():
    for fpath in MOCKED_FILES:
        run_cmd(f"rm -vrf {fpath}")
    
    run_cmd("rm -vrf ./internal/native/sse/traceback_test.go")
    run_cmd("rm -vrf ./internal/native/avx2/traceback_test.go")

def main():
    TRACE_TEST_FILE = "./internal/native/traceback_test.mock_tmpl"
    pattern = "*_text_amd64.go"

    # generate mocked function files
    global MOCKED_FILES
    for dir in [ "./internal/native/avx2/",  "./internal/native/sse/"]:
        for filepath in glob.glob(os.path.join(dir, pattern)):
            prefix = filepath.replace("_text_amd64.go", "")
            print(prefix)
            MOCKED_FILES.append(make_mock_file(prefix))

    # generate mocked trace test file
    run_cmd("sed -e 's/{{PACKAGE}}/sse/g' %s > ./internal/native/sse/traceback_test.go" %TRACE_TEST_FILE)
    run_cmd("sed -e 's/{{PACKAGE}}/avx2/g' %s > ./internal/native/avx2/traceback_test.go" %TRACE_TEST_FILE)

    # test the pcsp with Golang
    run_cmd("go version")
    run_cmd("go test -v -run=TestTraceback ./... > test.log")
    run_cmd("grep -q \"runtime: \" test.log && exit 1 || exit 0")

    clear_files()


if __name__ == '__main__':
    main()
