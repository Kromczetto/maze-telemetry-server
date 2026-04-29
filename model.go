package main

import "time"

type Run struct {
	Time      int64     `json:"time" bson:"time"`
	Cells     int       `json:"cells" bson:"cells"`
	Turns     int       `json:"turns" bson:"turns"`
	Algorithm string    `json:"algorithm" bson:"algorithm"`

	Width     int       `json:"width" bson:"width"`
	Height    int       `json:"height" bson:"height"`

	Maze      []int     `json:"maze" bson:"maze"` 
	Walls     []int     `json:"walls" bson:"walls"`

	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}