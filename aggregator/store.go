package main

import "github.com/Hhanri/toll_calculator/types"

type Storer interface {
	Insert(types.Distance) error
}
