package main

import "fmt"

func main() {
	var angkaPertama float64
	var angkaKedua float64
	var operator string
	var result float64

	fmt.Println("masukkan angka pertama :")
	_, err := fmt.Scanln(&angkaPertama)
	if err != nil {
		fmt.Println("input tidak valid harus berupa angka!", err)
		return
	}

	fmt.Println("masukkan operator (+, -, *, /) :")
	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Println("input tidak valid harus berupa operator matematika (+, -, *, /)", err)
		return
	}

	fmt.Println("masukkan angka kedua :")
	_, err = fmt.Scanln(&angkaKedua)
	if err != nil {
		fmt.Println("input tidak valid harus berupa angka!", err)
		return
	}

	switch operator {
	case "+":
		result = angkaPertama + angkaKedua
	case "-":
		result = angkaPertama - angkaKedua
	case "/":
		if angkaKedua == 0 {
			fmt.Println("tidak bisa membagi dengan angka 0")
			return
		}
		result = angkaPertama / angkaKedua
	case "*":
		result = angkaPertama * angkaKedua
	default:
		fmt.Println("data tidak valid")
		return
	}

	fmt.Println(angkaPertama, operator, angkaKedua, "=", result)
}
