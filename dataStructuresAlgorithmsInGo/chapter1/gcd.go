package main

func greatestCommonDivesor(a, b int) int {
	if a > b {
		return greatestCommonDivesor(b, a)
	}
	if b%a == 0 {
		return a
	}
	return greatestCommonDivesor(a, b%a)
}

//
// func main() {
// 	fmt.Println(33, 999, greatestCommonDivesor(33, 999))
// 	fmt.Println(1998, 98999, greatestCommonDivesor(1998, 9899))
// 	fmt.Println(7, 8, greatestCommonDivesor(7, 8))
// }
