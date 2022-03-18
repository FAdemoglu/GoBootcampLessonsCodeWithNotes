package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	/**
	error tipi Go dili içerisinde hatayı temsil eder
	Go'da fonksiyonlar birden çok fazla dönüş değerinse sahip olabilir.
	Bir çok error fonksiyon 2. ya da 3. dönüş tipi olarak tipini kullanır


	*/

	// strconv paketi string bir değeri float64,int,int32 vb. bir veri tipine dönüştürmek için kullanılır.
	// Buradaki Atoi kısaltması ise "ASCII to Int" ifadesini temsil eder.

	// strconv paketini kullanarak parametrede gönderdiğimiz değeri int tipine çevirmek istiyoruz.
	// Eğer bu değer geçerli bir int değeri ise geriye ilk dönüş tipi olarak integer (tam sayı) tipindeki değer ikinci dönüş değeri ise error tipinde hata olup olmadığını kontrol etmemizi sağlayacak değer geliyor.

	// Hatanın nil olmaması için "deneme" değerini parametrede kullanabilirsiniz.
	convertedValue, err := strconv.Atoi("1")

	// Go'da bir try-catch yapısı olmadığından dolayı hata yönetimi aşağıdaki şekilde yapılmaktadır.
	// Burada eğer error değerini atadığımız değişken nil yani null değil ise ekrana hatayı yazdırıp return ediyoruz.
	if err != nil {
		fmt.Println(err)
		return
	}

	// Eğer verdiğimiz değer geçerli bir değerse dönüştürülen değeri ekrana yazdırıyoruz.
	fmt.Println(convertedValue)

	// Tüm bunların yanı sıra kendi özel hatalarınızı da üretebilirsiniz.
	// bunun için errors.New() metodunu kullanarak özel hatalarınızı projenizde kullanabilirsiniz.

	customErr := errors.New("Benim özel hata mesajım")

	fmt.Println(customErr)

	// Bir diğer yöntem ise fmt.Errorf() bu da hata mesajlarınızı formatlayarak oluşturmanızı sağlar ve geriye error tipinde değer döner

	fmtErr := fmt.Errorf("Benim özel hata mesajım %d", 2)
	fmt.Println(fmtErr)

	// Son olarak eğer fmt.Errorf kullanmadan hata mesajlarınızı formatlamak isterseniz
	// errors.New içerisinde fmt.Sprintf metodunu kullanarak mesajlarınızı formatlayabilirsiniz.
	lastErr := errors.New(fmt.Sprintf("Benim özel hata mesajım %d", 3))

	fmt.Println(lastErr)
}
