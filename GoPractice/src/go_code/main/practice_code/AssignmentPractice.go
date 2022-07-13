package main

import "fmt"

var (
	practiceValueNo1 int
	practiceValueNo2 int
	practiceValueNo3 int
)

func main() {
	practiceValueNo1, practiceValueNo2 = 60 ,13
	practiceValueNo3 = practiceValueNo1 + practiceValueNo2
	fmt.Println(practiceValueNo3)

	practiceValueNo1 += practiceValueNo2
	fmt.Println(practiceValueNo1)

	practiceValueNo1 -= practiceValueNo2 //practiceValueNo1 -= practiceValueNo2--->practiceValueNo1 = practiceValueNo1 - practiceValueNo2
	fmt.Println(practiceValueNo1)

	practiceValueNo2 = 2
	practiceValueNo1 *= practiceValueNo2
	fmt.Println(practiceValueNo1)

	practiceValueNo2 = 30
	practiceValueNo1 /= practiceValueNo2
	fmt.Println(practiceValueNo1)

	practiceValueNo2 = 16
	practiceValueNo1 %= practiceValueNo2
	fmt.Println(practiceValueNo1)

	practiceValueNo2 = 2
	practiceValueNo1 <<= practiceValueNo2
	fmt.Println(practiceValueNo1)

	practiceValueNo1 >>= practiceValueNo2
	fmt.Println(practiceValueNo1)

	practiceValueNo2 = 13
	practiceValueNo1 &= practiceValueNo2
	fmt.Println(practiceValueNo1)

	practiceValueNo1 ^= practiceValueNo2
	fmt.Println(practiceValueNo1)

	practiceValueNo1 |= practiceValueNo2
	fmt.Println(practiceValueNo1)
}