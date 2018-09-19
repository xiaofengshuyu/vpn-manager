package order

// CommitOrderRequest is commit an order request
type CommitOrderRequest struct {
	Email       string `json:"email"`
	OrderNumber string `json:"orderNumber"`
	ProductCode string `json:"productCode"`
}
