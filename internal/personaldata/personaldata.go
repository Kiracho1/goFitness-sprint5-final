package personaldata

import "fmt"
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

func (p Personal) Print() {
	fmt.Printf(
	"Имя: %s\n" +
	"Вес: %.2f\n" +
	"Рост: %.2f\n",
	p.Name, p.Weight, p. Height)
}
