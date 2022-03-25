package main

import "fmt"

func main() {

	/*	c := Car{"Honda", "Kırmızı"}
		fmt.Printf("%#v\n", c)

		fmt.Println()
		c.ChangeColor("Mavi")
		fmt.Printf("%#v\n", c)*/

	c := Car{
		Vehicle: Vehicle{
			Name:  "Honde",
			Color: "Kırmızı",
		},
		GearboxType: "Otomatik",
	}
	fmt.Println(c)
	//Embedding sayesinde composition üzerinden change color a erişim sağlanabiliyor

	c.ChangeColor("Mavi")
	c.Print()
}

//Struct embedding
//Composition HAS-A
//Inheritance IS-A

type Car struct {
	Vehicle
	GearboxType string
}

func (c Car) Print() {
	fmt.Printf("Name : %s ,Color: %s,GearboxTye: %s", c.Name, c.Color, c.GearboxType)
}

type Ship struct {
	Vehicle
	EngineType string
}

type Vehicle struct {
	Name  string
	Color string
}

//Method receiverlar
//Buradaki c Car ifadesi bu metodun car structının metodu olduğunu gösteriyor
//c *Car pointer receiver demektir
/*func (c *Car) ChangeColor(color string) {
	c.Color = color
}*/
func (c *Vehicle) ChangeColor(color string) {
	c.Color = color
}
