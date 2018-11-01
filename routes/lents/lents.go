package lents

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func LentsPostHandler(w http.ResponseWriter, r *http.Request) {
	request, err := httputil.DumpRequest(r, true)
	if err != nil {
		return
	}

	fmt.Fprintf(w, string(request))
}
