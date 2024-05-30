package controllers

import (
	"fmt"
	"net/http"

	"github.com/Fatwaff/be_ksi"
)

func KSI_User(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, be_ksi.GetProfileHandler("PASETOPUBLICKEY", "MONGOSTRING", "db_ksi", r))
}
