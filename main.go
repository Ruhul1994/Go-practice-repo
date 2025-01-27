package main
import "fmt"
// func main () {
// 	 a := 10
// 	 b := 20
// 	c := a + b
// 	fmt.Println("The sum of a and b is: ", c);
// 	fmt.Println("Thank you Ruhul you did it")
// }

func main() {
	fmt.Println("starting the program...")
	great("Ruhul")
}

func great(name string) {
	fmt.Printf("Hello, %s\n", name)
}