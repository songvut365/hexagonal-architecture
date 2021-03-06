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

## APIs
- Get All Customer
```
$ curl localhost:8000/customers -i
```

- Get Customer By ID
```
curl localhost:8000/customers/1 -i
```

- Get Account
```
$ curl localhost:8000/accounts/1/account -i
```

- Create New Account
```
$ curl localhost:8000/accounts/1 -i \
    -X POST \
    -H "content-type:application/json" \
    -d '{"account_type":"saving", "amount":500}'
```

## Reference 
- [Go Programming - Hexagonal Architecture](https://www.youtube.com/watch?v=k3JZI-sQs2k)