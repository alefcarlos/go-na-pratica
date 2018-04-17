# As variáveis

## O que é uma variável ?

> É uma posição nomeada de memória, que é usada para guardar um valor que pode ser modificado pelo programa (LAUREANO, 2005, p. 12)

Criamos variáveis para guardar informações que podem ser reutilizadas durante a execução do programa.

## Tipos de dados

Para cada informação que queremos guardar temos um tipo apropriado. Os tipos mais comuns são:

* Inteiro
* Real
* Texto
* Lógico

Utilizamos o tipo de dado inteiro quando queremos armazenar números sem casas decimais, tanto positivo quanto negativo, e quando queremos utilizar números fracionados utilizamos o tipo de dado real.


Utilizamos o tipo de dado texto quando queremos armazenar uma sequência de alfanuméricos, por exemplo seu nome.

Utilizamos o tipo de dado lógico quando queremos armazenar valores lógicos(decisões), os valores podem somente ser `TRUE` ou `FALSE`.

## As variáveis em Go

Quando vamos utilizar variáveis em programas de computadores primeiro temos de verificar como elas funcionam em cada linguagem. Utilizando Go os tipos das variáveis são:

* int
* float32
* string
* bool

## Utilizando variáveis em Go

Antes de utilizarmos uma variável nós temos de declará-la, ou seja, criá-la.

A sintaxe de declaração de uma variável em Go é:

```gol
var <nome> <tipo>
```

Exemplo de criação de uma variável do tipo inteiro:

```go
var idade int
```

> DICA: Quando criamos variáveis no Go elas já são inicializada com um valor padrão. O padrão de int é `0`, o padrão de string é `""`, o padrão de bool é `false` e o float32 é iniciado com `0`.

Após realizada a declaração da variável, nós podemos atribuir valores a ela.

```go
idade = 24
```

E para utilizar, exemplo utilizando `fmt.Print()`

```go
fmt.Print(idade);
```

Resultado esperado:

```cmd
24
```

## Declarar e inicializar

Nós também podemos já declarar e inicializar uma variável, pois as vezes já sabemos o valor que vamos utilizar, por exemplo: um teste fixo, o retono de uma função...

Exemplo:

```go
var nome string = "Zé"
var total int = soma(1, 4)
var multiplicacao int = 1 * 1

fmt.Print(nome);
```

Temos outro jeito encurtado para criar e inicilizar uma variável:

```go
nome := "Maria"
```

O símbolo `:=` indica: crie e inicialize o valor da varável nome. 

> Não precisamos informar o tipo do dado quando estamos utilizando dessa maneira.

## Escopo

O escopo de uma variável é sempre onde ela foi declarada.

O que isso significa ? Só podemos utilizar uma variável no bloco em que nós a criamos.

Entenda com o exemplo abaixo:

```go
package main

func main(){
    var nome string = "Zé"
    fmt.Print(nome);

    // Não se preocupe com essa linha, vamos aprender sobre 
    // funções na próxima etapa
    funcao();
    fmt.Print(nome);
}

func funcao(){
    nome = "Outro Nome" //Aqui ocorrerá erro ao tentar executar seu programa. Dizendo que nome não foi encontrado.
}
```

Criamos a variável `nome` na funçao main, então ela só fica restrita ao bloco da main. Caso tentemos utilizar dentro de outra função, causará erro.

Agora veja esse outro exemplo:

```go
package main

//Tudo o que for digitado abaixo da declaração do package
// é considerado global
var nome string

func main(){
    nome = "Zé"
    fmt.Print(nome);

    // Não se preocupe com essa linha, vamos aprender sobre 
    // funções na próxima etapa
    funcao();
    fmt.Print(nome);
}

func funcao(){
    nome = "Outro Nome" //Aqui não ocorrerá mais erro
}
```

Podemos definir `variáveis globais` declarando-as abaixo do package, assim elas ficam disponíveis para todo o pacote.

## Nomeando variáveis

É sempre legal manter um padrão quando criamos uma variável. 

* Use nomes que faça sentido a variável
* Use palavras curtas

E quando existirem palavras compostas utilize o padrão *lowerCamelCase*, por exemplo dataAniversario.

> *lowerCamelCase:* a primeira letra da primeira palavra sempre minúscula e as seguites com a primeira letra maíscula.

O arquivo [main.go](main.go) contém exemplos funcionais.

O próximo passo é: [Funções](/funcoes)