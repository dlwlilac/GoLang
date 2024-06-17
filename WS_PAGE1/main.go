package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Company struct {
	Name    string
	Address string
	Phone   string
	Email   string
}

func power(base, exponent float64) float64 {
	return math.Pow(base, exponent)
}

func main() {

	// num0()
	// num1()
	// num12()
	// num2()
	// num3()
	// num31()
	// num4()
	// num41()
	// num5()
	num6()
	// spacial()
}

func num0() {
	i := 2

	fmt.Println("Example: If condition")

	if i == 0 {
		fmt.Println("Zero")
	} else if i == 1 {
		fmt.Println("One")
	} else if i == 2 {
		fmt.Println("Two")
	} else if i == 3 {
		fmt.Println("Three")
	} else {
		fmt.Println("Your i not in case.")
	}
}

func num1() {
	count := 0 // เพิ่มตัวแปรนับจำนวน

	for i := 1; i < 100; i++ {
		if i%3 == 0 {
			fmt.Print(i, ",")
			count++ // เพิ่มตัวนับเมื่อพบตัวเลขที่หารด้วย 3 ลงตัว
		}
	}

	fmt.Printf("\nTotal: %d\n", count) // แสดงผลรวมของจำนวนตัวเลขที่หารด้วย 3 ลงตัว

}

func num12() {
	result := power(20, 2)
	fmt.Println("20,2 =", result)
}

func num2() {
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}

	minValue, maxValue := findMinMax(x)

	fmt.Printf("Minimum number is: %d\n", minValue)
	fmt.Printf("Maximum number is: %d\n", maxValue)
}

// หาเลขน้อยกว่ามากกว่า
func findMinMax(arr []int) (int, int) {
	if len(arr) == 0 {
		return 0, 0
	}

	min := arr[0]
	max := arr[0]

	for _, num := range arr {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return min, max
}

func num3() {
	count := countNineIn1000(1, 1000)
	fmt.Println("Total number of 9 from 1 to 1000 =", count)
}

func countNineIn1000(start, end int) int {
	count := 0
	for i := start; i <= end; i++ {
		numStr := strconv.Itoa(i)
		for _, char := range numStr {
			if char == '9' {
				count++
			}
		}
	}
	return count
}

func num31() {
	count := someFunc(10000)
	fmt.Println("Total number of 9 from 1 to 10000 =", count)
}

func someFunc(end int) int {
	count := 0
	for i := 1; i <= end; i++ {
		count += countNumber(i)
	}
	return count
}

func countNumber(num int) int {
	count := 0
	numStr := strconv.Itoa(num)
	for _, char := range numStr {
		if char == '9' {
			count++
		}
	}
	return count
}

func num4() {
	var myWords = "AW SOME GO!"
	var result string

	for i := 0; i < len(myWords); i++ {
		if myWords[i] != ' ' { // เช็คว่าตัวอักษรไม่ใช่ช่องว่าง
			result += string(myWords[i])
		}
	}

	fmt.Println(result)

}

func num41() {
	text := "ine t"
	result := cutText(text)
	fmt.Println("Result:", result)
}

func cutText(input string) string {
	return strings.ReplaceAll(input, " ", "")
}

func num5() {
	peoples := map[string]map[string]string{
		"data_01": {
			"FirstName": "William",
			"LastName":  "YaHhhh",
		},
		"data_02": {
			"FirstName": "The",
			"LastName":  "Flood",
		},
		"data_03": {
			"FirstName": "Shin",
			"LastName":  "Godzilla",
		},
		"data_04": {
			"FirstName": "Yasuo",
			"LastName":  "Yone",
		},
	}
	for key, value := range peoples {
		fmt.Printf("%s:\n", key)
		for field, data := range value {
			fmt.Printf("  %s: %s\n", field, data)
		}
		fmt.Println()
	}

}

func num6() {
	company := Company{
		Name:    "Khao Man Kai",
		Address: "KFC 5Star Mc",
		Phone:   "1112 1111 2222",
	}
	fmt.Println("Company Name:", company.Name)
	fmt.Println("Company Address:", company.Address)
	fmt.Println("Company Phone:", company.Phone)

}

func spacial() {
	// fmt.Println("*")
	// fmt.Println("**")
	// fmt.Println("***")
	// fmt.Println("****")
	// fmt.Println("*****")
	// fmt.Println("******")

	x := 6

	for i := 1; i <= x; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
