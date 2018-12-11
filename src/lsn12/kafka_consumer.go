package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/luci/go-render/render"
)

var (
	wg sync.WaitGroup
)

func main() {
	consumer, err := sarama.NewConsumer(strings.Split("host4.gcloud.set:9093", ","), nil)
	if err != nil {
		fmt.Println("Failed to start consumer : ", err)
		return
	}

	partitionList, err := consumer.Partitions("nginx-log")
	partitionListError, err := consumer.Partitions("nginx-error")
	if err != nil {
		fmt.Println("Failed to start consumer for partition : ", err)
		return
	}

	fmt.Println("partitionList ", render.Render(partitionList))

	consueMsg := func(pc sarama.PartitionConsumer) {
		for msg := range pc.Messages() {
			fmt.Println("key ", string(msg.Key), " value ", string(msg.Value), " offset ", msg.Offset, " par ", msg.Partition)
		}
		time.Sleep(time.Second)
	}

	for {
		for partion := range partitionList {
			pc, err := consumer.ConsumePartition("nginx-log", int32(partion), sarama.OffsetNewest)

			if err != nil {
				fmt.Println("Failed to start consumer for partition ", partion, " err ", err)
				return
			}

			defer pc.AsyncClose()
			consueMsg(pc)
		}

		for partion := range partitionListError {
			pc, err := consumer.ConsumePartition("nginx-error", int32(partion), sarama.OffsetNewest)

			if err != nil {
				fmt.Println("Failed to start consumer for partition ", partion, " err ", err)
				return
			}

			defer pc.AsyncClose()
			consueMsg(pc)
		}

		time.Sleep(time.Second)
	}

	wg.Wait()
	consumer.Close()
	select {}
}
