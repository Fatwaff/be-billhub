package controllers

import (
	"fmt"
	"net/http"

	"github.com/Fatwaff/be_ksi"
)

func KSI_Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, be_ksi.SignUpHandler("MONGOSTRING", "db_ksi", r))
		return
	}
}