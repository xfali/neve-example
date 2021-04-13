# neve-example

该工程为neve的使用示例(一个提供RESTful API的KV存储服务)，集成neve相关的中间件

## 运行
```
go run cmd/main.go
```

## 访问
1. 获得版本号

GET /version
```
curl -X GET "http://localhost:8080/version"
```

2. 查询value

GET /keys/:key
```
curl -X GET "http://localhost:8080/keys/abc"
```

3. 设置key和value

PUT /keys/:key
```
curl -X PUT "http://localhost:8080/keys/abc" -d "123"
```

4. 删除key以及value

DELETE /keys/:key
```
curl -X DELETE "http://localhost:8080/keys/abc"
```

## API Docs (Swagger)
### 1. 访问
```
http://localhost:8080/swagger-ui/index.html
```
### 2. 更新
1. 安装
```
go get -u github.com/swaggo/swag/cmd/swag
```
2. 运行
```
swag init -g cmd/main.go -o docs/swagdoc
```



