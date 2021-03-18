package kafka

import (
	"fmt"
	_ "os"
	_ "os/signal"
	_ "syscall"
	"testing"
)

func TestProducerMsg(t *testing.T) {

	// producerApi := NewProducerApi([]string{"172.29.203.62:9092"})
	producerApi := NewProducerApiWithSASLPlainText([]string{"172.29.202.56:9092"}, "username", "password")

	defer producerApi.Close()

	c := producerApi.PutFromString("test-sasl", "test a message")

	fmt.Println(c)
}
