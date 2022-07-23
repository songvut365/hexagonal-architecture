# Hexagonal Architecture

## How to start

1. Start MySQL
```
$ docker run --name my-mysql -p 3306:3306 \
      -e MYSQL_ROOT_PASSWORD=1234 \
      -e MYSQL_DATABASE=banking \
      -e MYSQL_USER=songvut \
      -e MYSQL_PASSWORD=1234 \
      -d mysql
```

2. Run application
```
$ go run main.go
```

## Directory structure

```
hexagonal-architecture
├── README.md
├── config
│   ├── database.go
│   ├── timezone.go
│   └── viper.go
├── config.yaml
├── go.mod
├── go.sum
├── handler
│   └── customer.go
├── logs
│   └── logs.go
├── main.go
├── repository
│   ├── customer.go
│   ├── customer_db.go
│   └── customer_mock.go
└── service
    ├── customer.go
    └── customer_service.go
```

## How to run project
```
$ go run .
```

## Reference 
- [Go Programming - Hexagonal Architecture](https://www.youtube.com/watch?v=k3JZI-sQs2k)