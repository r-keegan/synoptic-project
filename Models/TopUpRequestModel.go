package Models

type TopUpRequest struct {
	CardID string
	Pin    string
	Amount int
}
