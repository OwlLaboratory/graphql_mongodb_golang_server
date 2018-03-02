package entity

type entity interface {
	ID() string
	Created() string
	Updated() string
}
