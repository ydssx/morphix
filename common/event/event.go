package event

type QueueHandler struct {
	Queue   string
	Handler interface{}
}
