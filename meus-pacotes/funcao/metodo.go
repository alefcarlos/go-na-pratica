/*
Para definir um novo pacote:

package <nome>

Um pacote serve para agruparmos uma série de arquivos em comum

Por exemplo, criamos aqui um pacote chamado modelo, nesse pacote estarão agrupados todos
os arquivos de struct
*/
package funcap

import "time"

/*
Funções são blocos de códigos que nos permite manipular parâmetros recebidos
São bastante usadas para separar e reutilizar blocos de códigos

Por exemplo, eu tenho uma função chamado soma, quando eu quiser somar dois números
eu posso sempre chamar esse método passando os parâmetros

uma função PÚBLICA é definiada assim:

//<Nome> essa função faz isso
func <Nome>(<parâmetros>) <retorno>{
	corpo da função
}

o nome da função pública começa com letra MAIÚSCULA

Uma função pública pode ser usada por pacotes externos
*/

//CalculaIdade cálcula quantos anos se passaram desde uma data.
func CalculaIdade(anoNascimento int) int {
	//time.Now() retorna uma struct com as informações de data atual
	//O método Year retorna o número do ano
	anoAgora := time.Now().Year()

	//Podemos subtrai duas variáveis do mesmo tipo
	//int - int
	return anoAgora - anoNascimento
}

/*
Funções são blocos de códigos que nos permite manipular parâmetros recebidos
São bastante usadas para separar e reutilizar blocos de códigos

Por exemplo, eu tenho uma função chamado soma, quando eu quiser somar dois números
eu posso sempre chamar esse método passando os parâmetros

uma função PRIVADA é definiada assim:

func <nome>(<parâmetros>) <retorno>{
	corpo da função
}

o nome da função PRIVADA começa com letra minúscula
uma função privada não pode ser utilizada por pacotes externos

Uma caracterisca legal aqui é que os dois parâmetros da função são do tipo inteiro, então
podemos utilizar essa maneira de escrita:

<parametro1>, <parametro2>, <parametro3> [tipo]
*/
func soma(numero1, numero2 int) int {
	return numero1 + numero2
}
