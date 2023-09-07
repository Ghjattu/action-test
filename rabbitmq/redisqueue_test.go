package rabbitmq

import "testing"

func TestProduce(t *testing.T) {
	RedisMQ.Producer(&Message{
		Type: "test",
	})
}
