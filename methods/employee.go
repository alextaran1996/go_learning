package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Employee struct {
	Name   string
	Age    int
	Height int
	Weight int
	Gender string
}

type Age struct {
	age int
}

func (empl Employee) addNew(exsworker []Employee) []Employee {
	regex := regexp.MustCompile(`^[a-zA-Z]*$`) // fixed regexp. Name begins and ends with letters from a-Z and can contain multiple letters
	if !regex.MatchString(empl.Name) {
		fmt.Println("Name shouldn't include non-letter values")
		os.Exit(1)
	}
	if strings.ToLower(empl.Gender) != "m" && strings.ToLower(empl.Gender) != "w" {
		panic("Gender has only M or W values") // realized panic instead of os.Exit
	}
	exsworker = append(exsworker, empl)
	return exsworker
}

func (a *Age) addNew() {
	a.age++
}

/* Function for removing employee.Using pointer cuz function copy variable and work with it and
size of structure variable it quiet big
*/
func (empl *Employee) del(exsworker []Employee) []Employee { // (empl *Employee) called reciver
	for ind, val := range exsworker {
		if val == *empl {
			copy(exsworker[ind:], exsworker[ind+1:]) // Move all elements from the right side on desired position
			return exsworker[:len(exsworker)-1]      // Don't include last element of the slice as it will repeat twice
		}
	}
	return exsworker
}

// Add new structure and method for it that will have overlapping names with Employee method strcture

func main() {
	// Two different structures can have  identical method name but reciver in function should be different
	alex := Employee{Name: "Aliaksandr", Age: 12, Height: 159, Weight: 60, Gender: "M"}
	tim := Employee{Name: "Tim", Age: 28, Height: 180, Weight: 90, Gender: "M"}
	greg := Employee{Name: "Greg", Age: 65, Height: 209, Weight: 80, Gender: "M"}
	workers := alex.addNew(nil)
	workers = tim.addNew(workers)
	workers = greg.addNew(workers)
	workers = tim.del(workers)
	old := Age{42}
	old.addNew() // the same as (*old).addNew().Here is compiler implicitly gets variable address
	fmt.Println(old)
	fmt.Println(workers)

}
