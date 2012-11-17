// Copyright 2012 Braille Printer Team. All rights reserved.
// Use of this source code is governed by the Apache 2.0 license.

package brailleprinter

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"appengine"
	"appengine/datastore"
)

type ClientKey struct {
	Id        string
	PrinterID string
	Used      bool      `datastore:",noindex"`
	CTime     time.Time `datastore:",noindex"`
	UsedTime  time.Time `datastore:",noindex"`
}

// GET /qrcode
// Requirement login as admin previlege.
func qrCodeHandler(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "root/qrcode.html")
}

// POST /qrcode/generate-key
func qrCodeKeyGenerateHandler(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	c := appengine.NewContext(r)
	var tmpKey string

	// Check for duplicated clientKey
	for {
		tmpKey = genKey()
		qry := datastore.NewQuery("ClientKey").Filter("Id =", tmpKey).Filter("Used =", false)
		cnt, _ := qry.Count(c)
		if cnt == 0 {
			break
		}
	}

	clientKey := ClientKey{
		Id: tmpKey,
		PrinterID: "label_printer_0",
		Used: false,
		CTime: time.Now(),
	}

	_, err := datastore.Put(c, datastore.NewIncompleteKey(c, "ClientKey", nil), &clientKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, tmpKey)
}
