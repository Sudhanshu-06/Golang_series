package main

import "fmt"

// maps-> object,dict
func main() {
// creating map

    // m := make(map[string]string)

	// setting an element
	// m["name"]="golang"
	// m["area"]="backend"

	// get an element
	// fmt.Println(m["name"],m["area"])
	// if key does not exist in the map then it returns zero value
	// fmt.Println(m["phone"])

	// m:= make(map[string]int)
	// m["age"] =30
	// m["price"] =300
	// fmt.Println(m["phone"])

	// fmt.Println(len(m))

	// delete(m,"price")
	// fmt.Println(m)

	// clear(m)
	// fmt.Println(m)

	// m:= map[string]int{"price":40,"phone":3}
	// fmt.Println(m)

	m:= map[string]int{"price":40,"phone":3}
	
	_,ok:= m["price"]

	if ok{
		fmt.Println("all ok")
	}else{
		fmt.Println("not ok")
	}




    
}