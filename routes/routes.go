package routes

import (
	"go-proj/database"
	m "go-proj/model"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"gorm.io/gorm"
)

func Router(app *fiber.App) {

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v2 := api.Group("/v2")
	v1.Use(basicauth.New(basicauth.Config{Users: map[string]string{
		"testgo": "6102565",
	}}))
	v1.Get("/test", func(c *fiber.Ctx) error { // test auth
		return c.SendString("Hi")
	})
	v1.Post("/User", func(c *fiber.Ctx) error { // add user
		var user m.User

		if err := c.BodyParser(&user); err != nil {
			return c.Status(503).SendString(err.Error())
		}
		database.DBConn.Create(&user)

		return c.JSON(map[string]interface{}{
			"data":    user,
			"message": "data was created ! ",
		})
	})
	v2.Get("/Data", func(c *fiber.Ctx) error {
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
	})
	v2.Get("/DataP", func(c *fiber.Ctx) error {
		var user []m.User

		type getSearch struct {
			data string
		}

		getSG := new(getSearch)
		getSG.data = c.Query("Search")
		database.DBConn.Scopes(SrchEidNmeLnme(getSG)).Find(&user)
		return c.SendString("")
	})
}

func SrchEidNmeLnme(v interface{}) func(*gorm.DB) {
	var arr = make([]func(*gorm.DB) *gorm.DB, 0)
}
func customStructFilter(v interface{}) []func(*gorm.DB) *gorm.DB {
	var arr = make([]func(*gorm.DB) *gorm.DB, 0)
	vl := reflect.ValueOf(v).Elem()
	typ := reflect.TypeOf(v).Elem()
	for idx := 0; idx < vl.NumField(); idx++ {
		if !vl.Field(idx).IsNil() {
			arr = append(arr, func(d *gorm.DB) *gorm.DB {
				return d.Where(typ.Field(idx).Tag.Get("qrstr"), vl.Field(idx).Interface())
			})
		}
	}
	return arr
}
