// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var urlMap = map[string]string{}
var rItems = []rItem{}

type rItem struct {
	URL     string `json:"url"`
	Article string `json:"article"`
}

// Register HTTP handlers that redirect old blog paths to their new locations.
func init() {
	file, e := ioutil.ReadFile("rewrites.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	err := json.Unmarshal(file, &rItems)
	if err != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	for _, v := range rItems {
		urlMap[v.URL] = v.Article
	}

	for p := range urlMap {
		dest := "/" + urlMap[p]
		http.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, dest, http.StatusMovedPermanently)
		})
	}
}
