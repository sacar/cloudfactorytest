package rest

import (
	"fmt"
	"net/http"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a simple rest api. Find the api details in readme file of this project")
}
