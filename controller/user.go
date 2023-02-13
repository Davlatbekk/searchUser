package controller

import (
	"app/models"
	"strconv"
	"strings"

	"github.com/bxcodec/faker/v3"
)

var Users []models.User

func CreateUser(data models.User) {
	Users = append(Users, data)
}

func GetListUser(req models.GetListRequest) (resp []models.User, err bool) {
	users := []models.User{}
	if req.Limit+req.Offset > len(Users) {
		return []models.User{}, true
	}
	for _, val := range Users[req.Offset:] {
		if req.Limit == 0 {
			break
		}
		name := strings.ToLower(val.Name)
		surname := strings.ToLower(val.Surname)
		sear := strings.ToLower(req.Search)
		if (strings.HasPrefix(name, sear) || strings.HasPrefix(surname, sear)) && (DataFilter(req.FromDate, req.ToDate, val)) {
			users = append(users, val)
		}
		req.Limit--
	}
	return users, false
}

func GenerateUser(count int) {
	for i := 0; i < count; i++ {
		Users = append(Users, models.User{
			Id:       i + 1,
			Name:     faker.FirstName(),
			Surname:  faker.LastName(),
			Birthday: faker.Date(),
		})
	}
}

func DataFilter(from string, to string, user models.User) bool {
	from_year, _ := strconv.Atoi(from[:4])
	from_month, _ := strconv.Atoi(from[5:7])
	from_day, _ := strconv.Atoi(from[8:10])

	to_year, _ := strconv.Atoi(to[:4])
	to_month, _ := strconv.Atoi(to[5:7])
	to_day, _ := strconv.Atoi(to[8:10])

	year_user, _ := strconv.Atoi(user.Birthday[:4])
	month_user, _ := strconv.Atoi(user.Birthday[5:7])
	day_user, _ := strconv.Atoi(user.Birthday[8:10])

	if year_user > from_year && year_user < to_year {
		return true
	} else if month_user > from_month && month_user < to_month {
		return true
	} else if day_user > from_day && day_user < to_day {
		return true
	}
	return false
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
