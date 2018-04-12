# User inputs

Quando queremos receber informações do usuário através da nossa aplicação chamamos de **user-input**.

Podemos receber essa informação de incontáveis maneiras: formulário, um clique ou até um texto digitado no console.

## O que vamos aprender ?

Aprenderemos a interpretar o que o usuário digita no console da aplicação utilizando a função `Scan()` do pacote `fmt` e também o método `ReadString()` do pacote `bufio`.


## `fmt.Scan()`

>Utilizamos essa função quando queremos ler um ou mais input do usuário.

### Assinatura da função

```go
func Scan(a ...interface{}) (n int, err error) {
	return Fscan(os.Stdin, a...)
}
```

Parâmetros
* `a` ponteiros a receberem valores

Retornos:

* `n` é o número de informações que a função conseguiu ler (cada espaço é uma informação)
* `err` é a informação do erro, caso occora.

Essa função lê os inputs e tenta fazer a conversão para os tipo das variáveis desejadas, caso ocorra algum erro o retorno `err` irá conter as informações.

### Exemplo de usos

Única informação

```go
var variavel int
n, err = fmt.Scan(&variavel)

//variavel terá valor que o usuário digitou
fmt.Print(variavel)
```

Várias informações

```go
var variavel,variavel2, variavel3 int
n, err = fmt.Scan(&variavel, &variavel2, &variavel3)

fmt.Print(variavel1)
fmt.Print(variavel2)
fmt.Print(variavel3)
```

## `reader.ReadString('\n')`

Utilizamos o método `ReadString()` quando queremos ler toda a linha que foi digitada pelo usuário.

## Assinatura do método

```go
func (b *Reader) ReadString(delim byte) (string, error) {
	bytes, err := b.ReadBytes(delim)
	return string(bytes), err
}
```

Parâmetros
* `delim`

Retornos:

* `string` 
* `error` 

Como podemos ver o método `ReadString()` é utilizada a partir de uma referência da struct `Reader`, então primeiros devemos criar um `Reader` para poder utilizá-lo:

```go
reader := bufio.NewReader(os.Stdin)
```

## Exemplo de usos

```go
reader := bufio.NewReader(os.Stdin)
fmt.Println("Qual seu nome completo ?")

nome, err := reader.ReadString('\n')

fmt.Print(nome)
```