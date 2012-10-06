// Copyright 2012 Braille Printer Team. All rights reserved.
// Use of this source code is governed by the Apache 2.0 license.

package brailleprinter

import (
	"net/http"
	"bytes"
	"bufio"
//	"appengine"
//	"appengine/datastore"
	"strings"
	svg "github.com/ajstarks/svgo"
	brl_ko "github.com/suapapa/go_braille/ko"
	brl_en "github.com/suapapa/go_braille"
	"log"
)

type PrintQ struct {
	Type string
	Key string
	Origin string
	Result string
}

// API: POST /printq/add
func printqAddHandler(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	src := r.FormValue("input")
	lang := r.FormValue("lang")
	if lang == "" {
		// TODO: lang이 없거나 auto이면 언어 판단해야함
		lang = "ko"
	}

	var bStr string; var bLen int

	if lang == "ko" {
		bStr, bLen = brl_ko.Encode(src)
	} else if lang == "en" {
		bStr, bLen = brl_en.Encode(src)
	}

	buf := new(bytes.Buffer)
	svgw := bufio.NewWriter(buf)

	canvas := svg.New(svgw)
	defer canvas.End()

	drawBrailleStr(canvas, bStr, bLen)

	log.Println(buf)
}
