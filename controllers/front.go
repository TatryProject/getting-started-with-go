package controllers

import (
	"encoding/json"
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()

	/*
	 * Go handlers /users and /users/ routes differently
	 * so we specify both.
	 */
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON(data interface{}, w http.ResponseWriter) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
