package main

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/Hhanri/toll_calculator/types"
	"github.com/gorilla/websocket"
)

const sendInterval = time.Second
const wsEndpoint = "ws://127.0.0.1:3000/ws"

func genCoord() float64 {
	n := float64(rand.Intn(90) + 1)
	f := rand.Float64()
	return n + f
}

func genGeo() types.Geo {
	return types.Geo{
		Lat: genCoord(),
		Lng: genCoord(),
	}
}

func generateObuIds(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}
	return ids
}

func main() {

	obuIds := generateObuIds(10)

	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		for i := 0; i < len(obuIds); i++ {
			data := types.ObuData{
				ObuId: obuIds[i],
				Geo:   genGeo(),
			}
			if err := conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}
		}
		time.Sleep(sendInterval)
	}
}
