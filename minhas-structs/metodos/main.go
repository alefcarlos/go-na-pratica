/*Um programa é executado sempre pelo pacote main

para executar o programa: go run <nomedoarquivo.go>

go run main.go
*/
package main

import "fmt"

/*
 O golang não tem paradgima definido, mas também não é orientado a objeto, então não temos o conceito de Classes aqui.
 Temos o que chamado de Struct, que nada mais é do que uma coleção de campos.

 criado uma struct assim:


 type <Name> struct{
	 <campo> <tipo>
 }

 Toda struct que comece coma letra maiúscula deve ter um comentário no padrão:
 <Name> Essa struct é isso..

 Podemos criar métodos para nossas structs:

 func <recebedor> <Nome>() <retorno> {

 }
*/

//Pessoa - Modelo de pessoa com descrição customizada
type Pessoa struct {
	Nome  string
	Idade int
}

//Falar método que retorna
func (*Pessoa) Falar() string {
	return "Eu sei falar"
}

//Andar método
func (*Pessoa) Andar() string {
	return "Eu sei andar"
}

//Contar -
func (*Pessoa) Contar(numero int) string {
	return fmt.Sprintf("Eu sei contar, você me passou %d", numero)
}

//Função de entrada do programa
func main() {

	/*Aqui vamos criar nossas structs
	Podemos instânciar já passando os valores dos campos na ordem
	*/
	pessoa1 := Pessoa{"Alef", 25}

	/*Essa instrução irá escrever na tela, mas irá pular linha
	  A próxima instrução que escrever na tela, irá aparecer embaixo
	*/
	fmt.Println("Quem sou eu ?", pessoa1)

	fmt.Println("Falar: ", pessoa1.Falar())
	fmt.Println("Andar: ", pessoa1.Andar())
	fmt.Println("Contar: ", pessoa1.Contar(10))
}
