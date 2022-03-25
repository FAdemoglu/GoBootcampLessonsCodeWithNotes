package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := Book{Name: "My First Book"}

	j, err := json.Marshal(a) //=>içine verilen herhangi bir değeri json yapısına çevirir
	if err != nil {
		fmt.Errorf("JSON dönüştürme işlemi sırasında hata alındı : %s", err.Error())
	}
	fmt.Printf("JSON dönüşümü öncesi : %#v\n", a)
	fmt.Printf("JSON dönüşümü sonrası : %s\n", string(j))

	var request Book

	err = json.Unmarshal(j, &request) //json veriyi verilen requestteki structta dönüştürür

	if err != nil {
		fmt.Errorf("JSON okuma işlemi sırasında hata alındı : %s", err.Error())
	}

	fmt.Printf("%#v\n", request)
}

type Book struct {
	Name    string   `json:"bookName"`          //Burada biz http server yazdığımızda dönen responseda client "bookName" olarak alacak
	Reviews []string `json:"reviews,omitempty"` //reviewlar nil gelirse bunları boş bir slice a çevirir
}

//Marshall ve unmarshalldan farklı decode ve encode fonksiyonları vardır fark şudur
//encode ve decodeda json'ın yapısına bakılmaz yani bozuk bir json da dönüştürülmeye çalışılır
//fakat marshall json bütünlüğünü kontrol eder
