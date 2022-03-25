package main

import "fmt"

func main() {

	//a := Person{"asdads", "adasda", 12, 23}

	//b := Person2{"asdads", "adasda", 12, 23}

	//a=b bu yapılamaz
	//a = Person(b) //=> bu dönüşüm yapılabilir ta ki  struct alanlarının sırası ve değişken tipleri aynı olduğu sürece

	b := new(Person) //bu şekilde yapıldığına bu structın pointerının döner
	c := Person{
		FirstName: "asadfs",
		LastName:  "ewfsdf",
		Age:       23,
		Salary:    30,
	}

	//fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Salary    float64
}

//Yukarıdaki person structında alanların yerleri değişirse structın kapladığı
//alan da değişebiliyor bunun sebebi memory padding
/*
type Person2 struct {
	FirstName string
	LastName  string
	Age       int
	Salary    float64
}*/
