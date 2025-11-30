package main

import (
	"log"
	"queue-worker/config"
	"queue-worker/consumer"
)

func main() {
	cfg := config.Load()

	log.Println("Starting Go worker...")

	if err := consumer.StartConsumer(cfg); err != nil {
		log.Fatal("Worker failed:", err)
	}
}
