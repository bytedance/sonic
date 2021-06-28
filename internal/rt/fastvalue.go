/*
 * Copyright 2021 ByteDance Inc.
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

package rt

import (
	"reflect"
	"unsafe"
)

var reflectRtypeItab = findReflectRtypeItab()

const (
	_KindMask    = (1 << 5) - 1
	_DirectIface = 1 << 5
)

type GoType struct {
	nb     uintptr
	ptrd   uintptr
	hash   uint32
	tflags uint8
	align  uint8
	falign uint8
	kflags uint8
	traits unsafe.Pointer
	gcdata *byte
	str    int32
	p      int32
}

func (self *GoType) Size() int {
	return int(self.nb)
}

func (self *GoType) Hash() uint32 {
	return self.hash
}

func (self *GoType) Kind() reflect.Kind {
	return reflect.Kind(self.kflags & _KindMask)
}

func (self *GoType) Pack() (t reflect.Type) {
	(*GoIface)(unsafe.Pointer(&t)).Itab = reflectRtypeItab
	(*GoIface)(unsafe.Pointer(&t)).Value = unsafe.Pointer(self)
	return
}

func (self *GoType) NoPtr() bool {
	return self.ptrd == 0
}

func (self *GoType) Indir() bool {
	return (self.kflags & _DirectIface) == 0
}

func (self *GoType) String() string {
	return self.Pack().String()
}

type GoMap struct {
	Count      int
	Flags      uint8
	B          uint8
	Overflow   uint16
	Hash0      uint32
	Buckets    unsafe.Pointer
	OldBuckets unsafe.Pointer
	Evacuate   uintptr
	Extra      unsafe.Pointer
}

type GoMapIterator struct {
	Key         unsafe.Pointer
	Elem        unsafe.Pointer
	T           *GoMapType
	H           *GoMap
	Buckets     unsafe.Pointer
	Bptr        *unsafe.Pointer
	Overflow    *[]unsafe.Pointer
	OldOverflow *[]unsafe.Pointer
	StartBucket uintptr
	Offset      uint8
	Wrapped     bool
	B           uint8
	I           uint8
	Bucket      uintptr
	CheckBucket uintptr
}

type GoItab struct {
	it unsafe.Pointer
	vt *GoType
	hv uint32
	_  [4]byte
	fn [1]uintptr
}

type GoIface struct {
	Itab  *GoItab
	Value unsafe.Pointer
}

type GoEface struct {
	Type  *GoType
	Value unsafe.Pointer
}

func (self GoEface) Pack() (v interface{}) {
	*(*GoEface)(unsafe.Pointer(&v)) = self
	return
}

type GoPtrType struct {
	GoType
	Elem *GoType
}

type GoMapType struct {
	GoType
	Key        *GoType
	Elem       *GoType
	Bucket     *GoType
	Hasher     func(unsafe.Pointer, uintptr) uintptr
	KeySize    uint8
	ElemSize   uint8
	BucketSize uint16
	Flags      uint32
}

func (self *GoMapType) IndirectElem() bool {
	return self.Flags&2 != 0
}

type GoStructType struct {
	GoType
	Pkg    *byte
	Fields []GoStructField
}

type GoStructField struct {
	Name     *byte
	Type     *GoType
	OffEmbed uintptr
}

type GoSlice struct {
	Ptr unsafe.Pointer
	Len int
	Cap int
}

type GoString struct {
	Ptr unsafe.Pointer
	Len int
}

func PtrElem(t *GoType) *GoType {
	if t.Kind() != reflect.Ptr {
		panic("not a pointer: " + t.String())
	} else {
		return (*GoPtrType)(unsafe.Pointer(t)).Elem
	}
}

func MapType(t *GoType) *GoMapType {
	if t.Kind() != reflect.Map {
		panic("not a map: " + t.String())
	} else {
		return (*GoMapType)(unsafe.Pointer(t))
	}
}

func UnpackType(t reflect.Type) *GoType {
	return (*GoType)((*GoIface)(unsafe.Pointer(&t)).Value)
}

func UnpackEface(v interface{}) GoEface {
	return *(*GoEface)(unsafe.Pointer(&v))
}

func findReflectRtypeItab() *GoItab {
	v := reflect.TypeOf(struct{}{})
	return (*GoIface)(unsafe.Pointer(&v)).Itab
}
