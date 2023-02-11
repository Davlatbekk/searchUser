package main

import (
	"fmt"

	"app/controller"
	"app/models"
)

func main() {

	page := 1000
	controller.GenerateUser(page)
	fmt.Println((page / 10))

	var sum int

	fmt.Scan(&sum)

	users, err := controller.GetListUser(models.GetListRequest{
		Offset: (sum - 1) * 10,
		Limit:  10,
	})

	if err {
		fmt.Println("users out of range")
		return
	}

	for _, user := range users {
		fmt.Println(user)
	}
}
