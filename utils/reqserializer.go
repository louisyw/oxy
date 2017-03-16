//Serializes http.Request objects into text for logging (safely)

package utils

import (
	"net/http"
	"net/http/httputil"
)

func SerializeHttpReq(req *http.Request) string {
	if dump, err := httputil.DumpRequest(req, false); err != nil {
		return err.Error()
	} else {
		return string(dump)
	}
}
