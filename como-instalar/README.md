# Como instalar

## Windows

1. Acessar URL de download [aqui](https://golang.org/dl/)
2. Clicar em Microsoft Windows

Após relizar o download, instalar normalmente. Utilizar a instalação padrão.

> O instalador irá instalar o Go na pasta `C:\Go`

## Mac OS X

1. Acessar URL de download [aqui](https://golang.org/dl/) (Pode-se instalar também pelo [homebrew](http://brewformulas.org/Go))
2. Clicar em Apple MacOS

Após relizar o download, instalar normalmente. Utilizar a instalação padrão.

> O instalador irá instalar o Go na pasta `/usr/local/go`

# GOPATH e GOROOT

`GOROOT` é o caminho da pasta de instalação do Go.

> O instalador  automáticamente criará essa variável de ambiente. Por padrão o valor é `C:\Go`

`GOPATH` é o caminho de onde fica seu workspace, ou seja, onde os seus códigos fontes e bibliotecas estão localizadas. Por padrão é `%USERPROFILE%\go`.

> `%USERPROFILE%` é o caminho da pasta do seu usuário logado. Por exemplo, o seu usuário é jose, então a pasta será `C:\Users\jose`

## Testar instalação

1. Acesse seu `GOPATH(%USERPROFILE%\go)` e crie a pasta `src` e a acesse
2. Crie uma nova pasta chamada github.com e a acesse
3. Crie uma nova pasta chamada alefcarlos e a acesse
4. Clone esse projeto em sua máquina `git clone https://github.com/alefcarlos/go-na-pratica.git`

Execute o arquivo main.go.

> Para excutar uma aplicação utilize o comando `go run <nomearquivo>.go`

```cmd
cd go-na-pratica/como-instalar
go run main.go
```

O resultado esperado é:

```cmd 
Teste de instalação do go!
```

O próximo passo é: [Workspaces](/workspaces)