/*Um programa é executado sempre pelo pacote main

para executar o programa: go run <nomedoarquivo.go>

go run main.go
*/
package main

import (
	//Ao utilizarmos a função NewReader, é automáticamente importado
	"bufio"

	//Ao utilizarmos a função Print do pacote fmt ele é automaticamente importado
	"fmt"
	"os"
)

//Função de entrada do programa
func main() {
	/*Essa instrução irá escrever na tela, mas sem pular linha
	  A próxima instrução que escrever na tela, irá aparecer ao lado
	*/
	fmt.Println("User inputs")

	/*
		Utilizamos a função NewReader do pacote bufio para
		ler um texto que foi informado pelo usuário
	*/
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Qual seu nome completo ?")

	/*Após definirmos nosso reader, temos que utilizar o Método ReadString()
	para ler um input do usuário

	passamos o argumento '\n', para indicar que o reader deve parar de ler quando o usuário teclar enter

	a função panic() força a parada da aplicação indicando que algo não esperado aconteceu

	*/
	nome, err := reader.ReadString('\n')
	if err != nil {
		//Caso ocorra algum erro, devemos parar a execução da aplicação
		panic(err)
	}

	/*
		Criamos uma variável do tipo inteiro que vamos
		armazenar a idade do usuário
	*/
	var idade int
	fmt.Println("Quantos anos você tem ?")

	/*Utilizamos a função Scan() do pacote fmt para
	ler um input do usuário e já convertê-lo para um tipo de dado desejado

	Quando queremos ignorar um retorno, ou seja, essa informação não é relevante, utilizamos o _(underscore)

	No exemplo abaixo, o número de informações retornadas não é relevante, somente o erro

	Para testar o erro:
		Quando perguntar a idade, escreva qualquer coisa que não seja número.
	*/
	_, err = fmt.Scan(&idade)
	if err != nil {
		//Caso ocorra algum erro, devemos parar a execução da aplicação
		panic(err)
	}

	fmt.Printf("Seu nome é %s e você tem %d anos. \n", nome, idade)

	/*A função Scan() pode receber vários parâmetros, cada parâmetro representa uma informação lida.

	Um exemplo é a sequência de números 1 2 3 4 5
	*/
	fmt.Println("Digte 3 números separados por espaço")

	var numero1, numero2, numero3 int
	n, err := fmt.Scan(&numero1, &numero2, &numero3)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Você digitou %d número: %d, %d e %d", n, numero1, numero2, numero3)
}
