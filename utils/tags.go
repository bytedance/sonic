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

func (t *TagChanger) Replace(old reflect.Type) reflect.Type  {
    if t.types == nil {
        t.types = make(map[reflect.Type]reflect.Type)
    }

    // 如果已经处理过该类型，直接返回
    if n, ok := t.types[old]; ok {
        return n
    }

    // 处理结构体类型
    if old.Kind() == reflect.Struct {
        // 创建一个新的结构体类型
        newType := reflect.StructOf(t.createFieldsWithReplacedTags(old))
        t.types[old] = newType
        return newType
    }

    // 处理map类型
    if old.Kind() == reflect.Map {
        // 递归处理map的value类型
        nn := t.Replace(old.Elem())
		n := reflect.MapOf(old.Key(), nn)
		t.types[old] = n
        return n
    }

    // 处理slice/array类型
    if old.Kind() == reflect.Slice || old.Kind() == reflect.Array {
        // 递归处理元素类型
       nn := t.Replace(old.Elem())
       n := reflect.SliceOf(nn)
       t.types[old] = n
       return n
    }

    // 处理指针类型
    if old.Kind() == reflect.Ptr {
        // 递归处理指针指向的类型
        nn := t.Replace(old.Elem())
        n := reflect.PtrTo(nn)
        t.types[old] = n
        return n
    }

    // 其他类型直接返回
    return old
}

func (t *TagChanger) createFieldsWithReplacedTags(old reflect.Type) []reflect.StructField {
    numField := old.NumField()
    fields := make([]reflect.StructField, numField)

    for i := 0; i < numField; i++ {
        oldField := old.Field(i)
        tag := oldField.Tag

        // 递归处理处理字段类型
		fieldType := t.Replace(oldField.Type)

        // 使用自定义的replace函数替换tag
        newTag := t.replacer(oldField, tag)

        fields[i] = reflect.StructField{
            Name: oldField.Name,
            Type: fieldType,
            Tag:  newTag,
        }
    }

    return fields
}


func ReplacerAddOmitemptyfunc(field reflect.StructField, tag reflect.StructTag) reflect.StructTag {
	tags, err := structtag.Parse(string(tag))
	if err != nil {
		return tag
	}

	// 查找 json tag
	jsonTag, err := tags.Get("json")
	if err != nil {
		return reflect.StructTag(`json:",omitempty"`)
	}

	// 修改 json tag 的值
	if !jsonTag.HasOption("omitempty") {
		jsonTag.Options = append(jsonTag.Options, "omitempty")
	}

	// 重新构建 tag
	tags.Set(jsonTag)
	newTag := tags.String()
	return reflect.StructTag(newTag)
}

func MarshalWithTagChanger(t *TagChanger, v interface{}) ([]byte, error) {
	// 先替换tag
	oldType := reflect.TypeOf(v)
	if oldType.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("v must be a pointer")
	}
	newType := t.Replace(oldType)
	newObj := reflect.NewAt(newType, unsafe.Pointer(reflect.ValueOf(v).UnsafeAddr())).Interface()

	// 再序列化
	return sonic.Marshal(newObj)
}

func UnmarshalWithTagChanger(t *TagChanger, data []byte, v interface{}) error {
	// 先替换tag
	oldType := reflect.TypeOf(v)
	if oldType.Kind()!= reflect.Ptr {
		return fmt.Errorf("v must be a pointer")
	}
	newType := t.Replace(oldType)
	newObj := reflect.NewAt(newType, unsafe.Pointer(reflect.ValueOf(v).UnsafeAddr())).Interface()
	// 再反序列化
	return sonic.Unmarshal(data, newObj)
}
