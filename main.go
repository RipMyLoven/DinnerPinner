package main

import (
	"DinnerPinner/advt"
	"DinnerPinner/math"
	"DinnerPinner/str"
	"fmt"
)

func main() {
	var num1, num2 int

	fmt.Print("Input num1: ")
	fmt.Scanln(&num1)
	fmt.Print("Input num2: ")
	fmt.Scanln(&num2)

	instanceInt := advt.Inti{
		Num1: num1,
		Num2: num2,
	}

	resultIntBool := math.MathBool(instanceInt.Num1, instanceInt.Num2)
	fmt.Println("Math Bool:", resultIntBool)

	resultIntPlus := math.MathPlus(instanceInt.Num1, instanceInt.Num2)
	fmt.Println("Math Plus:", resultIntPlus)

	resultIntMinus := math.MathMinus(instanceInt.Num1, instanceInt.Num2)
	fmt.Println("Math Minus:", resultIntMinus)

	resultIntDilen := math.MathDilen(instanceInt.Num1, instanceInt.Num2)
	fmt.Println("Math Dilen:", resultIntDilen)

	fmt.Println("|-----------------------------------|")

	num1 = 1
	num2 = 2

	resultIntBool = math.MathBool(num1, num2)
	fmt.Println("Math Bool:", resultIntBool)

	resultIntPlus = math.MathPlus(num1, num2)
	fmt.Println("Math Plus:", resultIntPlus)

	resultIntMinus = math.MathMinus(num1, num2)
	fmt.Println("Math Minus:", resultIntMinus)

	resultIntDilen = math.MathDilen(num1, num2)
	fmt.Println("Math Dilen:", resultIntDilen)

	fmt.Println("|-----------------------------------|")

	var name1, name2 string

	fmt.Print("Input Name1: ")
	fmt.Scanln(&name1)
	fmt.Print("Input Name2: ")
	fmt.Scanln(&name2)

	instanceStr := advt.SrtAdvt{
		Name1: name1,
		Name2: name2,
	}

	resultStrBool := str.StrBool(instanceStr.Name1, instanceStr.Name2)
	fmt.Println("Str Bool:", resultStrBool)

	resultStrInt := str.StrInt(instanceStr.Name1, instanceStr.Name2)
	fmt.Println("Str to Int:", resultStrInt)

	resultStrName := str.StrName(instanceStr.Name1, instanceStr.Name2)
	fmt.Println("Str to Name:", resultStrName)
}
