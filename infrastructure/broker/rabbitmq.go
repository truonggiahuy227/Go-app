package broker

import (
	"crypto/tls"
	"encoding/json"
	"time"

	"akawork.io/dto"
	"akawork.io/infrastructure"
	"akawork.io/infrastructure/logger"
	"github.com/streadway/amqp"
)

/**
 * Connects to RabbitMQ server
 */
func Connect(host string, port int, userName string, password string) (conn *amqp.Connection) {
	for {
		// Make a infrastructure
		uri := amqp.URI{
			Scheme:   "amqps",
			Host:     host,
			Port:     port,
			Username: userName,
			Password: password,
		}
		conn, err := amqp.DialTLS(uri.String(), &tls.Config{MinVersion: tls.VersionTLS12})

		if err != nil {
			logger.Error("Failed to connect to RabbitMQ %s", err.Error())
			return nil
		}
		// Reconnect to RabbitMQ in case connection died
		//log.Printf("Trying to reconnect to RabbitMQ at %s\n", conn)
		time.Sleep(500 * time.Millisecond)
		return conn
	}
	//return conn
}

/**
 * Establish connection when connection died
 */
func ReConnection(host string, port int, userName string, password string) (conn *amqp.Connection) {
	var rabbitErr *amqp.Error
	// create the rabbit mq error channel
	rabbitCloseError := make(chan *amqp.Error)
	for {
		rabbitErr = <-rabbitCloseError
		if rabbitErr != nil {
			rabbitConn := Connect(host, port, userName, password)
			rabbitCloseError = make(chan *amqp.Error)
			rabbitConn.NotifyClose(rabbitCloseError)
			// run your setup process here
		}
	}
}

/**
 * Creates a RabbitMQ queue
 */
func CreateQueue(ch *amqp.Channel, name string) amqp.Queue {
	queue, err := ch.QueueDeclare(
		name,  // name of the queue
		true,  // should the message be persistent? also queue will survive if the cluster gets reset
		false, // auto delete if there's no consumers (like queues that have anonymous names, often used with fanout exchange)
		false, // exclusive means I should get an error if any other consumer subscribes to this queue
		false, // no-wait means I don't want RabbitMQ to wait if there's a queue successfully setup
		nil,   // arguments for more advanced configuration
	)

	if err != nil {
		logger.Error("Failed to declare a queue %s: %s", name, err.Error())
	}

	return queue
}

/**
 * Pushes a message to RabbitMQ queue
 * @Return {bool}
 */
func PushMessage(rbChannel *amqp.Channel, rbExchange string, rbRouteKey string, action string, obj interface{}) error {
	objTask := dto.Task{
		MessageType: action,
		Data:        obj,
	}
	//m := map[string]string{
	//	"JobId": jobId,
	//}
	task, err := json.Marshal(objTask)
	if err != nil {
		defer rbChannel.Close()
		logger.Error("Rabbitmq - Failed to pair message %s", err.Error())
		return err
	}

	err = rbChannel.Publish(
		rbExchange, // exchange
		rbRouteKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  infrastructure.DefaultContentType,
			Body:         task,
		})

	if err != nil {
		defer rbChannel.Close()
		logger.Error("Rabbitmq - Failed to publish a message %s", err.Error())
		return err
	}

	return nil
}

/**
 * Connects to RabbitMQ server
 */
func ConnectChannel(conn *amqp.Connection) *amqp.Channel {
	// Creates a RabbitMQ channel
	rbChannel, err := conn.Channel()
	if err != nil {
		logger.Error("Error Create Channel %s", err.Error())
		return nil
	}

	// Set RabbitMQ QoS
	err = rbChannel.Qos(1, 0, false)
	if err != nil {
		logger.Error("Error Set Qos Channel %s", err.Error())
	}

	return rbChannel
}

/**
 * Handles error
 */
func HandleError(err error) {
	logger.Error("[RabbitMQ] %s", err.Error())
}
