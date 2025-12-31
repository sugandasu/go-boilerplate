package dto

type Response struct {
	Data    any    `json:"data,omitempty"`
	Message string `json:"message"`
	Errors  any    `json:"errors,omitempty"`
}

type Pagination struct {
	Total int `json:"total"`
}

type PaginationResponse struct {
	Items      any        `json:"items"`
	Pagination Pagination `json:"pagination"`
}
