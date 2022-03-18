package main

import "fmt"

/**
Constantlar hem paket seviyesinde hem de fonksiyon içerisinde tanımlanabilir

Constant tipler sadece string numeric ve bool tipinde değişkenler olabilir

*/

const myVar = "message" //untyped

const typedConst string = "message" //typed

const name string = "Patika"

/**
Go da typed ile bir değişken tanımlarsanız açık bir şekilde verdiğiniz için
kendi string tipinden ürettiğiniz bir tip olsa bile o tip ile bir atama yapamazsınız.

*/

/**

 */

func main() {
	changeName(name) //=>Pass by value şeklinde bir gönderim yapılır
	fmt.Println(name)
}

func changeName(name string) {
	name += ".dev"
	fmt.Println(name)
}
