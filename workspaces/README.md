# Workspaces

> Geralmente utilizamos um único workspace contendo todos os nossos projetos.

Em nosso `GOPATH` deve conter 3 pastas:

> Se você ainda não sabe o que é `GOPATH` [leia aqui](/como-instalar)

* src contém nossos códigos fontes
* pkg conté os códigos das bibliotecas
* bin contém executáveis utilizados pelo Go

> Quando você acaba de instalar o Go, a pasta src não é criada, então se não existir, sinta-se a vontade de criá-la

## src

Aqui tem uma característca interessante do Go, todos os projetos são separados por repositório.

Primeiramente temos uma pasta com o nome do repositório utilizado, por exemplo, o `go-na-pratica` utiliza o Github como repostiório, então os códigos fontes estarão dentro da pasta github.com.

> %GOPATH%/src/github.com/alefcarlos/go-na-pratica

> É sempre legal você também manter seus projetos assim, caso você utlize o Bitbucket, cria uma pasta bitbucket.com, por exemplo.

## Alterar caminho do workspace

Caso você não queira utiliza o caminho padrão de workspace, geralmente em **[Windows]** `C:\Users\seu-usuario\go` **[Mac]** `\Users\jose` , você pode alterar o valor da variável de ambiente `GOPATH`

Próximo passo é: [Visual Code + Plugin Go](/visual-code)