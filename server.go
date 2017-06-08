package lazyhacker

import (
	"net/http"
)

func init() {
	http.Handle("/.well-known//", http.FileServer(http.Dir(".")))
}

