package japan

import "fmt"

func main() {
	places := [5]int{30, 20, 20, 10}
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
