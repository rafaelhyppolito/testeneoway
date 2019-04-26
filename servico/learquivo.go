package servico

import (
	"strings"
	//"fmt"
	"bufio"
	"os"
)

//Verifica se existe espaço duplo ao longo da String e retorna TRUE ou FALSE
func existeespacoduplo(texto string) bool {
	var ant, atu int
    for _, v := range texto {
		atu = int(v)
            if v == 32 {
				if ant == atu{
					return true
				}
			}
		ant = int(v)
	}
	return false
}

//Remove todos os espaços duplos do texto e retorna com os separadores csv (;)
func removeespacosduplo(texto string) string {
	for existeespacoduplo(texto){
		texto = strings.ReplaceAll(texto,"  ", " ")
	}	
	texto = strings.ReplaceAll(texto,",", ".")
	texto = strings.ReplaceAll(texto," ", ",")
	return texto
}

// Funcao que le o conteudo do arquivo e retorna um slice the string com todas as linhas do arquivo
func LerTexto(caminhoDoArquivo string) ([]string, error) {
	// Abre o arquivo
	arquivo, err := os.Open(caminhoDoArquivo)
	// Caso tenha encontrado algum erro ao tentar abrir o arquivo retorne o erro encontrado
	if err != nil {
		return nil, err
	}
	// Garante que o arquivo sera fechado apos o uso
	defer arquivo.Close()

	// Cria um scanner que le cada linha do arquivo
	var linhas []string
	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		linhas = append(linhas, removeespacosduplo(scanner.Text()))
	}

	// Retorna as linhas lidas e um erro se ocorrer algum erro no scanner
	//fmt.Printf("Linhas: %s\n\r",linhas)
	return linhas, scanner.Err()
}