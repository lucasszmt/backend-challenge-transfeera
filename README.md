# backend-challenge-transfeera
Este repositório corresponde ao desafio de Backend da transfeera, para rodar o projeto é necessaria a versão
[1.19 da linguagem GO](https://go.dev/dl/) bem como [Docker](https://docs.docker.com/get-docker/)


Para rodar o projeto localmente basta rodar o seguinte comando
``` 
$ make init
```
Isso fará com que subam containers tanto para o banco quanto para o microserviço, bem como rodará 
as migrações iniciais para criação de tabelas e preenchimento de dados

E para rodar os testes basta apenas rodar o comando `make coverage_tests`
## Endpoints
### Criação de recebedores(receivers)
Deverá ser feita uma requisição do tipo POST para o endpoint `localhost:8000/api/v1/receiver` com o body contendo
os dados `name, email, doc, pix_key_type e pixkey` onde `pix_key_type` deve ser do tipo `{"cpf", "cnpj", "random_key", "email", "phone"}` e 
o campo doc é correpondente a `CPF ou CNPJ`
```
curl --location --request POST 'localhost:8000/api/v1/receiver' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Lucas Szeremeta",
    "email": "lucasszmt@gmail.com",
    "doc": "084.125.359-52",
    "pix_key_type": "cpf",
    "pix_key": "084.125359-52"
}'
```

### Update de um recebedor
As mesmas regras se aplicam do endpoint de create, porem a requisição deve ser `PATCH` e deve
conter o `ID` do recebedor
```
curl --location --request PATCH 'localhost:8000/api/v1/receiver' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "0f45db07-245f-47e1-8b0e-3b9a7905f082",
    "name": "Lucas Szeremeta",
    "email": "lucaszmt@gmail.com",
    "doc": "000.125.359-52",
    "pix_key_type": "cpf",
    "pix_key": "08412535952",
    "status": "draft"
}'
```

### Recuperar um recebedor por ID
Endpoint responsável por recuperar dados de um recebedor onde deve ser passado o id do mesmo na rota: `/api/v1/receiver/{id}`
```
curl --location --request GET 'localhost:8000/api/v1/receiver/{id}'
```

### Listagem de recebedores
Retorna uma lista de recebedores limitados a 10, onde a mesma recepe um query param `page` que recebe o num da página
```
curl --location --request GET 'localhost:8000/api/v1/receiver?page=1'
```

### Busca de recebedores
Possível realizar a busca de recebedores por seu "Status", "Nome", "Tipo da chave" ou "Valor da chave", via um query param
`query`, e um parametro opcional `limit` pode ser passado também, para que se limite o número de items retornados 
```
curl --location --request GET 'localhost:8000/api/v1/receiver/search?query=lu&limit=20'
```

### Deleção de recebedor(es)
Endpoint para deleção de users, deve ser passado apenas uma lista de ids, que se deseja exluir, no body de
uma requisição `DELETE` como no exemplo abaixo
```
curl --location --request DELETE 'localhost:8000/api/v1/receiver' \
--header 'Content-Type: application/json' \
--data-raw '{
    "ids": [
        "05e12547-9420-4bce-bd88-f40dc5a596a2",
        "40b0b875-8c6e-456b-99f9-4aea2bcea693"
    ]
}'
```