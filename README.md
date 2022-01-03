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
# MySQL 8.0
CREATE USER 'qnote' IDENTIFIED BY 'pa$$w0rd';
CREATE DATABASE qnote CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
GRANT ALL PRIVILEGES ON qnote.* TO 'qnote';
```

如果您的 MySQL 版本低于 8.0，请参考 [此贴](https://dba.stackexchange.com/questions/76788/create-a-mysql-database-with-charset-utf-8) 来设置数据库的字符集。

## REST Client

`./restClient` 目录下的文件是用于开发时测试的， 可以使用 [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) 插件

## reference

https://github.com/lightyears1998/qnote-backend