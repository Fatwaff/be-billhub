package controllers

import (
	"fmt"
	"net/http"

	"github.com/Fatwaff/be_ksi"
)

func KSI_Billboard(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, be_ksi.TambahBillboardHandler("PASETOPUBLICKEY", "MONGOSTRING", "db_ksi", r))
		return
	}
	if r.Method == http.MethodPut {
		fmt.Fprintf(w, be_ksi.EditBillboardHandler("PASETOPUBLICKEY", "MONGOSTRING", "db_ksi", r))
		return
	}
	if r.Method == http.MethodDelete {
		fmt.Fprintf(w, be_ksi.HapusBillboardHandler("PASETOPUBLICKEY", "MONGOSTRING", "db_ksi", r))
		return
	}
	// Set CORS headers for the main request.
	fmt.Fprintf(w, be_ksi.GetBillboarHandler("MONGOSTRING", "db_ksi", r))

}
