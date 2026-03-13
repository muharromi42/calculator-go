package main

import "fmt"

func main() {
	var angkaPertama float64
	var angkaKedua float64
	var operator string
	var result float64

	fmt.Println("masukkan angka pertama :")
	fmt.Scanln(&angkaPertama)
	fmt.Println("masukkan operator :")
	fmt.Scanln(&operator)
	fmt.Println("masukkan angka kedua :")
	fmt.Scanln(&angkaKedua)

	switch operator {
	case "+":
		result = angkaPertama + angkaKedua
	case "-":
		result = angkaPertama - angkaKedua
	case "/":
		result = angkaPertama / angkaKedua
	case "*":
		result = angkaPertama * angkaKedua
	default:
		fmt.Println("data tidak valid")
	}

	fmt.Println(angkaPertama, operator, angkaKedua, "=", result)
}
