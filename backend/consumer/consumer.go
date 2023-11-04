package main

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ParkingSpotUpdate struct {
	id           string
	occupied     bool
	occupiedTime string
}

type ParkingSpot struct {
	id                string
	latitude          int
	longitude         int
	parkingSpotZone   string
	occupied          bool
	occupiedTimeStamp string
}

func main() {
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
		panic(err)
	}

	topics := []string{"team8"}
	c.SubscribeTopics(topics, nil)
	var newspot ParkingSpotUpdate
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))

			err1 := json.Unmarshal([]byte(msg.Value), &newspot)

			if err1 != nil {
				fmt.Printf("%s\n", err1)
			}

			fmt.Printf("Struct is:", newspot)
			fmt.Printf("Id is:%s\n", newspot.id)
			fmt.Printf("Occupied is:%t\n", newspot.occupied)
			fmt.Printf("Time is:%s\n", newspot.occupiedTime)
			break
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

	c.Close()
}
