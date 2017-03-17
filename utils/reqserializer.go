//Serializes http.Request objects into text for logging (safely)

package utils

import (
	"encoding/json"
	"net/http"
)

func SerializeHttpReq(req *http.Request) string {
	dump, _ := json.Marshal(req)
	return string(dump)
}
