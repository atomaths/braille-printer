// Copyright 2012 Braille Printer Team. All rights reserved.
// Use of this source code is governed by the Apache 2.0 license.

package brailleprinter

import (
	"fmt"
	"net/http"
	//brl "github.com/suapapa/go_braille"
	svg "github.com/ajstarks/svgo"
	brl_ko "github.com/suapapa/go_braille/ko"
	brl_svg "github.com/suapapa/go_braille/svg"
)

func init() {
	http.HandleFunc("/braille", brailleHandler)
	http.HandleFunc("/root", rootHandler)
	http.HandleFunc("/login", loginHandler)
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
	http.ServeFile(w, r, "root.html")
}

// root 로그인 처리
func loginHandler(w http.ResponseWriter, r *http.Request) {
	rootInsert(w, r)
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
	src := r.FormValue("b-input")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	bStr, bLen := brl_ko.Encode(src)
	fmt.Fprint(w, bStr)

	canvas := svg.New(w)
	defer canvas.End()

	drawBrailleStr(canvas, bStr, bLen)

}
