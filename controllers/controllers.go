package controllers

import (
	"go-proj/database"
	m "go-proj/model"

	"github.com/gofiber/fiber/v2"
)

func SendHi(c *fiber.Ctx) error { // test auth
	return c.SendString("Hi")
}

func UserADD(c *fiber.Ctx) error { // add user
	var user m.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	database.DBConn.Create(&user)

	return c.JSON(map[string]interface{}{
		"data":    user,
		"message": "data was created ! ",
	})
}

func GetGEN(c *fiber.Ctx) error {
	var User []m.User

	database.DBConn.Find(&User)

	type UserRes struct {
		EmployeeId       int
		Name             string
		Lastname         string
		Birhtday         string
		Age              int
		Email            string
		Tel              string
		PeopleGeneration string
	}
	sum_genz := 0
	sum_geny := 0
	sum_genx := 0
	sum_babyboomer := 0
	sum_gi := 0
	var datauser []UserRes

	for _, v := range User {
		typeStr := ""
		if v.Age < 23 {
			typeStr = "GenZ"
			sum_genz += 1

		} else if v.Age >= 24 && v.Age <= 41 {
			typeStr = "GenY"
			sum_geny += 1

		} else if v.Age >= 42 && v.Age <= 56 {
			typeStr = "GenX"
			sum_genx += 1

		} else if v.Age >= 57 && v.Age <= 75 {
			typeStr = "Baby Boomer"
			sum_babyboomer += 1

		} else if v.Age > 76 {
			typeStr = "G.I. Generation"
			sum_gi += 1

		} else {
			typeStr = "ERROR TYPE"

		}

		d := UserRes{
			EmployeeId:       v.EmployeeId,
			Name:             v.Name,
			Lastname:         v.Lastname,
			Birhtday:         v.Birhtday,
			Age:              v.Age,
			Email:            v.Email,
			Tel:              v.Tel,
			PeopleGeneration: typeStr,
		}

		datauser = append(datauser, d)
		// sumAmount += v.Amount
	}

	return c.JSON(map[string]interface{}{
		"data":           datauser,
		"all_user_count": len(User),
		"sum_genz":       sum_genz,
		"sum_geny":       sum_geny,
		"sum_genx":       sum_genx,
		"sum_babyboomer": sum_babyboomer,
		"sum_gi":         sum_gi,
	})
}
func GetParm(c *fiber.Ctx) error {
	var user []m.User
	vari := c.Query("Search")
	database.DBConn.Raw("SELECT * FROM `users` WHERE employee_id = ? OR name = ? OR Lastname = ?", vari, vari, vari).Scan(&user)
	return c.JSON(user)
}
