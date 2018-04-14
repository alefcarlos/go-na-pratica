/*Um programa é executado sempre pelo pacote main

para executar o programa: go run <nomedoarquivo.go>

go run main.go
*/
package main

import (
	"fmt"
)

var global string

//Função de entrada do programa
func main() {
	global = "Dentro da função main()"
	fmt.Println(global)

	/*Utilizando jeito curto para
	declarar e inicializar uma variável do tipo lógico
	*/
	variavelEscopoMain := true
	fmt.Println(variavelEscopoMain)

	funcao()
	fmt.Println(global)

	//As variáveis em go quando criadas, assumem os valores default, nunca nil

	//uma variável do tipo string é um texto
	//uma variáve do tipo int é um número
	//uma variável do tipo bool é true/false, ou seja, verdadeiro ou falso

	//Criar uma variável, sem inicializá-la explicitamente
	var nome string  // é inicializado com valor ""
	var idade int    // é inicializado com o valor 0
	var falso bool   // é inicilizado com valor false
	var real float32 // É inicialdo com o valor 0

	fmt.Println(nome)
	fmt.Println(idade)
	fmt.Println(falso)
	fmt.Println(real)

	/*Criar uma variável já inicializada
	var variavel int = 1
	Um jeito mais simples de se escrever é utilizando a forma abaixo
	variavel := valor

	O símbolo := cria e inicializa uma variável
	*/
	nomeI := "alef"
	idadeI := 1
	verdadeiro := true
	realNovo := 5.2

	fmt.Println(nomeI)
	fmt.Println(idadeI)
	fmt.Println(verdadeiro)
	fmt.Println(realNovo)
}

func funcao() {
	global = "Dentro da função funcao()"

	// variavelEscopoMain = false // descomente aqui para ver o erro
}
