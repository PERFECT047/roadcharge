package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
	"github.com/perfect047/roadcharge/types"
)

const sendInterval = 1
const vehicleCount = 20
const wsEndpoint = "ws://127.0.0.1:30000/ws"

func genLocation() (float64, float64) {
	return genCoord(), genCoord()
}

func genCoord() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()

	return n + f
}

func genObuIds(n int) []int {
	OBUIds := make([]int, n)

	for i := 0; i < n; i++ {
		OBUIds[i] = rand.Intn(math.MaxInt)
	}

	return OBUIds
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	OBUIds := genObuIds(vehicleCount)
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)

	if err != nil {
		log.Fatal(err)
	}
	for {
		for i := range vehicleCount {
			lat, long := genLocation()
			data := types.OBUData{
				OBUID:     OBUIds[i],
				Latitude:  lat,
				Longitude: long,
			}

			fmt.Printf("Vehicle details: %+v\n", data)

			if err := conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}
		}
		time.Sleep(sendInterval * time.Second)
	}
}
