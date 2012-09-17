// Copyright 2011 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package helloworld

import (
	"fmt"
	"net/http"
	//brl "github.com/suapapa/go_braille"
	brl_ko "github.com/suapapa/go_braille/ko"
)

func init() {
	http.HandleFunc("/", handle)
}

func handle(w http.ResponseWriter, r *http.Request) {
	s := "오빤 강남 스타일"
	fmt.Fprint(w, s + "\n")
	fmt.Fprint(w, brl_ko.Encode(s) + "\n")
}
