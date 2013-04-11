package main

import (
	"errors"
)

func InvalidSliceError(srcFunc, sliceName string) error {
	return errors.New("Slice : " + sliceName + " is invalid in " + srcFunc)
}

func OutOfBoundsError(srcFunc string) error {
	return errors.New("Index out of bounds in " + srcFunc)
}
