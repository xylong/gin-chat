package model

type IModel interface {
	ParseAction(action string) error
}
