/*Um programa é executado sempre pelo pacote main

para executar o programa: go run <nomedoarquivo.go>

go run main.go
*/
package main

//Ao utilizarmos a função Print do pacote fmt ele é automaticamente importado
import (
	"fmt"
)

//Função de entrada do programa
func main() {
	/*Essa instrução irá escrever na tela, mas sem pular linha
	  A próxima instrução que escrever na tela, irá aparecer ao lado
	*/
	fmt.Print("Ola, mundo com Print!")

	/*Essa instrução irá escrever na tela, mas irá pular linha
	  A próxima instrução que escrever na tela, irá aparecer embaixo
	*/
	fmt.Println("Ola, mundo com Println!")

	/*Essa instrução nos permite escrever algo tela utilizando formatadores
	  %s interpreta um parâmetro do tipo string(texto)
	  \n força pular a linha
	*/
	fmt.Printf("Ola, mundo com Printf! meu nome é %s \n", "alef")
}
