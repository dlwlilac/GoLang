package controllers

import (
	"fmt"
	"regexp"
	"strconv"

	m "go-fiber-test/models"

	// "github.com/go-playground/validator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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
	return c.SendString(fmt.Sprintf("%d! = %d", number, fact))

}

func ASCII(c *fiber.Ctx) error {

	taxID := c.Query("tax_id")
	asciiValues := ""
	for _, char := range taxID {
		asciiValues += strconv.Itoa(int(char)) + " "
	}

	response := fmt.Sprintf("%s! Your ASCII values: %s", taxID, asciiValues)

	return c.SendString(response)

}

func ValidTest(c *fiber.Ctx) error {
	user := new(m.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	validate := validator.New()
	validate.RegisterValidation("username", validateUsername)
	validate.RegisterValidation("url", validateUrl)

	if err := validate.Struct(user); err != nil {
		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			if err.Tag() == "username" {
				errorMessages = append(errorMessages, "ชื่อผู้ใช้: ใช้อักษรภาษาอังกฤษ (a-z), (A-Z), ตัวเลข (0-9) และเครื่องหมาย (_), (-) เท่านั้น เช่น Example_01")
			} else if err.Tag() == "url" {
				errorMessages = append(errorMessages, "ชื่อเว็บไซต์: 2-30 ตัวอักษรต้องเป็นตัวอักษรภาษาอังกฤษตัวเล็ก (a-z) ตัวเลข (0-9) ห้ามใช้เครื่องหมายอักขระพิเศษยกเว้นเครื่องหมายขีด (-) ห้ามเว้นวรรคและห้ามใช้ภาษาไทย ")
			} else {
				errorMessages = append(errorMessages, err.Error())
			}
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errorMessages,
		})
	}
	return c.JSON(user)

}

func validateUrl(fl validator.FieldLevel) bool {
	url := fl.Field().String()
	regex := regexp.MustCompile(`^[a-zA-Z0-9-.]+$`)
	return regex.MatchString(url)
}

func validateUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	if username == "" {
		return true
	}
	regex := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	return regex.MatchString(username)
}
