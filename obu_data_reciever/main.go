package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/perfect047/roadcharge/types"
)

type DataReciever struct {
	conn     *websocket.Conn
	producer DataProducer
}

func NewDataReciever() (*DataReciever, error) {
	kafkaTopic := "obudata"
	p, err := NewKafkaProducer(kafkaTopic)
	if err != nil {
		return nil, err
	}
	p = NewLogMiddleware(p)
	return &DataReciever{
		producer: p,
	}, err
}

func (dr *DataReciever) handleWS(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	dr.conn = conn

	go dr.wsRecieveLoop()
}

func (dr *DataReciever) wsRecieveLoop() {
	fmt.Printf("New OBU client connected!\n")
	var data types.OBUData
	for {
		err := dr.conn.ReadJSON(&data)

		if err != nil {
			log.Printf("websocket data read error: %+v\n", err)
			continue
		}

		fmt.Printf("Recived data from [%d]: %+v \n", data.OBUID, data)
		if err := dr.producer.ProduceData(data); err != nil {
			fmt.Printf("Kafka produced error: %+v\n", err)
		}

	}
}

func main() {
	recv, err := NewDataReciever()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/ws", recv.handleWS)
	http.ListenAndServe(":30000", nil)
}
