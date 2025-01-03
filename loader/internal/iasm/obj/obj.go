//
// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package obj

import (
	"fmt"
	"io"
	"os"
)

// Format represents the saved binary file format.
type Format int

const (
	// ELF indicates it should be saved as an ELF binary.
	ELF Format = iota

	// MachO indicates it should be saved as a Mach-O binary.
	MachO
)

var formatTab = [...]func(w io.Writer, code []byte, base uint64, entry uint64) error{
	ELF:   nil,
	MachO: assembleMachO,
}

var formatNames = map[Format]string{
	ELF:   "ELF",
	MachO: "Mach-O",
}

// String returns the name of a specified format.
func (self Format) String() string {
	if v, ok := formatNames[self]; ok {
		return v
	} else {
		return fmt.Sprintf("Format(%d)", self)
	}
}

// Write assembles a binary executable.
func (self Format) Write(w io.Writer, code []byte, base uint64, entry uint64) error {
	if self >= 0 && int(self) < len(formatTab) && formatTab[self] != nil {
		return formatTab[self](w, code, base, entry)
	} else {
		return fmt.Errorf("unsupported format: %s", self)
	}
}

// Generate generates a binary executable file from the specified code.
func (self Format) Generate(name string, code []byte, base uint64, entry uint64) error {
	var fp *os.File
	var err error

	/* create the output file */
	if fp, err = os.Create(name); err != nil {
		return err
	}

	/* generate the code */
	if err = self.Write(fp, code, base, entry); err != nil {
		_ = fp.Close()
		_ = os.Remove(name)
		return err
	}

	/* close the file and make it executable */
	_ = fp.Close()
	_ = os.Chmod(name, 0755)
	return nil
}
