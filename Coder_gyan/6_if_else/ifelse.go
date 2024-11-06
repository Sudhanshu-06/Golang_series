package main

import "fmt"

func main() {
	// age:=16

	// if age >= 18{
	// 	fmt.Println("person age is valid for voting")
	// }else{
	// 	fmt.Println("person age is not valid for voting")
	// }

	// age := 16

	// if age >= 18 {
	// 	fmt.Println("person age is valid for voting")
	// } else if age >= 12 {
	// 	fmt.Println("person age is teenager")
	// } else {
	// 	fmt.Println("person age is kid")
	// }

	// var role="admin"
	// var hasPermission = false

	// 	 if role=="admin" || hasPermission {
	// 	 	 fmt.Println("Access granted")
	// 	  }

	// 	if role=="admin" && hasPermission {
	// 		fmt.Println("Access granted")
	// 	}

	// we can declare a variable inside if construct

	if age :=15; age >=18{
		fmt.Println("person age is valid for voting",age)
	}else if age>=12{
		fmt.Println("person is teengaer",age)
	}else{
		fmt.Println("person is kid",age)
	}

	// go lang does not have ternary operator you will have to use normal if else 

}
