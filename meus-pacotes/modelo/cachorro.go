/*
Para definir um novo pacote:

package <nome>

Um pacote serve para agruparmos uma série de arquivos em comum

Por exemplo, criamos aqui um pacote chamado modelo, nesse pacote estarão agrupados todos
os arquivos de struct
*/
package modelo

//Cachorro Modelo que representa as informações de um cachorro
type Cachorro struct {
	Nome  string
	Idade int
	Raca  string
}

/*Aqui definimos uma struct PRIVADA, ou seja, ela está disponível apenas no pacote modelo
pacotes externos NÃO poderão acessar essa struct
*/
type animal struct {
}
