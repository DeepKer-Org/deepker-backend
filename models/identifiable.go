package models

type Identifiable interface {
	GetID() uint
	SetID(uint)
}
