package controller

import (
	"app/models"

	"github.com/bxcodec/faker/v3"
)

var Users []models.User

func CreateUser(data models.User) {
	Users = append(Users, data)
}

func GetListUser(req models.GetListRequest) (resp []models.User, err bool) {

	if req.Limit+req.Offset > len(Users) {
		return []models.User{}, true
	}

	return Users[req.Offset : req.Limit+req.Offset], false
}

func GenerateUser(count int) {
	for i := 0; i < count; i++ {
		Users = append(Users, models.User{
			Id:      i + 1,
			Name:    faker.FirstName(),
			Surname: faker.LastName(),
		})
	}
}

// func GetByIdUser(id int) models.User {
// 	for _, val := range Users {
// 		if val.Id == id {
// 			return val
// 		}
// 	}
// 	return models.User{}
// }

// func UpdateUser(data models.User) {
// 	for in, val := range Users {
// 		if val.Id == data.Id {
// 			Users[in].Name = data.Name
// 			Users[in].Surname = data.Surname
// 			return
// 		}
// 	}
// 	fmt.Println("Bunday User yo'q")
// }

// func DeleteUser(data models.User) {
// 	for in, val := range Users {
// 		if val.Id == data.Id {
// 			Users = append(Users[:in], Users[in+1:]...)
// 			return
// 		}
// 	}
// 	fmt.Println("Bunday User yo'q")
// }

// getbyid
// update
// delete
