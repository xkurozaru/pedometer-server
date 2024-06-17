package common

import "github.com/aidarkhanov/nanoid"

type NanoID string

func NewNanoID() NanoID {
	return NanoID(nanoid.New())
}

func (id NanoID) String() string {
	return string(id)
}
