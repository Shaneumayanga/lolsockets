package lolsockets

import "net/http"

type checkOrigin func(r *http.Request) bool

var (
	CheckOrigin checkOrigin = func(r *http.Request) bool {
		return true
	}
)
