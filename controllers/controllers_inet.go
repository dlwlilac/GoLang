package controllers

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	m "go-fiber-test/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"go-fiber-test/database"
)

func Factorial(c *fiber.Ctx) error {

	numStr := c.Params("number")        // รับค่า number จาก Params
	number, err := strconv.Atoi(numStr) // แปลง string เป็นตัว int

	if err != nil { // ตรวจว่าค่าที่ได้มาเป็นตัวเลขมั้ย
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Pls Enter Number",
		})
	}
	fact := 1                      // ค่าเริ่มต้นของ Fact
	for i := 2; i <= number; i++ { // i = 2 เพราะ fact 0,1 คือ 1
		fact *= i // เอาเลขทุกตัวที่ได้มา * กัน
	}
	return c.SendString(fmt.Sprintf("%d! = %d", number, fact)) // แสดงผลตามรูปแบบ String ที่วางไว้

}

func ASCII(c *fiber.Ctx) error {

	taxID := c.Query("tax_id") // ดึงค่าของ queryParams ที่ชื่อ "tax_id"

	asciiValues := "" // ประกาศสตริงว่างเพื่อเก็บค่าของ ASCII

	for _, char := range taxID {
		// วนลูปผ่านตัวอักษรแต่ละตัวในสตริง taxID
		// ใช้ strconv.Itoa(int(char)) เพื่อแปลงค่าจากตัวเลข integer
		// เป็นสตริง และนำไปต่อท้ายใน asciiValues พร้อมกับเว้นวรรคระหว่างแต่ละค่า
		asciiValues += strconv.Itoa(int(char)) + " " // เปลี่ยนจาก rune เป็น int
	}

	response := fmt.Sprintf("%s! Your ASCII values: %s", taxID, asciiValues)

	return c.SendString(response)

}

func ValidTest(c *fiber.Ctx) error {
	user := new(m.User) // สร้างตัวแปร user จากโครงสร้าง m.User ใหม่

	// แปลงข้อมูลที่ได้รับจาก request body ไปเป็นโครงสร้าง user
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	validate := validator.New()                               // สร้างตัวแปร validate
	validate.RegisterValidation("username", validateUsername) // ส่ง username ไปเช็คที่ func validateUsername
	validate.RegisterValidation("url", validateUrl)           // ส่ง Url ไปเช็คที่ func validateUrl

	if err := validate.Struct(user); err != nil { // ตรวจสอบข้อมูลในโครงสร้าง models user
		var errorMessages []string

		for _, err := range err.(validator.ValidationErrors) {
			if err.Tag() == "username" { // ถ้าข้อผิดพลาดมาจากการตรวจสอบฟิลด์ username
				errorMessages = append(errorMessages, "ชื่อผู้ใช้: ใช้อักษรภาษาอังกฤษ (a-z), (A-Z), ตัวเลข (0-9) ตัวอักษรตัวแรกต้องเป็นพิมพ์ใหญ่ และเครื่องหมาย (_), (-) เท่านั้น เช่น Example_01")
			} else if err.Tag() == "url" { // ถ้าข้อผิดพลาดมาจากการตรวจสอบฟิลด์ url
				errorMessages = append(errorMessages, "ชื่อเว็บไซต์: 2-30 ตัวอักษรต้องเป็นตัวอักษรภาษาอังกฤษตัวเล็ก (a-z) ตัวเลข (0-9) ห้ามใช้เครื่องหมายอักขระพิเศษยกเว้นเครื่องหมายขีด (-) ห้ามเว้นวรรคและห้ามใช้ภาษาไทย ")
			} else {
				errorMessages = append(errorMessages, err.Error())
			}
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errorMessages,
		})
	}
	response := map[string]interface{}{
		"message": "Register Success",
		"user":    user,
	}
	return c.JSON(response)

}

func validateUrl(fl validator.FieldLevel) bool {
	url := fl.Field().String()                      // ดึงค่า Url
	regex := regexp.MustCompile(`^[a-zA-Z0-9-.]+$`) // เงื่อนไขตรวจสอบ
	return regex.MatchString(url)                   // ตรวจสอบ Url ว่าตรงกับเงื่อนไขไหม
}

func validateUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()                      // ดึงค่า Username
	regex := regexp.MustCompile(`^[A-Z][a-zA-Z0-9_-]+$`) // เงื่อนไขตรวจสอบ
	return regex.MatchString(username)                   // ตรวจสอบ Username ว่าตรงกับเงื่อนไขไหม
}

// * Start Dogs //

func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //delelete = null
	return c.Status(200).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

func AddDog(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetDogsJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) // 10 ตัว

	var pinkCount, greenCount, redCount, nocolor int

	var dataResults []m.DogsRes
	for _, v := range dogs {
		typeStr := ""
		if v.DogID >= 200 && v.DogID <= 251 {
			typeStr = "pink"
			pinkCount++
		} else if v.DogID >= 100 && v.DogID <= 150 {
			typeStr = "green"
			greenCount++
		} else if v.DogID >= 10 && v.DogID <= 50 {
			typeStr = "red"
			redCount++
		} else {
			typeStr = "no color"
			nocolor++
		}

		d := m.DogsRes{
			Name:  v.Name,
			DogID: v.DogID,
			Type:  typeStr,
		}
		dataResults = append(dataResults, d)
	}

	r := m.ResultData{
		Data:  dataResults,
		Name:  "golang-test",
		Count: len(dogs),
		ColorCounts: m.ColorCounts{
			Pink:    pinkCount,
			Green:   greenCount,
			Red:     redCount,
			NoColor: nocolor,
		},
	}
	return c.Status(200).JSON(r)
}

func GetDogsDel(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Unscoped().Where("deleted_at IS NOT NULL").Find(&dogs) //delelete = !null
	return c.Status(200).JSON(dogs)
}

func Getscope(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Scopes(DogIDGreaterThan50_100).Find(&dogs)
	return c.Status(200).JSON(dogs)
}

func DogIDGreaterThan50_100(db *gorm.DB) *gorm.DB {
	return db.Where("dog_id > ?  AND dog_id < ?", 50, 100)
}

// * Eng Dogs //

// * Start Company //

func Addcompany(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var Company m.Company

	if err := c.BodyParser(&Company); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&Company)
	return c.Status(201).JSON(Company)
}

func Updatecompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company m.Company
	id := c.Params("id")

	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&company)
	return c.Status(200).JSON(company)
}

func Removecompany(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var company m.Company

	result := db.Delete(&company, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.SendStatus(200)
}

func Getcompanys(c *fiber.Ctx) error {
	db := database.DBConn
	var company []m.Company

	db.Find(&company) //delelete = null
	return c.Status(200).JSON(company)
}

// * End Company //
