package common

import "github.com/aidarkhanov/nanoid"

type NanoID string

func NewNanoID() NanoID {
	return NanoID(nanoid.New())
}
