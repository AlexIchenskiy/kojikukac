package consumer

//Copyright (c) Microsoft Corporation. All rights reserved.
//Copyright 2016 Confluent Inc.
//Licensed under the MIT License.
//Licensed under the Apache License, Version 2.0
//
//Original Confluent sample modified for use with Azure Event Hubs for Apache Kafka Ecosystems
/*Paste into terminal:
export KAFKA_EVENTHUB_ENDPOINT="cbq-hackathon.servicebus.windows.net:9093"
export KAFKA_EVENTHUB_CONNECTION_STRING="Endpoint=sb://cbq-hackathon.servicebus.windows.net/;SharedAccessKeyName=n;SharedAccessKey=p3fH0pzw46YajywaIyAaWRK+HGqMBLgBV+AEhNWlq+4=;EntityPath=team8"
*/
import (
	//"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:6969", "http service address")

var upgrader = websocket.Upgrader{} // use default options
var message string

/*
	type ParkingSpotUpdate struct {
		Id         string
		IsOccupied bool
		Time       string
	}

	type ParkingSpot struct {
		id                string
		latitude          int
		longitude         int
		parkingSpotZone   string
		occupied          bool
		occupiedTimeStamp string
	}
*/
func Consume(from string) {
	consumerGroup := "consumergroup"

	// https://github.com/edenhill/librdkafka/blob/master/CONFIGURATION.md
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "cbq-hackathon.servicebus.windows.net:9093",
		"sasl.mechanisms":   "PLAIN",
		"security.protocol": "SASL_SSL",
		"sasl.username":     "$ConnectionString",
		"sasl.password":     "Endpoint=sb://cbq-hackathon.servicebus.windows.net/;SharedAccessKeyName=n;SharedAccessKey=p3fH0pzw46YajywaIyAaWRK+HGqMBLgBV+AEhNWlq+4=;EntityPath=team8",
		"group.id":          consumerGroup,
		"auto.offset.reset": "earliest",
		"debug":             "consumer",
	})

	if err != nil {
		panic(err)
	}

	topics := []string{"team8"}
	c.SubscribeTopics(topics, nil)
	//var newspot ParkingSpotUpdate

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("%s\n", string(msg.Value))

			message = string(msg.Value)
			var writer http.ResponseWriter

			flag.Parse()
			log.SetFlags(0)
			http.HandleFunc("/echo", Echo)
			log.Fatal(http.ListenAndServe(*addr, nil))

			Echo(writer, &http.Request{})

			/*
				err1 := json.Unmarshal([]byte(string(msg.Value)), &newspot)

				if err1 != nil {
					fmt.Printf("%s\n", err1)
				} else {

					fmt.Printf("Id is:%s\n", newspot.Id)
					fmt.Printf("Occupied is:%t\n", newspot.IsOccupied)
					fmt.Printf("Time is:%s\n", newspot.Time)
				}
			*/

		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

	c.Close()
}

func Echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	//_, message, err := c.ReadMessage()
	if err != nil {
		log.Println("read:", err)
	}

	err = c.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println("write:", err)
	}

}
