package model

type ResponseClient[T any] struct {
	Message string `json:"message"`
	Data    *T
}
