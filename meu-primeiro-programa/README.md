# Meu primeiro programa

## **Olá, mundo !**

A primeira interação de um ser humano com uma nova tecnologia é conhecida como `Olá, mundo`, o famoso hello-world.

Isso é apenas para nos familirizar com o ambiente de denseolvimento.

Geralmente inicializamos um arquivo simples com uma simples instrução: `escreva no console:  Olá, mundo!`

## O que é um software/programa ?

> É uma coleção de instruções que descrevem uma tarefa a ser realizada por um computador *(LAUREANO, 2005, p. 4)*.

## O que vamos aprender ?

Veremos um anatômia da estrutra de um arquivo Go e também as funções que escrevem informações no console da aplicação.


## Arquivo `.go`

Todo arquivo que contém códigos de Golang devem ter a extesão `.go`

```go
package main // <-- Todo arquivo deve pertencer a um pacote. O arquivo de entrada de um programa deve ser chamado main

//Função de entrada do programa
func main() {
  // A sua aplicação começa a ser executada a partir daqui
}

```

> Para executar uma aplicação go digite no console `go run <nome>.go`

## `fmt.Print`

Utilizamos essa função quando queremos apresentar alguma informação simples no console. 

> Caso a próxima instrução seja alguma que escreva algo na tela, o conteúdo aparecerá logo após o último caractere.

### Assinatura da função

```go
func Print(a ...interface{}) (n int, err error) {
	return Fprint(os.Stdout, a...)
}
```

### Exemplos de uso

```go
fmt.Print("Oi")
fmt.Print("Teste", " Continuação..")
```

Resultado esperado

```cmd
OiTeste Continuação..
```

## `fmt.Println`

A mesma funcionalidade que `fmt.Print`, porém adiciona um caractere `\n` no final, que é interpretado como nova linha.

### Assinatura da função

```go
func Println(a ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, a...)
}
```

### Exemplos de uso

```go
fmt.Println("Oi")
fmt.Println("Como vai ?")
fmt.Println("Vou bem...", " E você ?")
```

Resultado esperado

```cmd
Oi
Como vai ?
Vou bem... E você ?

```

## `fmt.Printf`

Utilizamos essa função quando queremos facilitar a escrita na tela utilizando formatadores de tipos. 

Formatadores comuns:

* `%s` utilizada para strings
* `%d` utilizado para inteiros

>Acesse [esse link](https://gobyexample.com/string-formatting) para ver uma lista com vários exemplos

>**DICA**: Para pular linha adicione `\n`.

### Assinatura da função

```go
func Printf(format string, a ...interface{}) (n int, err error) {
	return Fprintf(os.Stdout, format, a...)
}
```

```go
fmt.Printf("\n Oi")
fmt.Printf("Como vai ? \n")
fmt.Printf("Meu nome é %s e tenho %d anos", "Alef", 15)
```

Resultado esperado

```cmd

Oi
Como vai ?
Meu nome é Alef e tenho 15 anos

```

O arquivo [main.go](main.go) contém exemplos funcionais.

O próximo passo é: [As variáveis](/variaveis)