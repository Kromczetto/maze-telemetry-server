package main

import "time"

type Run struct {
	Time int64 `json:"time" bson:"time"`
	Cells int `json:"cells" bson:"cells"`
	Turns int `json:"turns" bson:"turns"`
	Algorithm string `json:"algorithm" bson:"algorithm"`
	Maze []int `json:"maze" bson:"maze"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}