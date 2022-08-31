<h1 align="center">
    <img alt="Star Wars API" width="600" height="300" title="#Star Wars API" src="https://lumiere-a.akamaihd.net/v1/images/darth-vader-main_4560aff7.jpeg?region=0%2C67%2C1280%2C720&width=960" />
</h1>

<h1 align="center">
   🚀 <a href="#"> STAR WARS API </a> 🚀
</h1>

<p align="center">
 <a href="#-projeto">Projeto</a> |
 <a href="#-tecnologias">Tecnologias</a> |
 <a href="#-prerequisitos">Pré-requisitos</a> |
 <a href="#-features">Features</a> | 
 <a href="#-comoexecutar">Como executar</a> |  
 <a href="#-endpoints">Endpoints</a> | 
 <a href="#-author">Author</a> |
 <a href="#memo-licença">Licença</a> 
</p>

<p align="center">
  <img alt="GitHub language count" src="https://img.shields.io/github/languages/count/flpzow/starwars-api-go?color=red">

  <img alt="Repository size" src="https://img.shields.io/github/repo-size/flpzow/starwars-api-go?color=red">
 
  <img alt="License" src="https://img.shields.io/badge/license-MIT-red">

  <img src="https://img.shields.io/github/stars/flpzow/starwars-api-go?label=stars&message=MIT&color=red" alt="Stars">
</p>

## 💻 Projeto

Este é um projeto para proporcionar dados de Planetas da saga StarWars.

## 🛠️ Tecnologias
- **[Go](https://go.dev/)**
- **[Docker](https://www.docker.com/)**
- **[Gorilla Mux](https://github.com/gorilla/mux)**
- **[MongoDB](https://www.mongodb.com/)**
- **[GoDotEnv](https://github.com/joho/godotenv)**
- **[Testify](https://github.com/stretchr/testify)**


## 🧩 Pré-requisitos

Antes de executá-lo, é necessário ter instalado em sua máquina:

  - Golang (https://go.dev/);
  - Git (https://git-scm.com/);
  - Docker (https://www.docker.com/);
  - Docker-compose (https://docs.docker.com/compose/);
  - Editor de texto (de sua preferência);

## 📝 Features

- [x] Criar um planeta com Nome, Clima e Terreno.
- [x] Buscar em API externa (https://swapi.dev/api/planets) a quantidade de aparições do respectivo planeta.
- [x] Listar todos os planetas.
- [x] Buscar um planeta por ID.
- [x] Buscar um planeta por Nome.
- [x] Remover um planeta.

## ⚙️ Como executar

```bash
$ git clone https://github.com/flpzow/starwars-api-go
$ cd starwars-api-go
```

Para iniciá-lo, siga os passos abaixo:
```bash
# O projeto sobe aplicação e banco de dados em conteiners Docker.
$ docker-compose up --build
```

```bash
# Para verificar os logs, execute:
$ docker logs nome-do-container
```

## 🎯 Endpoints

* CREATE PLANET - /planets (POST) 
* LIST PLANETS - /planets (GET)
* GET PLANET BY ID - /planets/{planetId}" (GET) 
* GET PLANET BY NAME - /planets?search={name} (GET)
* REMOVE PLANET - /planets/{planetId}" (DELETE)

Exemplo de Payload:

```json
 {

    "name":"Tatooine",
    "climate":"arid",
    "terrain":"desert"
  }
```

 Exemplo de Requisição:
 
```sh
curl -X POST \
  http://localhost:8080/planets \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -d '{
	"name":"Tatooine",
	"climate":"arid",
	"terrain":"desert"
}'
```

## ✏️ Author

<img style="border-radius: 50%;" src="https://avatars.githubusercontent.com/flpzow" width="100px;" alt="Vinícius Neto"/> 
 <br />

## :memo: License

Esse projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE.md) para mais detalhes.