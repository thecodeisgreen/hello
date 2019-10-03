package main

import (
	"fmt"

	"hello/models/users"

	"go.mongodb.org/mongo-driver/bson"

	"hello/tcig.io/count"
)

func main() {
	count.Inc()
	fmt.Println(count.Get())
	fmt.Println("Hello")
	fmt.Println("Hello you")
	users.Create("bettan", "hubert")
	fmt.Println(users.GetOneByID("5d95d27f411f0b2a0283d243"))
	fmt.Println(users.Get(bson.M{"lastname": "bettan"}))
}
