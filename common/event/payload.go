package event

type PayloadUserCharge struct {
	UserId  int64   `json:"user_id"`
	OrderId string  `json:"order_id"`
	Amount  float64 `json:"amount"`
}
