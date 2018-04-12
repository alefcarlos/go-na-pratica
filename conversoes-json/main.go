/*

Para codificar dados JSON, o pacote encoding/json nos fornece a função Marshal, que possuí a seguinte assinatura:

func Marshal(v interface{}) ([]byte, error)

A função recebe uma interface como parametro de entrada, ou seja, a função aceita qualquer tipo de estrutura como entrada.

Já no retorno, a função nos retorna uma tupla de ([]byte, error) um estilo de retorno muito comum em Go.

Vamos criar um exemplo onde criamos uma struct, e convertemos a mesma em JSON.
*/

package main

func main() {

}

type Pessoa struct {
	Nome  string
	Idade int
}
