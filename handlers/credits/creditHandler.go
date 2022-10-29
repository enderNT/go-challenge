package credits

import (
	"encoding/json"
	"fmt"
	"io"

	// "log"
	"net/http"
	// "strconv"

	"yoFioGOLANG/investments"
)

type invest struct {
	Investment int32
}

type result struct {
	Credit_type_300 int32
	Credit_type_500	int32
	Credit_type_700	int32
	Error 			bool
}

func CreditAssigment (w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		i := &invest{}
		err := json.NewDecoder(r.Body).Decode(i)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "error al decodificar")
			return
		}

		c := investments.Credit{}
		reslt := &result{}
		q1, q2, q3, err := c.Assign(i.Investment)
		if err != nil {
			reslt.Error = true
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			fmt.Printf("Error: %#v\n", err)
			return
		}

		reslt.Credit_type_300 = q1
		reslt.Credit_type_500 = q2
		reslt.Credit_type_700 = q3
		reslt.Error = false
		err = json.NewEncoder(w).Encode(reslt)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "error al codificar")
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "aplication/json")
		return
	}

	io.WriteString(w, "Hi, try the method/verb post :) with the next struct: \n\n{'Investment': 100}")
}
