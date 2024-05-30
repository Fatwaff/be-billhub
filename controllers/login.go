package controllers

import (
	"fmt"
	"net/http"

	"github.com/Fatwaff/be_ksi"
)

func KSI_Login(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the main request.
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, be_ksi.LogInHandler("PASETOPRIVATEKEY", "MONGOSTRING", "db_ksi", r))
		return
	}

}