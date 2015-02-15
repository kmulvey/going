package japan

import "fmt"

func main() {
	places := [20]int{712, 745, 230, 200, 648, 440, 115, 913, 627, 621, 186, 222, 741, 954, 581, 193, 266, 320, 798, 745}
	fmt.Println("Total: ", calc(places[:]))
}
func calc(places []int) int {
	total := 0

	for i := 1; i < len(places); i++ {
		if places[i] > places[i-1] {
			fmt.Println("decreace from ", places[i], " to ", places[i-1])
			total += places[i] - places[i-1]
		}
	}

	return total
}
