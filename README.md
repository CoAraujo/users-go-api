# users-go-api

## Descrição
* Projeto para consultar usuários e criar usuários. Criação ocorrerá de forma assíncrona com mensagens sendo enviadas para filas de uma mensageria (ActiveMQ) onde serão processadas assíncronamente por algum consumidor.
* Este projeto está feito em conjunto com outro projeto chamado [users-go-processor](https://github.com/CoAraujo/users-go-processor)
* O docker-compose foi comentado a fim de rodar os dois projetos em conjunto. Caso queira rodar apenas esse, basta descomentar.

## Requisitos Mínimos
* [Go 1.12+](https://golang.org/)
* [Docker](https://www.docker.com/)

## Tecnologias utilizadas
* [Go 1.12+](https://golang.org/)
* [MongoDB](https://www.mongodb.com/)
* [ActiveMQ](https://activemq.apache.org/)
* [Docker](https://www.docker.com/)
* [Docker-compose](https://docs.docker.com/compose/)
* [Echo Framework](https://echo.labstack.com/)
* [Go Mod](https://blog.golang.org/using-go-modules)

## Instalação
1. Baixe o repositório como arquivo zip ou faça um clone;
2. Descompacte os arquivos em seu computador;
3. Abra a pasta decompactada
4. Execute `make up`
5. Aguarde até a stack inteira estar deployada.
6. Acesse o ActiveMQ (www.localhost:8161) para visualizar as mensagens que foram enviadas para a fila.


## Chamadas

Request [GET User]:

`curl -X GET \
  http://localhost:8081/users/{userID} \
  -H 'Content-Type: application/json'
`

Response:

```javascript
{
   "_id":"123",
   "email":"emailteste",
   "username":"usernameTeste",
   "fullName":"fullnameTeste",
   "gender":"genderTeste",
   "status":"statusTeste",
   "birthDate":"birthdateTeste",
   "phones":{
      "phone":"phoneTeste",
      "cellphone":"cellphoneTeste",
      "ddd_cellphone":"21",
      "mobile_phone_confirmed":true
   },
   "clientId":"clientTeste"
}
```

***

Request [CREATE User]:

`curl -X POST \
  http://localhost:8081/users/ \
  -H 'Content-Type: application/json' \
  -d '{
   "_id":"123",
   "email":"emailteste",
   "username":"usernameTeste",
   "fullName":"fullnameTeste",
   "gender":"genderTeste",
   "status":"statusTeste",
   "birthDate":"birthdateTeste",
   "phones":{
      "phone":"phoneTeste",
      "cellphone":"cellphoneTeste",
      "ddd_cellphone":"21",
      "mobile_phone_confirmed":true
   },
   "clientId":"clientTeste"
}'
`

Response:

```javascript
{
   "_id":"123",
   "email":"emailteste",
   "username":"usernameTeste",
   "fullName":"fullnameTeste",
   "gender":"genderTeste",
   "status":"statusTeste",
   "birthDate":"birthdateTeste",
   "phones":{
      "phone":"phoneTeste",
      "cellphone":"cellphoneTeste",
      "ddd_cellphone":"21",
      "mobile_phone_confirmed":true
   },
   "clientId":"clientTeste"
}
```

***

## Collection Postman
* Importar a [collection](https://www.getpostman.com/collections/0535c574b7e864a6baca)

## Arquitetura de Solução
TODO

## Dúvidas?
`Caso tenha dúvidas ou precise de suporte, mande um email para rafacoaraujo@gmail.com`
