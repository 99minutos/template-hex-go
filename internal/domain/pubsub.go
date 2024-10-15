package domain

import "github.com/bytedance/sonic"

type Message struct {
	Data        []byte                 `json:"data" bson:"data"`
	MessageId   string                 `json:"message_id" bson:"message_id"`
	PublishTime string                 `json:"publish_time" bson:"publish_time"`
	Attributes  map[string]interface{} `json:"attributes,omitempty" bson:"attributes,omitempty"`
}
type RequestPubsub struct {
	Subscription string  `json:"subscription" bson:"subscription"`
	Message      Message `json:"message" bson:"message"`
}

func NewPubsubFromBytes(data []byte) (*RequestPubsub, error) {
	var req RequestPubsub
	err := sonic.Unmarshal(data, &req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}

func (m *Message) DecodeData(someType interface{}) error {
	return sonic.Unmarshal(m.Data, someType)
}
