// Copyright 2012 Braille Printer Team. All rights reserved.
// Use of this source code is governed by the Apache 2.0 license.

package brailleprinter

import (
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
	http.HandleFunc("/", indexHandler)
}

// http://braille-printer.appspot.com/ 첫 페이지 출력
func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
	/*
	s := "동해물과"
	fmt.Fprint(w, "<!DOCTYPE html><html><head></head><body>")
	fmt.Fprint(w, s + "<br>")
	fmt.Fprint(w, braille.Encode(s) + "\n")
	fmt.Fprint(w, "</body></html>")
	*/
}

// root 로그인 화면 출력
func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "root/root.html")
}

func drawBrailleStr(canvas *svg.SVG, bStr string, bLen int) {
	dot := 10
	margin := 3

	cw := bLen*(dot*2) + (bLen+1)*margin
	ch := margin*2 + dot*4
	canvas.Start(cw, ch)

	x := margin
	for _, c := range bStr {
		if c & 0x2800 != 0x2800 {
			continue
		}
		brl_svg.Draw(canvas, c, x, margin, dot)
		x += dot * 2
		x += margin
	}
}

func brailleHandler(w http.ResponseWriter, r *http.Request) {
	src := r.FormValue("input")
	lang := r.FormValue("lang")
	if lang == "" {
		// TODO: lang이 없거나 auto이면 언어 판단해야함
		lang = "ko"
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var bStr string; var bLen int

	if lang == "ko" {
		bStr, bLen = brl_ko.Encode(src)
	} else if lang == "en" {
		bStr, bLen = brl_en.Encode(src)
	}

	canvas := svg.New(w)
	defer canvas.End()

	drawBrailleStr(canvas, bStr, bLen)
}
