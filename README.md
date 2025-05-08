# Desafio Multithreading

## Module Path Update
The module path has been updated from `github.com/riccruzdev/desafiomulti` to `desafiomulti` to make it easier for users to clone and run the project without import path issues.

# instructions to test it
Por favor rode cmd/server/main.go CEP(arg of the cep para testar no formato 99999999 -  este argumento não é obrigatório, se não passar nada, ele irá buscar o cep 01153000)

O código está preparado para retornar o resultado da API mais rápida, ou seja, ele irá retornar o resultado da API que responder primeiro.

Além disso, irá retornar um timeout caso as duas APIs demorem mais de 1 segundo para responder.

Após isso, você pode jogar com os tempos de resposta das APIs para testar o código. Há 2 comentários no código para descomentar e comentar os tempos de espera das APIs. Assim pode escolher qual quer que retorne primeiro, ou se quer que ambas levem mais de 1 segundo.

Acredito que isto seja o bastante para você testar o código.

Obrigado.

Ric