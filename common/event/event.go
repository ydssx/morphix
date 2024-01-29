package event

var eventSubjectMap = map[Subject]interface{}{
	Subject_PaymentCompleted: &PayloadPaymentCompleted{},
	Subject_CancelPayment:    &PayloadCancelPayment{},
}

func GetEvent(subject Subject) interface{} {
	return eventSubjectMap[subject]
}