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
*/

//Pessoa - Modelo de pessoa com descrição customizada
type Pessoa struct {
	Nome  string
	Idade int
}

/*
https://godoc.org/golang.org/x/tools/cmd/stringer

Quando uma strcut implementa o método String(), é possível customizar
o resultado de um print. É utilizado quando usamos método por exemplo Print, Println, Printf
*/
func (p Pessoa) String() string {
	return fmt.Sprintf("Meu nome %s é e tenho %d anos", p.Nome, p.Idade)
}

//PessoaSemString Modelo de pessoa sem descrição customizada
type PessoaSemString struct {
	Nome  string
	Idade int
}

//Função de entrada do programa
func main() {

	/*Aqui vamos criar nossas structs
	Podemos inicializar já passando os valores dos campos na ordem
	*/
	pessoa1 := Pessoa{"Alef", 25}

	//Podemos inicializar e depois setar os valores dos campos
	pessoa2 := Pessoa{}
	pessoa2.Idade = 9
	pessoa2.Nome = "Pedro"

	//Podemos inicializar passando os nomes dos campos
	pessoa3 := PessoaSemString{Nome: "Karol", Idade: 18}

	/*Essa instrução irá escrever na tela, mas irá pular linha
	  A próxima instrução que escrever na tela, irá aparecer embaixo
	*/
	fmt.Println("Quem sou eu ?", pessoa1)

	/*Essa instrução nos permite escrever algo tela utilizando formatadores
	  %v interpreta a struct sem os nomes dos campos
	  %+v interpreta a struct com os nomes dos campos
	  \n força pular a linha
	*/
	fmt.Printf("Quem sou eu ? %v \n", pessoa2)
	fmt.Printf("Quem sou eu ? %v \n", pessoa3)
	fmt.Printf("Quem sou eu ? %+v \n", pessoa3)
}
