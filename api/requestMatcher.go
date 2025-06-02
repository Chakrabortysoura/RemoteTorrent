package api

import "net/http"

func NewRequestMatcher() *http.ServeMux {
	var requestMatcher = http.NewServeMux() // New servermux for reques matching

	//URL handle functions for the various api endpoints
	requestMatcher.HandleFunc("POST /configure/save-folder/", ConfigureDestFolder)
	requestMatcher.HandleFunc("POST /configure/password/", ConfigurePassword)
	requestMatcher.HandleFunc("GET /settings/", Settings)

	return requestMatcher
}
