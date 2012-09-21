// Copyright 2012 Braille Printer Team. All rights reserved.
// Use of this source code is governed by the Apache 2.0 license.

package brailleprinter

import (
	"net/http"
	"appengine"
	"appengine/datastore"
)

type Root struct {
	Id string
	Pw string
	Using bool
}

func rootInsert(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("root-id")
	pw := r.FormValue("root-pw")
	if id == "" || pw == "" {
		// TODO: id, pw 입력하라고 알려주도록 해야함
		http.ServeFile(w, r, "static/root.html")
		return
	}

	c := appengine.NewContext(r)
	root := Root{
		Id: id,
		Pw: pw,
		Using: true,
	}

	_, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Root", nil), &root)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
