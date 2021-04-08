# neve-example

该工程为neve的使用示例，集成neve相关的中间件

## swagger文档生成
安装
```
go get -u github.com/swaggo/swag/cmd/swag
```
运行
```
swag init -g cmd/main.go -o docs/swagdoc
```
访问
```
http://localhost:8080/swagger-ui/index.html
```


