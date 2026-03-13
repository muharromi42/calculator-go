package main

import "fmt"

func main() {
	var angka_pertama float64
	var angka_kedua float64
	var operator string
	var result float64

	fmt.Println("masukkan angka pertama :")
	fmt.Scanln(&angka_pertama)
	fmt.Println("masukkan operator :")
	fmt.Scanln(&operator)
	fmt.Println("masukkan angka kedua :")
	fmt.Scanln(&angka_kedua)

	switch operator {
	case "+":
		result = angka_pertama + angka_kedua
		fmt.Println("hasil opearsi :", result)
	case "-":
		result = angka_pertama - angka_kedua
		fmt.Println("hasil opearsi :", result)
	case "/":
		result = angka_pertama / angka_kedua
		fmt.Println("hasil opearsi :", result)
	case "*":
		result = angka_pertama * angka_kedua
		fmt.Println("hasil operasi :", result)
	default:
		fmt.Println("data tidak valid")
	}
}
