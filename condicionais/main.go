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
	fmt.Println("Condicionais")

	/* operadores relacionais:
	== 	é igual a
	!= é diferente de
	> 	maior que
	< 	menor que
	>= 	maior igual a
	<= 	menor igual a
	! 	negação
	*/

	if 3 > 1 {
		fmt.Println("3 é maior do que 1")
	}

	if 1 == 1 {
		fmt.Println("1 é igual a 1")
	}

	if 3 != 2 {
		fmt.Println("3 é diferente de 2")
	}

	hojeEMeuAniversario := true
	// hojeEMeuAniversario := false //descomente aqui para ver a outra mensagem

	if hojeEMeuAniversario {
		fmt.Println("Parabeéns pra você !")
	} else {
		fmt.Println("Volte amanhã, quem sabe é seu dia !")
	}

	if !hojeEMeuAniversario {
		fmt.Println("Hoje não é seu aniversário.")
	}

	/* Operadores lógicos:
	&&
	||
	*/

	if hojeEMeuAniversario && 1 > 0 {
		fmt.Println("Hoje é meu aniversário e 1 é maior do que 0")
	}

}
