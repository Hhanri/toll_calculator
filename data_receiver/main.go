package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Hhanri/toll_calculator/types"
	"github.com/gorilla/websocket"
)

func main() {
	var kafkaTopic string = "obu"

	logger := NewLogrusLogger()
	producer, err := NewKafkaProducer(kafkaTopic)
	if err != nil {
		log.Fatal(err)
	}
	producer = NewLogMiddleware(logger, producer)
	receiver, err := NewDataReceiver(producer)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/ws", receiver.handleWs)

	http.ListenAndServe(":3000", nil)
}

type DataReceiver struct {
	msgCh chan types.ObuData
	conn  *websocket.Conn
	prod  DataProducer
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

		// fmt.Printf("received OBU data from [%d] :: <lat %.3f | lng %.3f>\n", data.ObuId, data.Geo.Lat, data.Geo.Lng)

		if err := dr.produceData(data); err != nil {
			log.Println("kafka produce error:", err)
			continue
		}
		//dr.msgCh <- data
	}
}

func (dr *DataReceiver) produceData(data types.ObuData) error {
	return dr.prod.ProduceData(data)
}

func NewDataReceiver(producer DataProducer) (*DataReceiver, error) {
	return &DataReceiver{
		msgCh: make(chan types.ObuData, 128),
		prod:  producer,
	}, nil
}
