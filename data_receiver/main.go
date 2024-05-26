package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Hhanri/toll_calculator/types"
	"github.com/gorilla/websocket"
)

func main() {
	receiver := NewDataReceiver()

	http.HandleFunc("/ws", receiver.handleWs)

	http.ListenAndServe(":3000", nil)

}

type DataReceiver struct {
	msgCh chan types.ObuData
	conn  *websocket.Conn
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
		//dr.msgCh <- data
	}
}

func NewDataReceiver() *DataReceiver {
	return &DataReceiver{
		msgCh: make(chan types.ObuData, 128),
	}
}
