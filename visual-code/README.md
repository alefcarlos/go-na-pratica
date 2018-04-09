# Visual Code + Plugin Go

> Desde quando aprendi Go, eu utilizo o Visual Code como IDE para codificação de arquivos .go

O download pode ser feito por esse [link](https://code.visualstudio.com/Download) e está disponível para diversas plataformas.

Após efetuada a instalação será necessário instalar nosso plugin salvador, acesse o link para download [aqui](https://marketplace.visualstudio.com/items?itemName=lukehoban.Go) e clique em Install.

Quando terminar de instalar o plugin o Visual Code deve ser reiniciado e então uma mensagem irá aparecer, perguntando se você gostaria de intalar todas as ferramentas, clique em Yes to all.

Caso a mensagem não apareça, tecle `Ctrl + Shift + P` para abrir os comandos e então digite: `Go: Install`, selecione a opção `Go: Install/Update Tools` para a instalação manual.

> `DICA`: Caso você mude seu GOPATH, o Visual Code irá perguntar novamente.

Esse plugin instala 16 ferramentas essenciais para programação, como debugger, validador de código entre outros.

> Agora abra o arquivo como-instalar/main.go no Visual Code e mude a linha `fmt.Print("Teste de instalação do go!")` para `fmt.Print("Testando com Visual Code + Plugin!");`

Ao salvar o arquivo, você notará que o `;` sumiu automaticamente, esse é um dos recursos desse plugin ;)

> O Go não utiliza `;` para indicar fim de uma determinada instrução.

Execute a aplicação utilizando o comando `go run main.go`, o resultado espereado é `Testando com Visual Code + Plugin!`.

Próximo passo: [Meu Primeiro Programa](/meu-primeiro-programa)