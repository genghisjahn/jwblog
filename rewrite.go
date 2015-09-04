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
	"/2011/04/web-service-minutiae.html":                   "webservice",
	"/2011/01/b-is-for-backup.html":                        "backup",
	"/2010/12/new-laptop.html":                             "newlaptop",
	"/2010/10/looking-forward.html":                        "lookingforward",
	"/2010/10/try-kindle.html":                             "kindle",
	"/2010/10/wrong-major.html":                            "wrongmajor",
	"/2010/08/time-keeps-on-slippin.html":                  "timeslipping",
	"/2010/05/im-on-iphonegraphy-today.html":               "iphoneog",
	"/2010/05/learning-about-twitter.html":                 "twitter",
	"/2010/05/starting-blog.html":                          "startblog",
	"/2010/04/iphone-images.html":                          "iphoneimages",
	"/2010/04/time-to-learn.html":                          "timetolearn",
	"/2010/04/twitter-feeds-for-websites-that-dont.html":   "twitterfeeds",
	"/2010/03/numlock-morse-code.html":                     "morsecode",
	"/2009/05/auto-complete-example-project.html":          "ace",
	"/2009/05/more-autocomplete.html":                      "moreace",
	"/2009/02/auto-complete-on-google-search.html":         "acegsa",
	"/2009/02/debugging-and-windows-service-in-vbnet.html": "debugservice",
	"/2005/05/doomed.html":                                 "doomed",
}
