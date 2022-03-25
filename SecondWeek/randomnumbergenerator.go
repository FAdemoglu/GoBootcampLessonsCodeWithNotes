package main

import (
	"fmt"
	//"math/rand"//math rand random generator için iyi bir kütüphane değil
	"crypto/rand"
	"time"
)

//Burada sürekli aynı sayıyı verir random bir sayı sürekli gelmez
//sebebi go da sudo random generator denen bir kavram var
//karakteristik olarak seed değeri var seed değeri aynı olduğu sürece
//üretilen sayı hep aynı olacak

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	rand.Seed(time.Now().Unix()) //Burada seed değeri sürekli değiştiği için
	//sürekli farklı sayı ekrana yazar
	a := random(0, 100)
	fmt.Println(a)
}
