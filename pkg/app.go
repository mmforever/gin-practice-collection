package pkg

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type App struct {
	Port     string
	Merchant string
	Rider    string
}

var AppConfig *App

type Mysql struct {
	Host     string
	User     string
	Password string
	Dbname   string
}

var MysqlConfig *Mysql

type RedisCon struct {
	Addr     string
	Password string
	Db       int
}

var RedisConfig *RedisCon

type Kafka struct {
	Servers  string
	User     string
	Password string
	Topic    string
	Group    string
	Offset   string
}

var KafkaConfig = &Kafka{}

//读取配置文件
func AppSetup() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")

	if err := viper.ReadInConfig(); err != nil {
		logrus.WithFields(logrus.Fields{
			"config err": err,
		}).Info("config 配置有误")
		return
	}

	err := Sub("app", &AppConfig)
	logrus.WithFields(logrus.Fields{
		"err":        err,
		"app config": AppConfig,
	}).Info("获取 app 配置")

	err = Sub("mysql", &MysqlConfig)
	logrus.WithFields(logrus.Fields{
		"err":          err,
		"mysql config": MysqlConfig,
	}).Info("获取 mysql 配置")

	err = Sub("redis", &RedisConfig)
	logrus.WithFields(logrus.Fields{
		"err":          err,
		"redis config": RedisConfig,
	}).Info("获取 redis 配置")

	err = Sub("kafka", &KafkaConfig)
	logrus.WithFields(logrus.Fields{
		"err":          err,
		"kafka config": KafkaConfig,
	}).Info("获取 kafka 配置")

}

func Sub(key string, value interface{}) error {
	sub := viper.Sub(key)
	return sub.Unmarshal(value)
}
