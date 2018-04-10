/*
Para definir um novo pacote:

package <nome>

Um pacote serve para agruparmos uma série de arquivos em comum

Por exemplo, criamos aqui um pacote chamado modelo, nesse pacote estarão agrupados todos
os arquivos de struct
*/
package modelo

//Pessoa - Modelo de pessoa
type Pessoa struct {
	Nome  string
	Idade int
}
