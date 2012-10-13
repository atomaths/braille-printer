// Copyright 2012 Braille Printer Team. All rights reserved.
// Use of this source code is governed by the Apache 2.0 license.

package brailleprinter

import (
	"net/http"
	"net/url"
	"errors"
)

func parseQueryString(r *http.Request) (url.Values, error) {
	rawURL := "http://"
	if r.URL.IsAbs() {
		rawURL = r.URL.Scheme + "://"
	}
	rawURL += r.Host + r.RequestURI

	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, errors.New("Bad Request")
	}
	return u.Query(), nil
}
