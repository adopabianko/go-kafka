package producer_test

import (
	"fmt"
	"testing"

	"github.com/Shopify/sarama/mocks"
	"github.com/adopabianko/go-kafka/producer"
)

func TestSendMessage(t *testing.T) {
	t.Run("Send message ok", func(t *testing.T) {
		// membuat mocking untuk producer
		mockedProducer := mocks.NewSyncProducer(t, nil)

		// membuat expect producer success atau mengirim pesan
		mockedProducer.ExpectSendMessageAndSucceed()
		kafka := &producer.KafkaProducer{
			Producer: mockedProducer,
		}

		msg := "Message 1"

		err := kafka.SendMessage("test_topic", msg)
		if err != nil {
			t.Errorf("Send message should not be error but have: %v", err)
		}
	})

	t.Run("Send message not ok", func(t *testing.T) {
		mockedProducer := mocks.NewSyncProducer(t, nil)
		// membuat proses kirim pesan gagal
		mockedProducer.ExpectSendMessageAndFail(fmt.Errorf("Error"))
		kafka := &producer.KafkaProducer{
			Producer: mockedProducer,
		}

		msg := "Message 1"

		err := kafka.SendMessage("test_topic", msg)
		if err == nil {
			t.Error("This should be error")
		}
	})
}
