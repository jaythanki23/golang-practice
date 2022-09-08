package main

import "fmt"

// FUNCTIONS
// func add(x int, y int) int {
// 	return x + y
// }

// func compute(x, y int) (add int, sub int) {
// 	add = x + y
// 	sub = x - y
// 	return
// }

// func test2(myFunc func(int) int) {
// 	fmt.Println(myFunc(7))
// }

func returnFunc(x string) func() {
	return func() { fmt.Println(x) }
}

func main() {
	// VARIABLES
	// var name string = "Jay Thanki"
	// var number int = 12

	// fmt.Println("Hi")

	// FMT
	// fmt.Printf("Number: %d ", 2)
	// fmt.Printf("Number: %b ", 2)
	// fmt.Printf("Number: %x ", 2)

	// USER INPUT
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Printf("Enter the year you were born: ")
	// scanner.Scan()
	// input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	// fmt.Printf("You will be %d years old at the end of 2022 ", 2022-input)

	// CONDITIONALS
	// x := 5 < 7 || 6 > 10
	// x := 5 < 7 && 6 > 10
	// fmt.Print(x)

	// if, else if, else
	// age := 23

	// if age >= 18 {
	// 	fmt.Print("You can ride")
	// } else if age >= 14 {
	// 	fmt.Print("You can ride with a parent")
	// } else {
	// 	fmt.Print("You cannot ride")
	// }

	// LOOPS
	// for i := 0; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// i := 0
	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// ARRAYS
	// arr := [3]int{1, 2, 3}
	// fmt.Println(arr)

	//SLICES
	// var x [5]int = [5]int{1, 2, 3, 4, 5}
	// var s []int = x[1:4]
	// fmt.Println(s)
	// fmt.Println(len(s))
	// fmt.Println(cap(s))

	// var a []int = []int{1, 2, 3, 4, 5}
	// b := append(a, 6)
	// fmt.Println(b)

	// var a []int = []int{3, 4, 7, 89, 44, 33, 4, 44, 69}
	// for i, elem := range a {
	// 	if i+1 >= len(a) {
	// 		continue
	// 	} else {
	// 		for _, elem2 := range a[i+1:] {
	// 			if elem == elem2 {
	// 				fmt.Println(elem)
	// 			}
	// 		}
	// 	}
	// }

	// MAPS
	// var mp map[string]int = map[string]int{
	// 	"apples": 3,
	// 	"banana": 6,
	// 	"pears":  2,
	// }

	// fmt.Println(mp)
	// mp["apple"] = 5
	// mp["pears"] = 3
	// fmt.Println(mp)

	// val, ok := mp["apple"]
	// fmt.Println(val, ok)

	//FUNCTIONS
	// ans := add(5, 7)
	// fmt.Println(ans)

	// ans1, ans2 := compute(7, 5)
	// fmt.Println(ans1, ans2)

	// test := func(x int) int {
	// 	return x * -1
	// }

	// test2(test)

	// FUNCTIONS CLOSURES
	// returnFunc("hello")()

}
