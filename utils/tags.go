/**
 * Copyright 2025 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package utils

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/bytedance/sonic"
	"github.com/fatih/structtag"
)

// TagChanger is used to change the tag of a struct.
//
// WARNING: it is not thread-safe, must be used in a single thread or in initialization phase.
type TagChanger struct {
	types map[reflect.Type]reflect.Type // old -> new
	replacer func(field reflect.StructField, tag reflect.StructTag) reflect.StructTag
}

func NewTagChanger(replacer func(field reflect.StructField, tag reflect.StructTag) reflect.StructTag) *TagChanger {
	return &TagChanger{
		types: make(map[reflect.Type]reflect.Type),
		replacer: replacer,
	}
}

// Replace returns the type whose tag is replaced by replacer.
//
// WARNING: it is not thread-safe, must be used in a single thread or in initialization phase.
func (t *TagChanger) Replace(old reflect.Type) reflect.Type  {
    if t.types == nil {
        t.types = make(map[reflect.Type]reflect.Type)
    }

    if n, ok := t.types[old]; ok {
        return n
    }

    if old.Kind() == reflect.Struct {
        newType := reflect.StructOf(t.createFieldsWithReplacedTags(old))
        t.types[old] = newType
        return newType
    }

    if old.Kind() == reflect.Map {
        nn := t.Replace(old.Elem())
		n := reflect.MapOf(old.Key(), nn)
		t.types[old] = n
        return n
    }

    if old.Kind() == reflect.Slice || old.Kind() == reflect.Array {
       nn := t.Replace(old.Elem())
       n := reflect.SliceOf(nn)
       t.types[old] = n
       return n
    }

    if old.Kind() == reflect.Ptr {
        nn := t.Replace(old.Elem())
        n := reflect.PtrTo(nn)
        t.types[old] = n
        return n
    }

    return old
}

func (t *TagChanger) createFieldsWithReplacedTags(old reflect.Type) []reflect.StructField {
    numField := old.NumField()
    fields := make([]reflect.StructField, numField)

    for i := 0; i < numField; i++ {
        oldField := old.Field(i)
        tag := oldField.Tag

		fieldType := t.Replace(oldField.Type)

        newTag := t.replacer(oldField, tag)

        fields[i] = reflect.StructField{
            Name: oldField.Name,
            Type: fieldType,
            Tag:  newTag,
        }
    }

    return fields
}

// ReplacerAddOmitempty is a replacer that adds "omitempty" to the json tag.
func ReplacerAddOmitempty(field reflect.StructField, tag reflect.StructTag) reflect.StructTag {
	tags, err := structtag.Parse(string(tag))
	if err != nil {
		return tag
	}

	jsonTag, err := tags.Get("json")
	if jsonTag == nil {
		if err := tags.Set(&structtag.Tag{Key: "json", Name: field.Name, Options: []string{"omitempty"}}); err != nil {
			panic(err)
		}
		return reflect.StructTag(tags.String())
	}

	if !jsonTag.HasOption("omitempty") {
		jsonTag.Options = append(jsonTag.Options, "omitempty")
	}

	tags.Set(jsonTag)
	newTag := tags.String()
	return reflect.StructTag(newTag)
}

func ReplaceerAddStringInt64(field reflect.StructField, tag reflect.StructTag) reflect.StructTag {
    typ := field.Type
    if typ.Kind() == reflect.Ptr {
    	typ = typ.Elem()
    }
    if typ.Kind() != reflect.Int64 && typ.Kind() != reflect.Uint64 {
    	return tag
    }

	tags, err := structtag.Parse(string(tag))
	if err != nil {
		return tag
	}

	jsonTag, err := tags.Get("json")
	if jsonTag == nil {
		if err := tags.Set(&structtag.Tag{Key: "json", Name: field.Name, Options: []string{"string"}}); err != nil {
			panic(err)
		}
		return reflect.StructTag(tags.String())
	}

	if !jsonTag.HasOption("string") {
		jsonTag.Options = append(jsonTag.Options, "string")
	}

	tags.Set(jsonTag)
	newTag := tags.String()
	return reflect.StructTag(newTag)
}

// Marshal is used to marshal a struct whose tag is replaced by replacer into json.
func (t *TagChanger) Marshal(v interface{}) ([]byte, error) {
	oldType := reflect.TypeOf(v)
	if oldType.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("v must be a pointer")
	}
	newType := t.Replace(oldType.Elem())
	newObj := reflect.NewAt(newType, unsafe.Pointer(reflect.ValueOf(v).Pointer())).Interface()
	return sonic.Marshal(newObj)
}

// Unmarshal is used to unmarshal json into a struct whose tag is replaced by replacer.
func (t *TagChanger) Unmarshal(data []byte, v interface{}) error {
	oldType := reflect.TypeOf(v)
	if oldType.Kind()!= reflect.Ptr {
		return fmt.Errorf("v must be a pointer")
	}
	newType := t.Replace(oldType.Elem())
	newObj := reflect.NewAt(newType, unsafe.Pointer(reflect.ValueOf(v).Pointer())).Interface()
	return sonic.Unmarshal(data, newObj)
}
