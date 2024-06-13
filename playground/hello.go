package main

import "fmt"

type Human struct {
	Name   string
	Age    int
	Weight float32
}

func main() {
	scooter := Human{
		Name:   "Ryuichiro Suzuki",
		Age:    24,
		Weight: 60.3,
	}
	message := fmt.Sprintf("hi %s, you are %d years old and weigh %.1f kg", scooter.Name, scooter.Age, scooter.Weight)
	fmt.Println(message)
	fmt.Println(addOne(100))

	scooter.addWeight(4.4)

	message = fmt.Sprintf("hi %s, you are %d years old and weigh %.1f kg", scooter.Name, scooter.Age, scooter.Weight)
	fmt.Println(message)

	prices := map[string]int{"udon": 400, "ramen": 800}
	fmt.Println(prices["soba"])

	m := map[string]string{"udon": "400", "ramen": "800"}
	fmt.Println(m["soba"])

	scooter.rename("Nathan Cohen")

	fmt.Println(scooter)
}

func addOne(i int) int {
	return i + 1
}

func (h *Human) addWeight(f float32) {
	h.Weight += f
}

func (h *Human) rename(name string) Human {
	h.Name = name
	return h
}
