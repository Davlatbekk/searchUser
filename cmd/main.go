package main

import (
	"fmt"

	"app/controller"
	"app/models"
)

func main() {

	page := 1000
	controller.GenerateUser(page)

	fmt.Println((page / 10), "ta sahifa bor")
	fmt.Println("SAHOFANI TANLANG")

	var sum int

	fmt.Scan(&sum)

	fmt.Println("ISM BUYICHA QIDIRING")

	var search string

	fmt.Scan(&search)

	users, err := controller.GetListUser(models.GetListRequest{
		Offset:   (sum - 1) * 10,
		Limit:    20,
		Search:   search,
		FromDate: "2003-01-01",
		ToDate:   "2020-10-01",
	})

	if err {
		fmt.Println("users out of range")
		return
	}

	for _, user := range users {
		fmt.Println(user)
	}
}
