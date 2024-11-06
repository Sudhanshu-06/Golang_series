package main

import "fmt"

// for -> only construct in go for looping
func main(){
	// while loop
	// i:=1;
	// for i<=3{
	// 	fmt.Println(i)
	// 	i=i+1
	// }

	// infinite loop

	// for{
	// 	fmt.Println("1")
	// }

	// classic for loop
// 	for i:=0; i<=5; i++{
// 		fmt.Println(i)
// 	}

// break and continue
// for i:=1; i<10; i++{
// 	 if i==2{
// 	/	continue;
// 	}
// 	if i==4{
// 		break;
// 	}
// 	fmt.Println(i)
// }

// range

for i:=range 4{
	fmt.Println(i)
}

}