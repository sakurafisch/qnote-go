# qnote-go

本项目是对 qnote-backend 的重写。本项目正处于并将长期处于社会主义初级阶段。

## 技术栈

- Golang
- Gin
- GORM
- MYSQL

# Run

```
go mod tidy -v
go run main.go
```

or
```
go mod tidy -v
fresh
```

## API

Plz refer to `api.md`

## SQL

```mysql
GRANT ALL PRIVILEGES ON qnote.* TO 'qnote'@'localhost';
```

## reference

https://github.com/lightyears1998/qnote-backend