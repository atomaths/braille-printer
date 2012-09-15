// Copyright 2011 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package helloworld

import (
	"fmt"
	"net/http"
	"github.com/suapapa/go_braille"
)

func init() {
	http.HandleFunc("/", handle)
}

func handle(w http.ResponseWriter, r *http.Request) {
	s := "Hello World"
	fmt.Fprint(w, s + "\n")
	fmt.Fprint(w, braille.Encode(s) + "\n")
}
