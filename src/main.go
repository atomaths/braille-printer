// Copyright 2012 Braille Printer Team. All rights reserved.
// Use of this source code is governed by the Apache 2.0 license.

package brailleprinter

import (
	"fmt"
	"net/http"
	svg "github.com/ajstarks/svgo"
	brl_ko "github.com/suapapa/go_braille/ko"
	brl_en "github.com/suapapa/go_braille"
	brl_svg "github.com/suapapa/go_braille/svg"
)

func init() {
	http.HandleFunc("/root", rootHandler)
	http.HandleFunc("/braille", brailleHandler)
	http.HandleFunc("/printq/add", printqAddHandler)
	http.HandleFunc("/printq/list", printqListHandler)
	http.HandleFunc("/printq/item", printqItemHandler)
	http.HandleFunc("/printq/update", printqUpdateHandler)
	http.HandleFunc("/", indexHandler)
}

// http://braille-printer.appspot.com/ 첫 페이지 출력
func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// root 로그인 화면 출력
func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "root/root.html")
}

// API: POST /braille
func brailleHandler(w http.ResponseWriter, r *http.Request) {
	src := r.FormValue("input")
	lang := r.FormValue("lang")
	if lang == "" {
		// TODO: lang이 없거나 auto이면 언어 판단해야함
		lang = "ko"
	}
	format := r.FormValue("format")
	if format == "" {
		format = "svg"
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var bStr string

	if lang == "ko" {
		bStr, _ = brl_ko.Encode(src)
	} else if lang == "en" {
		bStr, _ = brl_en.Encode(src)
	}

	if format == "svg" {
		canvas := svg.New(w)
		defer canvas.End()
		brl_svg.DrawPage30(canvas, bStr)
	} else {
		fmt.Fprint(w, bStr)
	}
}
