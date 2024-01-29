package httpclient

import (
	"net/http"
	"strings"
)

func joinHeaderValues(header http.Header) map[string]string {
	h := make(map[string]string)
	for k, v := range header {
		h[k] = strings.Join(v, ",")
	}
	return h
}
