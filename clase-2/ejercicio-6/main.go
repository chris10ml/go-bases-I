package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	for employee, age := range employees {
		if employee == "Benjamin" {
			fmt.Printf("La edad de %s es %d \n",employee, age)
		}
		if age > 20 {
			fmt.Printf("%s es mayor de 21, tiene: %d \n",employee, age)
		}
	}

	employees["Federico"] = 25

	fmt.Println(employees)

	delete(employees, "Pedro")

	fmt.Println(employees)
}
