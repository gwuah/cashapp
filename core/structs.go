package core

type Pagination struct {
	CurrentPage  int   `json:"current_page,omitempty"`
	NextPage     int   `json:"next_page,omitempty"`
	PreviousPage int   `json:"previous_page,omitempty"`
	Count        int64 `json:"count"`
}

type Meta struct {
	Data       interface{} `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Message    string      `json:"message"`
}

type Response struct {
	Error bool `json:"error"`
	Code  int  `json:"code"`
	Meta  Meta `json:"meta"`
}

type CreateUserRequest struct {
	Tag string `json:"tag"`
}

type CreatePaymentRequest struct {
	From        int    `json:"from"`
	To          int    `json:"to"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
}
