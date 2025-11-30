package consumer

import (
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
	"queue-worker/config"
	"queue-worker/httpclient"
)

func StartConsumer(cfg config.Config) error {
	conn, err := amqp.Dial(cfg.RabbitURL)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		cfg.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		cfg.QueueName,
		"",
		false, 
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	apiClient := httpclient.NewAPIClient(cfg.NestAPIURL, cfg.HTTPTimeout)

	for msg := range msgs {
		log.Println("Received message")

		retry := 0
		processErr := ProcessMessage(msg.Body, apiClient)

		for processErr != nil && retry < cfg.MaxRetries {
			log.Println("Retrying due to error:", processErr)
			retry++
			processErr = ProcessMessage(msg.Body, apiClient)
		}

		if processErr != nil {
			log.Println("Discarding message after max retries:", processErr)
			msg.Nack(false, false)
			continue
		}

		msg.Ack(false)
		log.Println("Message processed successfully")
	}

	return nil
}
