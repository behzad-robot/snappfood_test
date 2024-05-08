package rabbit

import (
	"context"
	"math/rand"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitHelperService interface {
	Close()
	PublishTo(queueName string, body []byte) error
	SimpleConsumeOne(queueName string) ([]byte, bool, error)
}
type rabbitHelperService struct {
	rabbitConnection *amqp.Connection
	publisherChannel *amqp.Channel
}

func NewRabbitHelperService(rabbitConnection *amqp.Connection) RabbitHelperService {
	publisherChannel, err := rabbitConnection.Channel()
	if err != nil {
		panic(err)
	}
	return &rabbitHelperService{
		rabbitConnection: rabbitConnection,
		publisherChannel: publisherChannel,
	}
}
func (helper *rabbitHelperService) Close() {
	helper.publisherChannel.Close()
}
func (helper *rabbitHelperService) PublishTo(queueName string, body []byte) error {
	return helper.publisherChannel.PublishWithContext(context.Background(),
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: randomString(6),
			Body:          body,
		})
}
func (helper *rabbitHelperService) SimpleConsumeOne(queueName string) ([]byte, bool, error) {
	ch, err := helper.rabbitConnection.Channel()
	if err != nil {
		return nil, false, err
	}
	defer ch.Close()
	_, err = ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, false, err
	}
	// Set prefetch count to 1
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		return nil, false, err
	}
	// Consume a single message from the queue
	msg, ok, err := ch.Get(queueName, true)
	if err != nil || !ok {
		return nil, ok, err
	}
	return msg.Body, ok, err
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
