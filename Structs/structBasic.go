package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact  contactInfo
}

func main() {
 
	// alex := person{"Alex", "Anderson"} One way of defining

	// alex := person{firstName:"Alex", lastName:"Anderson"} 
	// fmt.Println(alex)

	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex) //Empty {} braces is the op
	fmt.Printf("%+v", alex) //{firstName:Alex lastName:Anderson}

	jim := person{
		firstName: "Jim",
		lastName: "Parker",
		contact: contactInfo {
			email: "jim1994@gmail.com",
			zipCode: 988780,
		},
	}

	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")
	jim.updateName("Jimmy") //Pointer Shortcut

	// jim.updateName("jimmy") //The name does not get updated to jimmy. GO is a pass by value lang.
	jim.print() 

}

func (pointerToPerson *person) updateName(newName string) { //pointer to person
	(*pointerToPerson).firstName = newFirstName
}


func (p person) print() {
	// fmt.Printf("%+v", jim)
	fmt.Printf("%+v", p)
}