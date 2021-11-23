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

package decoder

import (
	"encoding"
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
)
 
 type T struct {
	 S string `json:"s"`
 }
 
 func (self T) MarshalJSON() ([]byte, error) {
	 return []byte(self.S), nil
 }
 
 func (self T) MarshalText() ([]byte, error) {
	 return []byte(self.S), nil
 }
 
 type T2 struct {
	 A string `json:"a"`
 }
 
 func (self *T2) MarshalJSON() ([]byte, error) {
	 return []byte(self.A), nil
 }
 
 func (self *T2) MarshalText() ([]byte, error) {
	 return []byte(self.A), nil
 }
 
 func getRandomEface() interface{} {
	 switch rand.Intn(1) {
		 case 1: return "hello"
		 case 0: return &T2{A: "\"hello world\""}
		 default: return nil
	 }
 }
 
 func getRandomJsonMarshaler() json.Marshaler {
	 switch rand.Intn(2) {
		 case 1: return T{S: "\"hello world\""}
		 case 0: return &T2{A: "\"hello world\""}
		 default: return nil
	 }
 }
 
 func getRandomTextMarshaler() encoding.TextMarshaler {
	 switch rand.Intn(2) {
		 case 1: return T{S: "\"hello world\""}
		 case 0: return &T2{A: "\"hello world\""}
		 default: return nil
	 }
 }
 
 
 
 
//   func TestABIError(t *testing.T) {
// 	 cases :=  []struct {
// 		 name string
// 		 in string
// 		 flag uint64
// 		 obj interface{}
// 		 err error
// 	 }{
// 		//  {"error_stackoverflow", "{}", 9, &map[string]string{}, stackOverflow},
// 	 }
 
// 	 for i, c := range cases {
// 		 dc := NewDecoder(c.in)
// 		 dc.f = c.flag
// 		 err := dc.Decode(c.obj)
// 		 if c.err == nil && err != nil {
// 			 t.Fatal(i, c.name, err)
// 		 }else if c.err != nil && err == nil {
// 			 t.Fatal(i, c.name)
// 		 }else if c.err.Error() != err.Error() {
// 			 fmt.Printf("exp:%v\ngot:%v\n", c.err.Error(), err.Error())
// 			 t.Fatal(i, c.name)
// 		 }
// 	 }
//   }
 
 
  func TestABIOps(t *testing.T) {
	 cases :=  []struct {
		 name string
		 in string
		 obj interface{}
		 err error
	 }{
		 {"slice", "[0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15]",&[]int{}, nil},
	 }
 
	 for i, c := range cases {
		 dc := NewDecoder(c.in)
		 err := dc.Decode(c.obj)
		 if c.err == nil && err != nil {
			 t.Fatal(i, c.name, err)
		 }else if c.err != nil {
		    if err == nil {
			    t.Fatal(i, c.name)
			}else if c.err.Error() != err.Error() {
				fmt.Printf("exp:%v\ngot:%v\n", c.err.Error(), err.Error())
				t.Fatal(i, c.name)
			}
		}
	 }
  }

func BenchmarkSlice(b *testing.B) {
	src := "[0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15]"
	err := NewDecoder(src).Decode(&[]int{})
	if err != nil {
		b.Fatal(err)
	}
	b.SetBytes(int64(len(src)))
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		_ = NewDecoder(src).Decode(&[]int{})
	}
}