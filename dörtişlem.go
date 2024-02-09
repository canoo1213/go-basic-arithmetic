package main

import (
	"fmt"
)

func main() {
	
	var sayi1, sayi2 int
	var islem string

	fmt.Print("1. sayıyı giriniz: ")
	fmt.Scan(&sayi1)

	fmt.Print("2. sayıyı giriniz: ")
	fmt.Scan(&sayi2)

	fmt.Println("Yapılacak işlemi seçiniz: bölme(/), çarpma(*), toplama(+), çıkarma(-)")
	fmt.Scan(&islem)

	var sonuc int

	if islem == "/" {
		sonuc = sayi1 / sayi2
	} else if islem == "*" {
		sonuc = sayi1 * sayi2
	} else if islem == "+" {
		sonuc = sayi1 + sayi2
	} else if islem == "-" {
		sonuc = sayi1 - sayi2
	} else {
		fmt.Println("Hatalı işlem seçildi.")
		return
	}

	fmt.Printf("%d %s %d = %d\n", sayi1, islem, sayi2, sonuc)
}

