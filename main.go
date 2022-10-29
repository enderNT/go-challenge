package main

import (
	"net/http"
	"yoFioGOLANG/handlers/credits"
)

func main () {
	http.HandleFunc("/credit-assigment", credits.CreditAssigment)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
