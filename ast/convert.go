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

package ast

// NodeFrom converts any Go value to an ast.Node.
// It first marshals the value to JSON bytes, then parses the JSON into a Node.
//
// This is a convenience function that avoids manual serialization when you need
// to convert a Go struct or other value into an AST for further manipulation.
//
// Unlike NewAny(), the returned Node has a fully-parsed structure that supports
// all Get/Set operations.
//
// Example:
//
//	type User struct {
//	    Name string `json:"name"`
//	    Age  int    `json:"age"`
//	}
//	node, err := ast.NodeFrom(User{Name: "Alice", Age: 30})
//	name, _ := node.Get("name").String()  // "Alice"
func NodeFrom(v interface{}) (Node, error) {
	// Fast path: if v is already a Node, return it
	switch n := v.(type) {
	case Node:
		return n, nil
	case *Node:
		if n == nil {
			return NewNull(), nil
		}
		return *n, nil
	}

	// Marshal the value to JSON bytes using platform-specific implementation
	data, err := marshalValue(v)
	if err != nil {
		return Node{}, err
	}

	// Parse JSON into Node - use string() for safe copy since encoder may reuse buffer
	node, parseErr := NewParser(string(data)).Parse()
	if parseErr != 0 {
		return Node{}, parseErr
	}
	return node, nil
}

// Unmarshal decodes the Node's JSON representation into the value pointed to by v.
// It first encodes the Node to JSON bytes, then unmarshals into v.
//
// This is the inverse of NodeFrom and provides a convenient way to convert
// an ast.Node (or part of it) back into a Go struct or other value.
//
// Example:
//
//	node, _ := sonic.Get(jsonBytes, "user")
//	var user User
//	err := node.Unmarshal(&user)
func (self *Node) Unmarshal(v interface{}) error {
	if self == nil {
		return ErrNotExist
	}

	// Encode Node to JSON bytes
	data, err := self.MarshalJSON()
	if err != nil {
		return err
	}

	// Unmarshal JSON into v using platform-specific implementation
	return unmarshalValue(data, v)
}

// UnmarshalInto is an alias for Unmarshal for API consistency.
// Deprecated: Use Unmarshal instead.
func (self *Node) UnmarshalInto(v interface{}) error {
	return self.Unmarshal(v)
}
