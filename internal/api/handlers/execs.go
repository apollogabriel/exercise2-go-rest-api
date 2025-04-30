package handlers

import (
	"fmt"
	"net/http"
)

func ExecsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET execs Route"))
		fmt.Println("Hello GET execs Route")
	case http.MethodPost:
		w.Write([]byte("Hello POST execs Route"))
		fmt.Println("Hello POST execs Route")
	case http.MethodPut:
		w.Write([]byte("Hello PUT execs Route"))
		fmt.Println("Hello PUT execs Route")
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH execs Route"))
		fmt.Println("Hello PATCH execs Route")
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE execs Route"))
		fmt.Println("Hello DELETE execs Route")
	}
}
