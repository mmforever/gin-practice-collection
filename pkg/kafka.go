package pkg

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

func ProductMessage(msg map[string]interface{}) bool {

	msg_byte, err := json.Marshal(msg)

	logrus.WithFields(logrus.Fields{
		"err": err,
		"msg": msg,
	}).Info("解析 message 出错")

	config := &kafka.ConfigMap{
		"bootstrap.servers": KafkaConfig.Servers,
		"security.protocol": "SASL_PLAINTEXT",
		"sasl.mechanism":    "PLAIN",
		"sasl.username":     KafkaConfig.User,
		"sasl.password":     KafkaConfig.Password,
	}
	producer, err := kafka.NewProducer(config)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
			"msg": msg,
		}).Error("解析 message 出错")
		return false
	}

	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					logrus.WithFields(logrus.Fields{
						"err": ev.TopicPartition.Error,
					}).Error("Delivery failed:")
				} else {
					logrus.WithFields(logrus.Fields{
						"detail": ev.TopicPartition,
					}).Error("消息投递情况")
				}
			}
		}
	}()

	go func() {
		err := producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &KafkaConfig.Topic, Partition: kafka.PartitionAny},
			Value:          msg_byte,
		}, nil)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Error("消息投递情况")
			return
		}
		logrus.WithFields(logrus.Fields{}).Error("send message success")
		fmt.Println("send message success")
	}()
	producer.Flush(15 * 1000)
	return true
}

func ComsumerMessage() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": KafkaConfig.Servers,
		"security.protocol": "SASL_PLAINTEXT",
		"sasl.mechanism":    "PLAIN",
		"sasl.username":     KafkaConfig.User,
		"sasl.password":     KafkaConfig.Password,
		"group.id":          KafkaConfig.Group,
		"auto.offset.reset": KafkaConfig.Offset,
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{KafkaConfig.Topic}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		logrus.WithFields(logrus.Fields{
			"topic":  msg.TopicPartition,
			"detail": string(msg.Value),
			"err":    err,
		}).Info("消息投递情况")

	}

}
