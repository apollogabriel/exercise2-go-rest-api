package handlers

import (
	"fmt"
	"net/http"
)

func StudentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET students Route"))
		fmt.Println("Hello GET students Route")
	case http.MethodPost:
		w.Write([]byte("Hello POST students Route"))
		fmt.Println("Hello POST students Route")
	case http.MethodPut:
		w.Write([]byte("Hello PUT students Route"))
		fmt.Println("Hello PUT students Route")
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH students Route"))
		fmt.Println("Hello PATCH students Route")
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE students Route"))
		fmt.Println("Hello DELETE students Route")
	}
}
