package main

import "fmt"

//Anonym structlar bir fonksiyon içinde tanımlanan etk seferlik kulanılan
//Yaşam alanı o fonksiyon içerisinde kalan structlardır
func main() {
	var a = struct {
		FirstName string
	}{
		"My first Name",
	}

	var b = struct {
		FirstName string
	}{
		"My first Name",
	}

	a = b //atama yapılabilir çünkü  anonym structlar ayrıca alanları ve sıraları aynı ex:Firstname

	fmt.Println(a)
	fmt.Println(c)
}
