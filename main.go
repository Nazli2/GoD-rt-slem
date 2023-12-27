package main

import "fmt"

func main() {
	var sayi1 int
	var sayi2 int
	fmt.Print("sayi1'i giriniz :  ")
	fmt.Scan(&sayi1)
	fmt.Print("sayi2'yi giriniz :  ")
	fmt.Scan(&sayi2)

	var toplam int = sayi1 + sayi2
	var cikarma int = sayi1 - sayi2
	var bolme float64 = float64(sayi1) / float64(sayi2)
	var carpma int = sayi1 + sayi2

	fmt.Println("toplam: ", toplam)
	fmt.Println("cikarma", cikarma)
	fmt.Println("bolme", bolme)
	fmt.Println("carpma", carpma)

}
