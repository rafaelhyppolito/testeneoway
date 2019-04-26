package main

import (
	"github.com/rafaelhyppolito/testeneoway/repo"
	"fmt"
	"github.com/rafaelhyppolito/testeneoway/servico"
	"html/template"
	"net/http"      // Gerencia URLs e Servidor Web
)

func index(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("index.html")
	data := map[string]string{
		"Title": "NeoWay :)",
	}
	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, data)
}

//Funcao que realiza o carregamento do arquivo e sua importação para o banco de dados
func CarregaArquivo()  {
	connection := repo.Connect()
	//Limpa a tabela temporária
	repo.ExecSQL(repo.TruncateTmp(), connection)
	//Lê e higieniza os dados do arquivo inserindo-os na tabela temporária
	servico.LerTexto("base_teste.txt")
	//Insere na tabela final, convertendo os dados para os formatos corretos
	repo.ExecSQL(repo.InsertFinal(), connection)
	fmt.Println("Arquivo carregado com sucesso!")
	//repo.InsertSQL("INSERT INTO base(cpf,priv,incompleto) VALUES('12345678900',1,0)", repo.Connect())
}

func main() {

	http.HandleFunc("/", index)
	fmt.Println("Serviço ativo e ouvindo na porta 8080.")
	http.ListenAndServe(":8080", nil)

	//CarregaArquivo()
}