/*Um programa é executado sempre pelo pacote main

para executar o programa: go run <nomedoarquivo.go>

go run main.go
*/
package main

import (
	//Ao utilizarmos a função Print do pacote fmt ele é automaticamente importado
	"fmt"

	//Ao utilizarmos a struct modelo.Pessoa, é automaticamente será importado
	"github.com/alefcarlos/go-na-pratica/meus-pacotes/funcao"

	//Ao utilizamos o a função funcao.CalculaIdade, é automaticamente importado
	"github.com/alefcarlos/go-na-pratica/meus-pacotes/modelo"
)

//Função de entrada do programa
func main() {
	/*Essa instrução irá escrever na tela, mas sem pular linha
	  A próxima instrução que escrever na tela, irá aparecer ao lado
	*/
	fmt.Println("Teste de pacotes.")

	/*Para utilizar uma struct de outro package	devemos escrever <nomedopacote>.<struct>
	 */
	pessoa := modelo.Pessoa{}
	pessoa.Nome = "Alef"

	fmt.Printf("pessoa: %+v", pessoa)

	cachorro := modelo.Cachorro{Nome: "Belinha", Idade: 8, Raca: "Poodle"}
	fmt.Printf("pessoa: %+v", cachorro)

	/*A linha abaixo não irá compilar, pois tem o erro: undefined: modelo.Animal
	o nome da nossa strcut é animal(com a minúsculo)
	*/
	// animal := modelo.Animal{} //descomente essa linha para ver o erro

	/*A linh abaixo também não irá compilar, pois tem o erro: cannot refer to unexported name modelo.animal, ou seja,
	não conseguimos utilizar uma struct que não seja publica fora do mesmo pacote
	*/
	//animal2 := modelo.animal{}//descomente essa linha para ver o erro

	/*Podemos utilizar a mesma lógica para funções públicas, ou seja,
	funções de outros pacotes que comecem com a letra maíscula
	*/
	var idade = funcao.CalculaIdade(1993)
	fmt.Printf("idade é igual: %d", idade)

	/*Se tentarmos utilizar a função funcao.soma teremos o erro: cannot refer to unexported name funcao.soma*/
	// soma := funcao.soma(1, 2) //descomente essa linha para ver o erro

}
