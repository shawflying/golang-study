# gin
> golang的web api框架

## golang

```go
go get github.com/gin-gonic/gin
```

- https://studygolang.com/articles/11819

- https://www.cnblogs.com/nickchen121/p/11517442.html

- https://www.jianshu.com/p/98965b3ff638/

- https://www.cnblogs.com/ningskyer/articles/6297356.html
- https://geektutu.com/post/quick-go-gin.html
- https://www.bookstack.cn/books/gin-doc

- [gin框架中文文档] https://learnku.com/docs/gin-gonic/2018/gin-readme/3819
- [Gin Web 框架] https://rsy.me/gin-chn-document/
- [Gin实践 连载十二 优化配置结构及实现图片上传] https://www.bookstack.cn/read/gin-EDDYCJY-blog/golang-gin-2018-05-27-Gin实践-连载十二-优化配置结构及实现图片上传.md
- [golang 文档](https://www.kancloud.cn/uvohp5na133/golang/934170)

## 日志
```go
go get github.com/sirupsen/logrus
```
## 配置
```go
go get gopkg.in/yaml.v2
```

## 工具
```go
# 有demo
go get github.com/satori/go.uuid
go get github.com/google/uuid
go get github.com/robfig/cron
```

## mongodb
```go
go get go.mongodb.org/mongo-driver/mongo
```
## redis
```go
go get github.com/go-redis/redis

go get -u github.com/go-redis/redis
```
## http
```go
go get -u github.com/valyala/fasthttp
```

## 接口
```go
r.StaticFS("/static", http.Dir("html"))
// 在公司的浏览器中可以显示ico
r.StaticFile("/favicon.ico", "./html/favicon.ico")
```
