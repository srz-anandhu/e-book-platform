package api

import (
	"fmt"
	"net/http"
)

func DemoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "called demo handler")
}
