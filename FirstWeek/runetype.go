package main

import "fmt"

func main() {
	/**
	String tipi aslında bir byte dizisi barındırmaktadır.
	Rune tipi aslında int32 değeri ile  aynıdır ancak rune tipindeki bir değişkeni int32 tipindeki bir değere eşitleyemezsiniz
	Bunun için tip dönüşümü yapmanız gerekiyor.
	*/

	//rune karakteri metinsel ifade içerisindeki unicode karakterleri temsil etmektedir.
	//Go'da diğer dillerde olan char ceri tipi bulunmamaktadır bunun yerine rune tipini kullanabiliriz
	//Rune tipindeki değerler tek tırnak içerisinde yazılmaktadır
	//Tek tırnak içerisinde sadece bir karakter yazabilirsiniz

	r := '€'

	// aşağıdaki kodun çıktısı => int32, 114
	fmt.Printf("%T, %v\n", r, r)

	// rune tipindeki bir değeri string olarak yazmak isterseniz formatlarken %c ifadesini kullanmanız gerekmektedir.
	fmt.Printf("As a string: %s and as a character: %c\n", r, r)

	// bir string içerisindeki karakterleri for range ile almak istediğinizde her bir karakter rune tipinde gelecektir.
	// Bu karakterleri ekrana string olarak yazdırmak için %c ifadesini kullanabilirsiniz.
	myString := "Hello Patik@"
	for _, v := range myString {
		fmt.Printf("%c", v)
	}

}
