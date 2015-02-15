package voting

import "fmt"

func main() {
	names := [4]string{"Toshi", "Jun", "Teru", "Chihiro"}
	votes := [4]string{"Jun", "Chihiro", "Toshi", "Toshi"}
	fmt.Println("Total: ", tally(names[:], votes[:]))
}
func tally(names []string, votes []string) string {
	var results map[string]int = make(map[string]int)

	for i := 0; i < len(names); i++ {
		// exclude voting for themselves
		if names[i] != votes[i] {
			results[votes[i]] += 1
		}
	}
	var winningIndex int = 0
	var winningName string = ""
	var tie bool = false
	for k, v := range results {
		if v > winningIndex {
			winningName = k
			winningIndex = v
			tie = false
		} else if v == winningIndex {
			tie = true
		}
	}
	if tie {
		return ""
	} else {
		return winningName
	}
}
