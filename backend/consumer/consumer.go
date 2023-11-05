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
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var addr = flag.String("addr", "localhost:6969", "http service address")

var message string

func Consume() {
	consumerGroup := "consumergroup"

	// https://github.com/edenhill/librdkafka/blob/master/CONFIGURATION.md
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_EVENTHUB_ENDPOINT"),
		"sasl.mechanisms":   "PLAIN",
		"security.protocol": "SASL_SSL",
		"sasl.username":     "$ConnectionString",
		"sasl.password":     os.Getenv("KAFKA_EVENTHUB_CONNECTION_STRING"),
		"group.id":          consumerGroup,
		"auto.offset.reset": "earliest",
		"debug":             "consumer",
	})

	if err != nil {
		log.Fatal(err)
	}

	topics := []string{"team8"}
	c.SubscribeTopics(topics, nil)

	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := websocket.Accept(w, r, nil)
		if err != nil {
			// ...
		}
		defer c.CloseNow()

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
		defer cancel()

		var v interface{}
		err = wsjson.Write(ctx, c, &message)
		if err != nil {
			// ...
		}

		log.Printf("received: %v", v)

		c.Close(websocket.StatusNormalClosure, "")
	})
	go serve(fn)

	for true {
		msg, err := c.ReadMessage(-1)

		if err != nil {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
		log.Printf("%s\n", string(msg.Value))

		message = string(msg.Value)

		//var writer http.ResponseWriter

		/*flag.Parse()
		  log.SetFlags(0)*/

		//Echo(writer, &http.Request{})

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
	}

	c.Close()
}

func serve(fn http.HandlerFunc) {
	err := http.ListenAndServe("localhost:6969", fn)
	log.Fatal(err)
}
