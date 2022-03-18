package main

import (
	"fmt"
	"strconv"
)

func main() {
	//String ifadeler bir byte array olarak saklanmaktadır
	//Bir sayıyı string bir ifadeye çevirmek için string fonksiyoun kullanabilirsiniz

	s := string(100)

	fmt.Println(s)

	// Ancak bu işlemi strconv paketi ile yaptığınızda parametre olarak verilen int değerini olduğu gibi string'e çevirecek
	myVar := strconv.Itoa(100)
	fmt.Println(myVar)
}
