package controllers

import (
	"fmt"
	"net/http"

	"github.com/Fatwaff/be_ksi"
)

func KSI_Sewa(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, be_ksi.SewaHandler("PASETOPUBLICKEY", "MONGOSTRING", "db_ksi", r))
		return
	}
	if r.Method == http.MethodPut {
		fmt.Fprintf(w, be_ksi.EditSewaHandler("PASETOPUBLICKEY", "MONGOSTRING", "db_ksi", r))
		return
	}
	if r.Method == http.MethodDelete {
		fmt.Fprintf(w, be_ksi.HapusSewaHandler("PASETOPUBLICKEY", "MONGOSTRING", "db_ksi", r))
		return
	}
	// Set CORS headers for the main request.
	fmt.Fprintf(w, be_ksi.GetSewaHandler("PASETOPUBLICKEY", "MONGOSTRING", "db_ksi", r))

}
