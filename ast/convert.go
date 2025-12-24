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

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// NodeFrom converts any Go value to a fully-parsed Node.
// Unlike NewAny(), the returned Node supports all Get/Set operations.
func NodeFrom(v interface{}) (Node, error) {
	switch n := v.(type) {
	case Node:
		return n, nil
	case *Node:
		if n == nil {
			return NewNull(), nil
		}
		return *n, nil
	}

	data, err := marshalValue(v)
	if err != nil {
		return Node{}, err
	}

	node, perr := NewParser(string(data)).Parse()
	if perr != 0 {
		return Node{}, perr
	}
	return node, nil
}

// Unmarshal decodes the Node into the value pointed to by v.
// v must be a non-nil pointer.
func (self *Node) Unmarshal(v interface{}) error {
	if self == nil {
		return ErrNotExist
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return fmt.Errorf("Unmarshal requires a non-nil pointer")
	}

	return nodeToValue(self, rv.Elem())
}

func nodeToValue(node *Node, v reflect.Value) error {
	if node == nil || !node.Exists() {
		return nil
	}

	if err := node.Check(); err != nil {
		return err
	}

	if v.Kind() == reflect.Ptr {
		if node.Type() == V_NULL {
			v.Set(reflect.Zero(v.Type()))
			return nil
		}
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}
		return nodeToValue(node, v.Elem())
	}

	if v.Kind() == reflect.Interface {
		val, err := nodeToInterface(node)
		if err != nil {
			return err
		}
		if val != nil {
			v.Set(reflect.ValueOf(val))
		}
		return nil
	}

	nodeType := node.Type()

	switch nodeType {
	case V_NULL:
		v.Set(reflect.Zero(v.Type()))
		return nil

	case V_TRUE, V_FALSE:
		if v.Kind() == reflect.Bool {
			v.SetBool(nodeType == V_TRUE)
			return nil
		}
		return fmt.Errorf("cannot unmarshal bool into %v", v.Type())

	case V_STRING:
		str, err := node.String()
		if err != nil {
			return err
		}
		switch v.Kind() {
		case reflect.String:
			v.SetString(str)
			return nil
		case reflect.Slice:
			// Handle []byte - decode base64
			if v.Type().Elem().Kind() == reflect.Uint8 {
				decoded, err := base64.StdEncoding.DecodeString(str)
				if err != nil {
					return err
				}
				v.SetBytes(decoded)
				return nil
			}
		}
		return fmt.Errorf("cannot unmarshal string into %v", v.Type())

	case V_NUMBER:
		return unmarshalNumber(node, v)

	case V_ARRAY:
		return unmarshalArray(node, v)

	case V_OBJECT:
		return unmarshalObject(node, v)

	default:
		return fmt.Errorf("unknown node type: %d", nodeType)
	}
}

func nodeToInterface(node *Node) (interface{}, error) {
	switch node.Type() {
	case V_NULL:
		return nil, nil
	case V_TRUE:
		return true, nil
	case V_FALSE:
		return false, nil
	case V_STRING:
		return node.String()
	case V_NUMBER:
		return node.Float64()
	case V_ARRAY:
		return node.Array()
	case V_OBJECT:
		return node.Map()
	default:
		return nil, fmt.Errorf("unknown node type: %d", node.Type())
	}
}

func unmarshalNumber(node *Node, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := node.Int64()
		if err != nil {
			return err
		}
		v.SetInt(n)
		return nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		n, err := node.Int64()
		if err != nil {
			return err
		}
		if n < 0 {
			return fmt.Errorf("cannot unmarshal negative number %d into %v", n, v.Type())
		}
		v.SetUint(uint64(n))
		return nil

	case reflect.Float32, reflect.Float64:
		n, err := node.Float64()
		if err != nil {
			return err
		}
		v.SetFloat(n)
		return nil

	case reflect.String:
		num, err := node.Number()
		if err != nil {
			return err
		}
		v.SetString(string(num))
		return nil

	default:
		if v.Type() == reflect.TypeOf(json.Number("")) {
			num, err := node.Number()
			if err != nil {
				return err
			}
			v.Set(reflect.ValueOf(num))
			return nil
		}
		return fmt.Errorf("cannot unmarshal number into %v", v.Type())
	}
}

func unmarshalArray(node *Node, v reflect.Value) error {
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return fmt.Errorf("cannot unmarshal array into %v", v.Type())
	}

	if err := node.LoadAll(); err != nil {
		return err
	}

	length, err := node.Len()
	if err != nil {
		return err
	}

	if v.Kind() == reflect.Slice {
		if v.IsNil() || v.Cap() < length {
			v.Set(reflect.MakeSlice(v.Type(), length, length))
		} else {
			v.SetLen(length)
		}
	}

	for i := 0; i < length; i++ {
		elem := node.Index(i)
		if i < v.Len() {
			if err := nodeToValue(elem, v.Index(i)); err != nil {
				return err
			}
		}
	}

	return nil
}

func unmarshalObject(node *Node, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Map:
		return unmarshalMap(node, v)
	case reflect.Struct:
		return unmarshalStruct(node, v)
	default:
		return fmt.Errorf("cannot unmarshal object into %v", v.Type())
	}
}

func unmarshalMap(node *Node, v reflect.Value) error {
	if v.Type().Key().Kind() != reflect.String {
		return fmt.Errorf("map key must be string")
	}

	if v.IsNil() {
		v.Set(reflect.MakeMap(v.Type()))
	}

	if err := node.LoadAll(); err != nil {
		return err
	}

	var iterErr error
	node.ForEach(func(path Sequence, child *Node) bool {
		if path.Key == nil {
			return true
		}

		elemVal := reflect.New(v.Type().Elem()).Elem()
		if err := nodeToValue(child, elemVal); err != nil {
			iterErr = err
			return false
		}
		v.SetMapIndex(reflect.ValueOf(*path.Key), elemVal)
		return true
	})

	return iterErr
}

func unmarshalStruct(node *Node, v reflect.Value) error {
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldVal := v.Field(i)

		if !field.IsExported() {
			continue
		}

		tag := field.Tag.Get("json")
		if tag == "-" {
			continue
		}

		name, opts := parseJSONTag(tag)
		if name == "" {
			name = field.Name
		}
		hasStringTag := strings.Contains(opts, "string")

		if field.Anonymous && fieldVal.Kind() == reflect.Struct {
			if err := unmarshalStruct(node, fieldVal); err != nil {
				return err
			}
			continue
		}
		if field.Anonymous && fieldVal.Kind() == reflect.Ptr && fieldVal.Type().Elem().Kind() == reflect.Struct {
			if fieldVal.IsNil() {
				fieldVal.Set(reflect.New(fieldVal.Type().Elem()))
			}
			if err := unmarshalStruct(node, fieldVal.Elem()); err != nil {
				return err
			}
			continue
		}

		child := node.Get(name)
		if !child.Exists() {
			continue
		}

		if hasStringTag && child.Type() == V_STRING {
			str, err := child.String()
			if err != nil {
				return err
			}
			tempNode, perr := NewParser(str).Parse()
			if perr != 0 {
				return fmt.Errorf("failed to parse string-tagged value: %v", perr)
			}
			child = &tempNode
		}

		if err := nodeToValue(child, fieldVal); err != nil {
			return err
		}
	}

	return nil
}

func parseJSONTag(tag string) (name string, opts string) {
	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx], tag[idx+1:]
	}
	return tag, ""
}

// UnmarshalInto is an alias for Unmarshal for API consistency.
// Deprecated: Use Unmarshal instead.
func (self *Node) UnmarshalInto(v interface{}) error {
	return self.Unmarshal(v)
}
