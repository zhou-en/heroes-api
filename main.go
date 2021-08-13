package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getenv("DB_USERNAME"))
	fmt.Println(os.Getenv("DB_NAME"))
	fmt.Println(os.Getenv("DB_PASSWORD"))
	a := App{}
	a.Initialize(
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	a.Run(":8010")
}
