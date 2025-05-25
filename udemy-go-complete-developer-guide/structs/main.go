package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// You can have a struct within a struct :D
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// Notice that we explicitly state which field within the struct we want followed by the value.
	// This is pretty good practice rather than just specifying the value, as this avoids the possibility of having to constantly ensure the proper
	// order of the supplied values to the struct on creation.
	bruce := person{
		firstName: "Bruce",
		lastName:  "Wayne",
		contactInfo: contactInfo{
			zipCode: 999999,
			email:   "imbatman@gmail.com",
		},
	}

	bruce.print()
}

func (p person) print() {
	fmt.Printf("%v.\n", p)
}

// for the * in *person receiver, we should treat it as a form of description rather than an outright operator. basically describes
// that this receiver expects a pointer type value to be used
func (pointerToPerson *person) updateName(newFirstName string) {
	// *pointerToPerson asks for the value at that particularly given address in memory.
	// reason for doing it in this fashion is because the person type that we are using in this function is a `value type`.
	// In golang, they follow a pass-by-value mechanisms to which any variables passed into functions have a copy made in memory.
	// A `value type` variable simply contains a value. However, a `reference type` such as a slice, contains a reference to the actual address in memory
	// Hence you will notice that for `value type`, we would want to pass in a pointer to that variable whereas for `reference type` there isn't
	// a need to do so as even though it is copied in memory, the reference type value innately already contains a reference to the actual
	// address of the value we wish to modify.
	(*pointerToPerson).firstName = newFirstName
}
