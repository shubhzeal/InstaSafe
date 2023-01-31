package Model

type Transaction struct {
	Id        uint    `json:"id" gorm:"primary_key"`
	Amount    float64 `json:"amount"`
	Timestamp string  `json:"timestamp"`
}

type CreateTransactionInput struct {
	Amount    float64 `json:"amount" binding:"required"`
	Timestamp string  `json:"timestamp" binding:"required"`
}
