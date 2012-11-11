// Copyright 2012 Braille Printer Team. All rights reserved.
// Use of this source code is governed by the Apache 2.0 license.

package brailleprinter

import (
	"net/http"
	"strings"
	"time"
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
