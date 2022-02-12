package main

import (
	"fmt"
)

func canIDrink(age int) bool {

	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return false

}

//&메모리상의 주소
//*메모리상의 값을 살펴봄.
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	names := []string{"songseop", "test", "test2"}
	names = append(names, "ee")

	songseop := map[string]string{"name": "songseop", "age": "12"}
	favFood := []string{"kimchi", "ramen"}
	songseoptype := person{name: "songseop", age: 18, favFood: favFood}
	fmt.Println(songseoptype)
	for key, _ := range songseop {
		fmt.Println(key)
	}

	account := banking.Account{Owner: "songseop", Balance: 1000}
	account.Owner = "pepe"
	fmt.Println(account)

}
