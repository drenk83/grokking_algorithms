package main

import "errors"

var ErrNotFound = errors.New("not found")

func LinearSearch(arr []int, target int) (int, error) {
	for idx, val := range arr {
		if val == target {
			return idx, nil
		}
	}
	return -1, ErrNotFound
}
