package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Hhanri/toll_calculator/types"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gorilla/websocket"
)

var kafkaTopic string = "obu"

func main() {
	receiver, err := NewDataReceiver()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/ws", receiver.handleWs)

	http.ListenAndServe(":3000", nil)
}

type DataReceiver struct {
	msgCh chan types.ObuData
	conn  *websocket.Conn
	prod  *kafka.Producer
}

func (dr *DataReceiver) handleWs(w http.ResponseWriter, r *http.Request) {
	upg := websocket.Upgrader{
		WriteBufferSize: 1028,
		ReadBufferSize:  1028,
	}

	conn, err := upg.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	dr.conn = conn

	go dr.wsReceiveLoop()
}

func (dr *DataReceiver) wsReceiveLoop() {
	fmt.Println("New OBU connected !")
	for {
		var data types.ObuData
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println("read error:", err)
			continue
		}

		fmt.Printf("received OBU data from [%d] :: <lat %.3f | lng %.3f>\n", data.ObuId, data.Geo.Lat, data.Geo.Lng)

		if err := dr.produceData(data); err != nil {
			log.Println("kafka produce error:", err)
			continue
		}
		//dr.msgCh <- data
	}
}

func (dr *DataReceiver) produceData(data types.ObuData) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return dr.prod.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &kafkaTopic,
				Partition: kafka.PartitionAny,
			},
			Value: b,
		},
		nil,
	)

}

func NewDataReceiver() (*DataReceiver, error) {
	p, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers": "localhost",
		},
	)
	if err != nil {
		return nil, err
	}

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	return &DataReceiver{
		msgCh: make(chan types.ObuData, 128),
		prod:  p,
	}, nil
}
