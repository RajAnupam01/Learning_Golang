package main


// type Person struct {
// 	FirstName string
// 	LastName  string
// 	Age       int
// }

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type contact struct{
	Email string
	phone string
}

type Address struct {
	House int
	Area string
	State string
}

type Employee struct{
	Person_Detail Person
	Person_Contact contact
	Person_Adress Address
}

func main() {
	// var P1 Person
	// P1.FirstName = "Prince"
	// P1.LastName = "Aggarwal"
	// P1.Age = 24

	// fmt.Println("Person 1",P1)

	// P2 := Person {
	// 	FirstName: "Akash",
	// 	LastName: "Kumar",
	// 	Age:22,
	// }

	// fmt.Println("Person 2",P2)

	// var P3 = new(Person)
	// P3.FirstName="Ravi"
	// P3.LastName="Aggarwal"
	// P3.Age=22

	// fmt.Println("Person 3",P3)

    var employee1 Employee
	employee1.Person_Detail = Person {
		FirstName :"Prince",
		LastName : "Aggarwal",
		Age:24,
	}

	employee1.Person_Contact.phone= "2343555"
	employee1.Person_Contact.Email="prince@gmail.com"

	employee1.Person_Adress = Address {
		Area: "Delhi",
		House:12,
		State:"Jharkhand",
	}

}



