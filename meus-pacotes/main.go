/*Um programa é executado sempre pelo pacote main

para executar o programa: go run <nomedoarquivo.go>

go run main.go
*/
package main

import (
	//Ao utilizarmos a função Print do pacote fmt ele é automaticamente importado
	"fmt"

	//Ao utilizarmos a struct modelo.Pessoa, automáticamente será importado
	"github.com/alefcarlos/go-na-pratica/meu-pacotes/modelo"
)

//Função de entrada do programa
func main() {
	/*Essa instrução irá escrever na tela, mas sem pular linha
	  A próxima instrução que escrever na tela, irá aparecer ao lado
	*/
	fmt.Println("Teste de pacotes.")

	/*Para utilizar uma struct de outro package
	devemos utilizar o <nomedopacote>.<struct>
	*/
	pessoa := modelo.Pessoa{}
	pessoa.Nome = "Alef"

	fmt.Printf("pessoa: %+v", pessoa)

	cachorro := modelo.Cachorro{Nome: "Belinha", Idade: 8, Raca: "Poodle"}
	fmt.Printf("pessoa: %+v", cachorro)

	//A linha abaixo não irá compilar, pois tem o erro: undefined: modelo.Animal
	// animal := modelo.Animal{} //descomente essa linha para ver o erro

	/*A linh abaixo também não irá compila, pois tem o erro: cannot refer to unexported name modelo.animal, ou seja,
	não conseguimos utilizar uma struct que não seja publica fora do mesmo pacote
	*/
	//animal2 := modelo.animal{}//descomente essa linha para ver o erro
}
