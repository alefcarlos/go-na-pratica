# Como instalar

## Windows

1. Acessar URL de download [aqui](https://golang.org/dl/)
2. Clicar em Microsoft Windows

Após relizar o download, instalar normalmente. Utilizar a instalação padrão.

> O instalador irá instalar o Go na pasta `C:\Go`

## Testar instalação

Execute o arquivo main.go.

> Para excutar uma aplicação utilize o comando `go run <nomearquivo>.go`

```cmd
cd como-instalar
go run main.go
```

O resultado esperado é:

```cmd 
Teste de instalação do go!
```

# GOPATH e GOROOT

`GORROT` é o caminho da pasta de instalação do Go.

> O instalador irá automáticamente criará essa variável de ambiente. Por padrão o valor é `C:\Go`

`GOPATH` é o caminho de onde fica seu workspace, ou seja, onde os seus códigos fontes e bibliotecas estão localizadas. Por padrão é `%USERPROFILE%\go`.

> `%USERPROFILE%` é o caminho da pasta do seu usuário logado. Por exemplo, o seu usuário é jose, então a pasta será `C:\Users\jose`