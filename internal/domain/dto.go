package domain

type ListView[T any] struct {
	List []T `json:"list"`
	Page int `json:"page"`
}
