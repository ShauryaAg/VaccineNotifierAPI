package views

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	t, err := template.ParseFiles(wd + "/templates/password_reset.html")
	if err != nil {
		fmt.Println("error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println("error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
}
