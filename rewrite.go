// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "net/http"

// Register HTTP handlers that redirect old blog paths to their new locations.
func init() {
	for p := range urlMap {
		dest := "/" + urlMap[p]
		http.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, dest, http.StatusMovedPermanently)
		})
	}
}

var urlMap = map[string]string{
	"/2009/05/auto-complete-example-project.html":          "ace",
	"/2009/05/more-autocomplete.html":                      "moreace",
	"/2009/02/auto-complete-on-google-search.html":         "acegsa",
	"/2009/02/debugging-and-windows-service-in-vbnet.html": "debugservice",
	"/2005/05/doomed.html":                                 "doomed",
}
