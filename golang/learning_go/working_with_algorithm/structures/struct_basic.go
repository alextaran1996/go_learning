package main

import "fmt"

type employee struct { //Unexported structure for employee
	ID, Salary, Age                 int
	Name, Surname, Position, Addres string
}

type pos struct {
	x, y int
}

type round struct { // Embedded structure  for position strcture
	pos    // Anonimous field. Give an opportunity to escape from specifying all the embedded structures when assign value to a variable
	radius int
}

type wheel struct { // Embedded structure for round structure embedded in position structure
	round // anonimous field
	spoke int
}

func risesalary(empl *employee, coef int) { // Function thath doubles employees' salary
	empl.Salary = empl.Salary * coef
}

func main() {
	var alex = employee{ID: 321145, Name: "Alex", Surname: "Torv", Age: 22, Salary: 200, Position: "Cleaner"} // New employee
	var rick employee                                                                                         // New employee2
	rick.ID = 1241552                                                                                         // Set ID for rick
	position := &alex.Position
	fmt.Println(alex, "Senior "+*position)
	risesalary(&alex, 2) // Double alexs' salary
	fmt.Println(alex.Salary)
	var example wheel
	example.x = 2
	fmt.Println(example)
	var example2 = wheel{round{pos{12, 12}, 4}, 5} // Create elemenet of the wheel structure that uses embedded structures
	fmt.Println(example2)

}
