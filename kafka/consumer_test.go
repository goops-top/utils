package kafka

import (
	_ "fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestConsumerMsg(t *testing.T) {
	// no auth
	// consumerApi := NewConsumerApi([]string{"172.29.203.62:9092"}, "BGBiao", "earliest")
	// sasl/plaintext auth
	consumerApi := NewConsumerApiWithSASLPlainText([]string{"172.29.202.78:9092"}, "BGBiao", "earliest", "username", "password")

	defer consumerApi.Close()

	c := consumerApi.ConsumerMsgFromTopics([]string{"test-sasl"})

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigterm:
		log.Warnln("terminating: via signal")
	}
	c()

}
