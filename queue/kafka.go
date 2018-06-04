package queue

import (
	"github.com/Shopify/sarama"
	"fmt"
	"strings"
	"sync"
	"go_partice/tail_kfka/settings"
	"github.com/qq976739120/zhihu-golang-web/elastic"
	"github.com/qq976739120/zhihu-golang-web/pkg/util"
)

type KafkaClient struct {
	Produet_client  sarama.SyncProducer
	Consumer_client sarama.Consumer
	Wg     sync.WaitGroup
}

var KafkaCilent_ *KafkaClient

func InitKafka(addr string) (err error) {
	KafkaCilent_ = new(KafkaClient)
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	produet_client, err := sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		return
	}
	consumer_client, err := sarama.NewConsumer(strings.Split(addr, ","), nil)
	if err != nil {
		fmt.Println("init kafka failed, err:%v", err)
		return
	}
	KafkaCilent_.Produet_client = produet_client
	KafkaCilent_.Consumer_client = consumer_client
	return

}

func SendToKafka(data, topic string) (err error) {

	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	_, _, err = KafkaCilent_.Produet_client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed, err:%v data:%v topic:%v", err, data, topic)
		return
	}
	return
}

func TailToKafka()(err error){
	for {
		msg := util.GetOneLine()
		err = SendToKafka(msg.Msg,msg.Topic)
		if err != nil {
			fmt.Println("send to kafka failed, err:%v", err)
			continue
		}
	}
	return
}

func KafkaToES() (err error) {
	partitionList, err := KafkaCilent_.Consumer_client.Partitions(settings.Kafka_.Topic)
	fmt.Println("list",partitionList)
	if err != nil {
		fmt.Println("Failed to get the list of partitions: ", err)
		return
	}

	for partition := range partitionList {
		pc, errRet := KafkaCilent_.Consumer_client.ConsumePartition(settings.Kafka_.Topic, int32(partition), sarama.OffsetNewest)
		if errRet != nil {
			err = errRet
			fmt.Println("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		go func(pc sarama.PartitionConsumer) {
			KafkaCilent_.Wg.Add(1)
			for msg := range pc.Messages() {
				err =  elastic.SendToES(settings.Kafka_.Topic, msg.Value)
				fmt.Println("es",string(msg.Value))
				if err != nil {
					fmt.Println("send to es failed, err:%v", err)
				}
			}
			KafkaCilent_.Wg.Done()
		}(pc)
	}

	KafkaCilent_.Wg.Wait()
	return
}
