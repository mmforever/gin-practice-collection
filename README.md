# gin-practice-collection
基于 gin 的常用功能最佳实践的集合


## log 基于 sirupsen/logrus 和 lestrrat-go/file-rotatelogs 实现日志按日期切割
### 关键文件 pkg/log.go

使用方法：
logrus.WithFields(logrus.Fields{
			"claims": claims,
			"err":    err,
		}).Info("decode token")
    
    
   
   
   
   
## 读取配置文件 基于 spf13/viper 将配置文件读取到相应的结构体中，供别的包使用
### 关键文件 pkg/app.go
使用方法：增加其他的配置修改该文件，增加相应结构体和读取配置





## 基于 golang-jwt/jwt 实现 jwt 的加解密，并运用于 gin 的中间件中，这里实现了同一个项目有两个密钥的情况，name 表示密钥名称，
### 关键文件 pkg/jwt.go，middleware/TokenAuth.go

使用方法：
加密：pkg.Encode(name, data)
解密：pkg.Decode(name, token)




## 基于 confluent-kafka-go/kafka 实现投递和消费操作
### 关键文件 pkg/kafka.go
使用方法：
投递：pkg.ProductMessage(message)
消费：pkg.ComsumerMessage()





## 基于 go-redis/redis/v9 操作redis
### 关键文件 pkg/redis.go
使用方法：
err := pkg.Redis.Set(pkg.RedisCtx, key, 1, d).Err()







