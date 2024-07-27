package main

import "errors"

func AddOne(x int) (int, error) {
	if x == 0 {
		return 0, errors.New("x cannot be 0")
	}

	return x + 1, nil
}

func Double(x int) (int, error) {
	if x == 0 {
		return 0, errors.New("x cannot be 0")
	}

	return x * 2, nil
}

type functionAsType func() (int, error)

func Compose(f functionAsType, g functionAsType) func() int {
	first, err1 := f()
	if err1 != nil {
		return func() int {
			return 0
		}
	}

	second, err2 := g()
	if err2 != nil {
		return func() int {
			return 0
		}
	}

	return func() int {
		return first + second
	}
}
