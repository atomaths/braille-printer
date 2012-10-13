// Copyright 2012 Braille Printer Team. All rights reserved.
// Use of this source code is governed by the Apache 2.0 license.

package brailleprinter

import (
	"net/http"
	"bytes"
	"fmt"
	"time"
	"strings"
	"encoding/json"
	"appengine"
	"appengine/datastore"
	svg "github.com/ajstarks/svgo"
	brl_ko "github.com/suapapa/go_braille/ko"
	brl_en "github.com/suapapa/go_braille"
)

type PrintQ struct {
	Type string
	Key string
	Origin string
	Result []byte
	Status int
	CTime time.Time
}

// API: POST /printq/add
func printqAddHandler(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var authKey string
	if strings.Contains(r.Referer(), "http://localhost") ||
	   strings.Contains(r.Referer(), "http://braille-printer.appspot.com") {
		authKey = "examplekey"
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	input := r.FormValue("input")
	lang := r.FormValue("lang")
	if lang == "" {
		// TODO: lang이 없거나 auto이면 언어 판단해야함
		lang = "ko"
	}

	label := "label"

	if strings.Contains(input, "\n") {
		label = "paper"
	}

	var bStr string; var bLen int

	if lang == "ko" {
		bStr, bLen = brl_ko.Encode(input)
	} else if lang == "en" {
		bStr, bLen = brl_en.Encode(input)
	}

	buf := bytes.NewBuffer(make([]byte, 24288))

	canvas := svg.New(buf)
	defer canvas.End()
	drawBrailleStr(canvas, bStr, bLen)

	printq := PrintQ{
		Type: label,
		Key: authKey,
		Origin: input,
		Result: buf.Bytes(),
		Status: 0,
		CTime: time.Now(),
	}

	c := appengine.NewContext(r)
	_, err := datastore.Put(c, datastore.NewIncompleteKey(c, "PrintQ", nil), &printq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// API: GET /printq/list
func printqListHandler(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// type, key에 대당하는 query string 가져옴.
	qs, err := parseQueryString(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	label := qs.Get("type")
	if label == "" {
		label = "label"
	}
	key := qs.Get("key")
	if key == "" {
		key = "examplekey"
	}

	// Datastore에 조회할 쿼리 만듬.
	c := appengine.NewContext(r)
	q := datastore.NewQuery("PrintQ").Filter("Key =", key)
	if label != "all" {
		q = q.Filter("Type =", label)
	}
	q = q.Order("CTime").Limit(100)

	// [{"qid":1,"type":"label"},{"qid":2,"type":"paper"}] 형태로 리턴해준다.
	type QList struct {
                Qid int64 `json:"qid"`
                Type string `json:"type"`
        }

        qlist := make([]QList, 0, 100)

	for t := q.Run(c); ; {
		var x PrintQ
		qid, err := t.Next(&x)
		if err == datastore.Done {
			break
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		qlist = append(qlist, QList{Qid: qid.IntID(), Type: x.Type})
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	b, err := json.Marshal(qlist)
	fmt.Fprintf(w, "%s", string(b))
}
