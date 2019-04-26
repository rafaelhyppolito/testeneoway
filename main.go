package main

import (
	"github.com/rafaelhyppolito/testeneoway/repo"
	"fmt"
	"github.com/rafaelhyppolito/testeneoway/servico"
	//"html/template"
	"net/http"      // Gerencia URLs e Servidor Web
)

//Funcao que realiza o carregamento do arquivo e sua importação para o banco de dados
func CarregaArquivo(caminho string)  {
	connection := repo.Connect()
	//Limpa a tabela temporária
	repo.ExecSQL(repo.TruncateTmp(), connection)
	//Lê e higieniza os dados do arquivo inserindo-os na tabela temporária
	servico.LerTexto(caminho)
	//Insere na tabela final, convertendo os dados para os formatos corretos
	repo.ExecSQL(repo.InsertFinal(), connection)
	fmt.Println("Arquivo carregado com sucesso!")
	//repo.InsertSQL("INSERT INTO base(cpf,priv,incompleto) VALUES('12345678900',1,0)", repo.Connect())
}

//Funcao para tratar as requisicoes POST e GET
func index(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }
 
    switch r.Method {
    case "GET":     
         http.ServeFile(w, r, "front/index.html")
    case "POST":
        if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err: %v", err)
            return
		}
		arquivo := r.FormValue("arquivo")
		CarregaArquivo(arquivo)
		fmt.Fprintf(w, "Arquivo carregado com sucesso! \r\nAs informações estão no banco de dados!")
    default:
        fmt.Fprintf(w, "Desculpe, apenas métodos GET e POST são suportados.")
    }
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Serviço ativo e ouvindo na porta 8080.")
	http.ListenAndServe(":8080", nil)
}