# Meu primeiro programa

## **Olá, mundo !**

A primeira interação de um ser humano com uma nova tecnologia é conhecida como `Olá, mundo`, o famoso hello-world.

Isso é apenas para nos familirizar com o ambiente de denseolvimento.

Geralmente inicializamos um arquivo simples com uma simples instrução: `escreva no console:  Olá, mundo!`

## O que é um software/programa ?

> Software é uma sequência de instruções escritas para serem interpretadas por um computador com o objetivo de executar tarefas específicas. [Fonte](https://www.significados.com.br/software/)

## O que vamos aprender ?

Aprenderemos as funções que escrevam informações no console da aplicação.

## `fmt.Print`

Utilizamos essa função quando queremos apresentar alguma informação simples no console. 

> Caso seja utilizada alguma instrução que escreva algo na tela, será escrito logo a pós o último caractere.

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

### Assinatura da função

```go
func Printf(format string, a ...interface{}) (n int, err error) {
	return Fprintf(os.Stdout, format, a...)
}
```

O próximo passo é: [As variáveis](/variaveis)