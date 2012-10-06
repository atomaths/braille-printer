// Copyright 2012 Braille Printer Team. All rights reserved.
// Use of this source code is governed by the Apache 2.0 license.

package brailleprinter

import (
	"fmt"
	"net/http"
//	"appengine"
//	"appengine/datastore"
	"strings"
)

type PrintQ struct {
	Type string
	Key string
	Origin string
	Result string
}

func printqAddHandler(w http.ResponseWriter, r *http.Request) {
	if strings.ToUpper(r.Method) != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(w, "ok")
}
