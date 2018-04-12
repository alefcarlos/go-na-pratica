/*Um programa é executado sempre pelo pacote main

para executar o programa: go run <nomedoarquivo.go>

go run main.go
*/
package main

import (
	//Ao utilizarmos a função Print do pacote fmt ele é automaticamente importado
	"fmt"

	//Ao utilizarmos a função Now() do pacote time, ele é automaticamente importado
	"time"
)

//Função de entrada do programa
func main() {
	/*Essa instrução irá escrever na tela, mas sem pular linha
	  A próxima instrução que escrever na tela, irá aparecer ao lado
	*/
	fmt.Println("Funções")

	//Aqui criamos uma variável com o valor 1993
	anoNascimento := 1993

	//Aqui criamos outra variável com o valor do cálculo da idade a partir do ano de nascimento
	var minhaIdade = calculaIdade(anoNascimento)

	//Exibimos os resultados
	fmt.Printf("Eu nasci em %d e hoje tenho %d anos de idade \n", anoNascimento, minhaIdade)

	//Aqui estamos atualizando o valor da variável anoNascimento, veja que não precisamos mais do símbolo :=,
	//pois a variável já foi inicializada
	anoNascimento = 1992
	minhaIdade = calculaIdade(anoNascimento)
	fmt.Printf("Eu nasci em %d e hoje tenho %d anos de idade \n", anoNascimento, minhaIdade)

	//Aqui atualizamos o valor da variável minhaIdade passando um número inteiro como parâmetro
	minhaIdade = calculaIdade(1987)

	//Note aqui que o valor exibido em nasci em ... não será 1987,
	//pois não atualizamo o valor da variável anoNascimento, simplesmente passamos como parâmetro.
	fmt.Printf("Eu nasci em %d e hoje tenho %d anos de idade \n", anoNascimento, minhaIdade)

}

/*
Funções são blocos de códigos que nos permite manipular parâmetros recebidos
São bastante usadas para separar e reutilizar blocos de códigos

Por exemplo, eu tenho uma função chamado soma, quando eu quiser somar dois números
eu posso sempre chamar essa essa mesma função passando os parâmetros

uma função PRIVADA é definiada assim:

func <nome>(<parâmetros>) <retorno>{
	corpo da função
}

o nome da função PRIVADA começa com letra minúscula

*/
func calculaIdade(anoNascimento int) int {
	//time.Now() retorna uma struct com as informações de data atual
	//O método Year retorna o número do ano
	anoAgora := time.Now().Year()

	//Podemos subtrai duas variáveis do mesmo tipo
	//int - int
	return anoAgora - anoNascimento
}
