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
	"reflect"
	"testing"
)

type testUser struct {
	Name    string       `json:"name"`
	Age     int          `json:"age"`
	Email   string       `json:"email,omitempty"`
	Tags    []string     `json:"tags"`
	Profile *testProfile `json:"profile,omitempty"`
}

type testProfile struct {
	Bio     string `json:"bio"`
	Website string `json:"website"`
}

func TestNodeFrom_Struct(t *testing.T) {
	user := testUser{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
		Tags:  []string{"admin", "user"},
		Profile: &testProfile{
			Bio:     "Hello world",
			Website: "https://example.com",
		},
	}

	node, err := NodeFrom(user)
	if err != nil {
		t.Fatalf("NodeFrom failed: %v", err)
	}

	// Verify the node type
	if node.Type() != V_OBJECT {
		t.Errorf("Expected V_OBJECT, got %d", node.Type())
	}

	// Verify fields can be accessed
	name, err := node.Get("name").String()
	if err != nil {
		t.Errorf("Get name failed: %v", err)
	}
	if name != "Alice" {
		t.Errorf("Expected 'Alice', got '%s'", name)
	}

	age, err := node.Get("age").Int64()
	if err != nil {
		t.Errorf("Get age failed: %v", err)
	}
	if age != 30 {
		t.Errorf("Expected 30, got %d", age)
	}

	// Verify nested access
	bio, err := node.Get("profile").Get("bio").String()
	if err != nil {
		t.Errorf("Get profile.bio failed: %v", err)
	}
	if bio != "Hello world" {
		t.Errorf("Expected 'Hello world', got '%s'", bio)
	}

	// Verify array access
	tags := node.Get("tags")
	if tags.Type() != V_ARRAY {
		t.Errorf("Expected V_ARRAY, got %d", tags.Type())
	}
	tag0, _ := tags.Index(0).String()
	if tag0 != "admin" {
		t.Errorf("Expected 'admin', got '%s'", tag0)
	}
}

func TestNodeFrom_Primitives(t *testing.T) {
	// Test string
	strNode, err := NodeFrom("hello")
	if err != nil {
		t.Fatalf("NodeFrom string failed: %v", err)
	}
	if strNode.Type() != V_STRING {
		t.Errorf("Expected V_STRING, got %d", strNode.Type())
	}
	str, _ := strNode.String()
	if str != "hello" {
		t.Errorf("Expected 'hello', got '%s'", str)
	}

	// Test number
	numNode, err := NodeFrom(42)
	if err != nil {
		t.Fatalf("NodeFrom int failed: %v", err)
	}
	if numNode.Type() != V_NUMBER {
		t.Errorf("Expected V_NUMBER, got %d", numNode.Type())
	}
	num, _ := numNode.Int64()
	if num != 42 {
		t.Errorf("Expected 42, got %d", num)
	}

	// Test bool
	boolNode, err := NodeFrom(true)
	if err != nil {
		t.Fatalf("NodeFrom bool failed: %v", err)
	}
	if boolNode.Type() != V_TRUE {
		t.Errorf("Expected V_TRUE, got %d", boolNode.Type())
	}

	// Test nil
	nilNode, err := NodeFrom(nil)
	if err != nil {
		t.Fatalf("NodeFrom nil failed: %v", err)
	}
	if nilNode.Type() != V_NULL {
		t.Errorf("Expected V_NULL, got %d", nilNode.Type())
	}
}

func TestNodeFrom_Map(t *testing.T) {
	m := map[string]interface{}{
		"key1": "value1",
		"key2": 123,
		"key3": true,
	}

	node, err := NodeFrom(m)
	if err != nil {
		t.Fatalf("NodeFrom map failed: %v", err)
	}

	if node.Type() != V_OBJECT {
		t.Errorf("Expected V_OBJECT, got %d", node.Type())
	}

	key1, _ := node.Get("key1").String()
	if key1 != "value1" {
		t.Errorf("Expected 'value1', got '%s'", key1)
	}

	key2, _ := node.Get("key2").Int64()
	if key2 != 123 {
		t.Errorf("Expected 123, got %d", key2)
	}
}

func TestNodeFrom_Slice(t *testing.T) {
	s := []interface{}{1, "two", true, nil}

	node, err := NodeFrom(s)
	if err != nil {
		t.Fatalf("NodeFrom slice failed: %v", err)
	}

	if node.Type() != V_ARRAY {
		t.Errorf("Expected V_ARRAY, got %d", node.Type())
	}

	// Load all to get correct length (lazy parsing)
	if err := node.LoadAll(); err != nil {
		t.Errorf("LoadAll failed: %v", err)
	}

	length, lenErr := node.Len()
	if lenErr != nil {
		t.Errorf("Len failed: %v", lenErr)
	}
	if length != 4 {
		t.Errorf("Expected len 4, got %d", length)
	}
}

func TestNodeFrom_ExistingNode(t *testing.T) {
	// Test that NodeFrom returns existing Node as-is
	original := NewString("test")
	node, err := NodeFrom(original)
	if err != nil {
		t.Fatalf("NodeFrom Node failed: %v", err)
	}

	str, _ := node.String()
	if str != "test" {
		t.Errorf("Expected 'test', got '%s'", str)
	}

	// Test pointer to Node
	ptrNode, err := NodeFrom(&original)
	if err != nil {
		t.Fatalf("NodeFrom *Node failed: %v", err)
	}
	str2, _ := ptrNode.String()
	if str2 != "test" {
		t.Errorf("Expected 'test', got '%s'", str2)
	}

	// Test nil pointer
	var nilPtr *Node
	nilPtrNode, err := NodeFrom(nilPtr)
	if err != nil {
		t.Fatalf("NodeFrom nil *Node failed: %v", err)
	}
	if nilPtrNode.Type() != V_NULL {
		t.Errorf("Expected V_NULL, got %d", nilPtrNode.Type())
	}
}

func TestNode_Unmarshal_Struct(t *testing.T) {
	user := testUser{
		Name:  "Bob",
		Age:   25,
		Email: "bob@example.com",
		Tags:  []string{"developer"},
		Profile: &testProfile{
			Bio:     "Developer",
			Website: "https://bob.dev",
		},
	}

	// Convert to Node
	node, err := NodeFrom(user)
	if err != nil {
		t.Fatalf("NodeFrom failed: %v", err)
	}

	// Unmarshal back to struct
	var result testUser
	err = node.Unmarshal(&result)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	// Verify fields
	if result.Name != user.Name {
		t.Errorf("Name mismatch: expected %s, got %s", user.Name, result.Name)
	}
	if result.Age != user.Age {
		t.Errorf("Age mismatch: expected %d, got %d", user.Age, result.Age)
	}
	if result.Email != user.Email {
		t.Errorf("Email mismatch: expected %s, got %s", user.Email, result.Email)
	}
	if !reflect.DeepEqual(result.Tags, user.Tags) {
		t.Errorf("Tags mismatch: expected %v, got %v", user.Tags, result.Tags)
	}
	if result.Profile == nil || result.Profile.Bio != user.Profile.Bio {
		t.Errorf("Profile.Bio mismatch")
	}
}

func TestNode_Unmarshal_Primitives(t *testing.T) {
	// Test string
	strNode := NewString("hello")
	var str string
	if err := strNode.Unmarshal(&str); err != nil {
		t.Errorf("Unmarshal string failed: %v", err)
	}
	if str != "hello" {
		t.Errorf("Expected 'hello', got '%s'", str)
	}

	// Test number
	numNode := NewNumber("42")
	var num int
	if err := numNode.Unmarshal(&num); err != nil {
		t.Errorf("Unmarshal int failed: %v", err)
	}
	if num != 42 {
		t.Errorf("Expected 42, got %d", num)
	}

	// Test bool
	boolNode := NewBool(true)
	var b bool
	if err := boolNode.Unmarshal(&b); err != nil {
		t.Errorf("Unmarshal bool failed: %v", err)
	}
	if !b {
		t.Errorf("Expected true, got false")
	}
}

func TestNode_Unmarshal_Map(t *testing.T) {
	node, _ := NewParser(`{"a": 1, "b": "two", "c": true}`).Parse()

	var m map[string]interface{}
	if err := node.Unmarshal(&m); err != nil {
		t.Errorf("Unmarshal map failed: %v", err)
	}

	if m["b"] != "two" {
		t.Errorf("Expected 'two', got '%v'", m["b"])
	}
}

func TestNode_Unmarshal_Slice(t *testing.T) {
	node, _ := NewParser(`[1, 2, 3, 4, 5]`).Parse()

	var s []int
	if err := node.Unmarshal(&s); err != nil {
		t.Errorf("Unmarshal slice failed: %v", err)
	}

	if len(s) != 5 {
		t.Errorf("Expected len 5, got %d", len(s))
	}
	if s[2] != 3 {
		t.Errorf("Expected s[2]=3, got %d", s[2])
	}
}

func TestNode_Unmarshal_Nil(t *testing.T) {
	var node *Node
	var result string
	err := node.Unmarshal(&result)
	if err != ErrNotExist {
		t.Errorf("Expected ErrNotExist for nil node, got %v", err)
	}
}

func TestNodeFrom_Unmarshal_RoundTrip(t *testing.T) {
	original := testUser{
		Name:  "Charlie",
		Age:   35,
		Email: "charlie@test.com",
		Tags:  []string{"a", "b", "c"},
		Profile: &testProfile{
			Bio:     "Test bio",
			Website: "https://test.com",
		},
	}

	// Round trip: struct -> Node -> struct
	node, err := NodeFrom(original)
	if err != nil {
		t.Fatalf("NodeFrom failed: %v", err)
	}

	var result testUser
	if err := node.Unmarshal(&result); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if !reflect.DeepEqual(original, result) {
		t.Errorf("Round trip failed:\noriginal: %+v\nresult: %+v", original, result)
	}
}

func TestNodeFrom_Unmarshal_ModifyNode(t *testing.T) {
	original := testUser{
		Name: "Dave",
		Age:  40,
	}

	// Convert to Node
	node, err := NodeFrom(original)
	if err != nil {
		t.Fatalf("NodeFrom failed: %v", err)
	}

	// Modify the node
	node.Set("name", NewString("David"))
	node.Set("age", NewNumber("41"))

	// Unmarshal modified node
	var result testUser
	if err := node.Unmarshal(&result); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if result.Name != "David" {
		t.Errorf("Expected 'David', got '%s'", result.Name)
	}
	if result.Age != 41 {
		t.Errorf("Expected 41, got %d", result.Age)
	}
}

// Benchmark tests
func BenchmarkNodeFrom_Struct(b *testing.B) {
	user := testUser{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
		Tags:  []string{"admin", "user"},
		Profile: &testProfile{
			Bio:     "Hello world",
			Website: "https://example.com",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = NodeFrom(user)
	}
}

func BenchmarkNode_Unmarshal_Struct(b *testing.B) {
	user := testUser{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
		Tags:  []string{"admin", "user"},
		Profile: &testProfile{
			Bio:     "Hello world",
			Website: "https://example.com",
		},
	}
	node, _ := NodeFrom(user)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result testUser
		_ = node.Unmarshal(&result)
	}
}
