package main

import (
	"fmt"
	"html/template"
//	"database/sql"  // Pacote Database SQL para realizar Query
	"net/http"      // Gerencia URLs e Servidor Web
//	"text/template" // Gerencia templates
//	_ "github.com/lib/pq" // Driver PostreSQL para Go
)

func index(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("index.html")
	data := map[string]string{
		"Title": "NeoWay :)",
	}
	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Serviço ativo e ouvindo na porta 8080.")
	http.ListenAndServe(":8080", nil)
}