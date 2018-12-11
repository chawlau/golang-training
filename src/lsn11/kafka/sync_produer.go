package main

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机分区负载均衡
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err := sarama.NewAsyncProducer([]string{"host4.gcloud.set:9093"}, config)
	if err != nil {
		fmt.Println("producer close err :", err)
		return
	}

	defer client.Close()

	for {
		msg := &sarama.ProducerMessage{}
		msg.Topic = "nginx_log"
		msg.Value = sarama.StringEncoder("this is my first kafka message, my message is good")

		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send message failed ", err)
			return
		}

		fmt.Printf("pid %v offset %v \n", pid, offset)

		time.Sleep(1000 * time.Millisecond)
	}
}
