package customresponse

import "net/http"

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondwithJSON(w, code, map[string]string{"message": msg})
}
