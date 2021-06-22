
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

package testdata

// ffjson: skip
// easyjson:skip
type Book struct {
	BookId  int       `json:"id"`
	BookIds []int     `json:"ids"`
	Title   string    `json:"title"`
	Titles  []string  `json:"titles"`
	Price   float64   `json:"price"`
	Prices  []float64 `json:"prices"`
	Hot     bool      `json:"hot"`
	Hots    []bool    `json:"hots"`
	Author  Author    `json:"author"`
	Authors []Author  `json:"authors"`
	Weights []int     `json:"weights"`
}

// ffjson: skip
// easyjson:skip
type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Male bool   `json:"male"`
}

var book = Book{
	BookId:  12125925,
	BookIds: []int{-2147483648, 2147483647},
	Title:   "未来简史-从智人到智神",
	Titles:  []string{"hello", "world"},
	Price:   40.8,
	Prices:  []float64{-0.1, 0.1},
	Hot:     true,
	Hots:    []bool{true, true, true},
	Author:  author,
	Authors: []Author{author, author, author},
	Weights: nil,
}

var author = Author{
	Name: "json",
	Age:  99,
	Male: true,
}

var data = []byte(`{"id":12125925,"ids":[-2147483648,2147483647],"title":"未来简史-从智人到智神","titles":["hello","world"],"price":40.8,"prices":[-0.1,0.1],"hot":true,"hots":[true,true,true],"author":{"name":"json","age":99,"male":true},"authors":[{"name":"json","age":99,"male":true},{"name":"json","age":99,"male":true},{"name":"json","age":99,"male":true}],"weights":[]}`)