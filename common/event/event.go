package event

type QueueHandler struct {
	Queue   string
	Handler interface{}
}

var EventSubPayloadMap = map[Subject]any{
	Subject_PaymentProcessed: PayloadUserCharge{},
}
