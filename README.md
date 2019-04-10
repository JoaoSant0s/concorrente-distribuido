# GoLand

- Instale o GoLand
- Abra o projeto atividade_4 pelo goland
- Verifique se perde a referência de scripts
- Para tentar corrigir, vá em File->Settings-> Go -> GOPATH
- Na Opção "Project GOPATH" coloque o diretório atual. "atividade_4" deve ser sua raiz, para não ter conflito com outras bibliotecas de outras atividades

# concorrente-distribuido

Repositório focado para guardar e trabalhar atividades e projetos em programação concorrente distribuída

Para rodar o exercicio 3, primeiro deleta qlq resquicio de pasta que exista no go/src que involva o exercicio 3 que tu tenha usado antes pra testar la, deleta tudo msm pq ele tenta pegar o primeiro import que encontra, dps pega as pastas primos e shared do atividade3/exercicio e por no go/src do pc que ele pega os imports certinhos.
No caso fica go/src/primos e go/src/shared e abre as duas no vscode ou oq preferir dps só rodar normalmente com go run


Para o Rabbit1m precisa instalar o "ampq"

`
go get github.com/streadway/amqp
`

# Desenvolvimento

É preciso o comando:

`
go get [package path]
`

Em algumas desses plots (normalmente eles pedem a dependência) como por exemplo:

`
go get github.com/ajstarks/svgo
`


Como proposto pelo professor, o rabbitmq do middleware tem refêrencias encontradas aqui: 

- [Rabbitmq](https://www.rabbitmq.com/tutorials/tutorial-one-go.html)
- [go-amqp-example](https://github.com/andreagrandi/go-amqp-example)


`
Não foi possível se conectar ao servidor de mensageria!!: dial tcp [::1]:8080: connectex: Nenhuma conexão pôde ser feita porque a máquina de destino as recusou ativamente.                                                 exit status 1                                                          
`