
# Go TodoApp API 

## About the project

- Based on REST API 
- Powered up with gin framework
- Сlean Architecture with dependency injection
- Running via docker
- Work with db done with slqx
- JWT middleware for registration and authentication

## Technologies
* web framework: [gin](https://github.com/gin-gonic/gin)
* containerize: [docker](https://www.docker.com/)
* swagger: [swaggo](https://github.com/swaggo/swag)
* database:
    * [sqlx](https://github.com/jmoiron/sqlx) (db wrapper)
    * [postgres](https://github.com/lib/pq)
* config: [viper](https://github.com/spf13/viper)
* logger: [logrus](https://github.com/sirupsen/logrus)

## Usage
- Clone repository
```bash
git clone github.com/SimilarEgs/go-todo-app
```

- Run via docker
```bash
make compose-up
```

## Swagger 
```bash
make swag
```
```bash
http://localhost:8080/swagger/index.html
```
![SharedScreenshot](https://user-images.githubusercontent.com/90198202/180997065-50c58625-627e-4067-9791-166d1470d41d.jpg)
