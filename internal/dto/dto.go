package dto

type List[T any] struct {
	TotalCount int64 `json:"total_count"`
	Items      []T   `json:"items"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

type CreatedResponse struct {
	ID string `json:"id"`
}
