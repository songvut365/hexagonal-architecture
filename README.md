# Hexagonal Architecture

## How to start

1. Start MySQL
```
$ docker run --name my-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=1234 -e MYSQL_DATABASE=banking -d mysql
```

2. Run appliation
```
$ go run main.go
```

## Directory structure

```
hexagonal-architecture
|-- go.mod
|-- go.sum
|-- main.go
|-- README.md
|
|---repository
|     |--customer.go
|     |--customer_db.go
|
|---service
      |--customer.go
      |--customer_service.go
```

## How to run project
```
$ go run .
```

## Reference 
- [Go Programming - Hexagonal Architecture](https://www.youtube.com/watch?v=k3JZI-sQs2k)