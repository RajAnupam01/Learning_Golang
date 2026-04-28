package main

import "fmt"

func main() {
	// day := 3

	// switch day {
	// case 1:
	// 	fmt.Println("Monday")
	// case 2:
	// 	fmt.Println("Tuesday")
	// case 3:
	// 	fmt.Println("Wednesday")
	// case 4:
	// 	fmt.Println("Thursday")
	// case 5:
	// 	fmt.Println("Friday")
	// case 6:
	// 	fmt.Println("Saturday")
	// case 7:
	// 	fmt.Println("Sunday")
	// default:
	// 	fmt.Println("Invalid Day")
	// }

	// month:= "January"

	// switch month {
	// case "January","February","March":
	// 	fmt.Println("Winter")
	// case "April","May","June":
	// 	fmt.Println("Summer")
	// default:
	// 	fmt.Println("Other Seasons")
	// }

	temp := 25

	switch  {
	case temp < 0:
		fmt.Println("Freezing.")
	case temp >= 0 && temp <10:
		fmt.Println("Cold")
	case temp >=10 && temp < 20:
		fmt.Println("Cool")
	case temp >=20 && temp <30:
		fmt.Println("Warm")
	default:
		fmt.Println("Hot")
	}

}
