package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/luci/go-render/render"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机分区负载均衡
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	producer, err := sarama.NewAsyncProducer([]string{"host4.gcloud.set:9093"}, config)
	if err != nil {
		fmt.Println("producer close err :", err)
		return
	}

	defer producer.Close()

	for i := 0; i < 10; i++ {
		producer.Input() <- &sarama.ProducerMessage{
			Topic: "nginx-log",
			Key:   nil,
			Value: sarama.StringEncoder("this is my " + strconv.Itoa(i) + "kafka message, my message is good"),
		}
	}

	for i := 0; i < 10; i++ {
		select {
		case msg := <-producer.Errors():
			fmt.Println(msg.Err)
		case msg := <-producer.Successes():
			fmt.Println("msg ", render.Render(msg))
		case <-time.After(time.Second):
			fmt.Println("Timeout waiting for msg #", i)
			goto done
		}
	}
done:
	CloseProducer(producer)
}

func CloseProducer(p sarama.AsyncProducer) {
	var wg sync.WaitGroup
	p.AsyncClose()

	wg.Add(2)
	go func() {
		for range p.Successes() {
			fmt.Println("Unexpected message on Successes()")
		}
		wg.Done()
	}()
	go func() {
		for msg := range p.Errors() {
			fmt.Println(msg.Err)
		}
		wg.Done()
	}()
	wg.Wait()
}
