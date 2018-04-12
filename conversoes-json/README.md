# Lendo e escrevendo JSON

>O Go fornece um pacote padrão para se trabalhar com JSON **`encoding/json`**.

A codificação e decodificação do JSON para os tipos do GO são estremamente rápidos e fáceis. O pacote fornece duas funcões simplistas que é o `Marshal` e `Unmarshal`.

No entando, o pacote também nos fornece as funções `Encoder` e `Decoder` que nos permite maior controle ao ler e gravar fluxos de dados JSON.
