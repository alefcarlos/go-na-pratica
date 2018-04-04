/*Um programa é executado sempre pelo pacote main

para executar o programa: go run <nomedoarquivo.go>

go run main.go
*/
package main

import (
	"fmt"
)

//Função de entrada do programa
func main() {

	//As variáveis em go quando instânciadas, assumem os valores default, nunca nil

	//Instânciar variável, sem inicializá-la explicitamente
	var nome string // é inicializado com valor ""
	var idade int   // é inicializado com o valor 0
	var falso bool  // é inicilizado com valor false

	fmt.Println(nome)
	fmt.Println(idade)
	fmt.Println(falso)

	//Instânciar variável já inicializada
	nomeI := "alef"
	idadeI := 1
	verdadeiro := true

	fmt.Println(nomeI)
	fmt.Println(idadeI)
	fmt.Println(verdadeiro)
}
