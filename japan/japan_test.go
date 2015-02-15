package japan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcOne(t *testing.T) {
	places := [4]int{30, 20, 20, 10}
	assert.Equal(t, 0, calc(places[:]), "they should be equal")
}
func TestCalcTwo(t *testing.T) {
	places := [3]int{5, 7, 3}
	assert.Equal(t, 2, calc(places[:]), "they should be equal")
}
func TestCalcThree(t *testing.T) {
	places := [9]int{6, 8, 5, 4, 7, 4, 2, 3, 1}
	assert.Equal(t, 6, calc(places[:]), "they should be equal")
}
func TestCalcFour(t *testing.T) {
	places := [10]int{749, 560, 921, 166, 757, 818, 228, 584, 366, 88}
	assert.Equal(t, 2284, calc(places[:]), "they should be equal")
}
func TestCalcFive(t *testing.T) {
	places := [20]int{712, 745, 230, 200, 648, 440, 115, 913, 627, 621, 186, 222, 741, 954, 581, 193, 266, 320, 798, 745}
	assert.Equal(t, 6393, calc(places[:]), "they should be equal")
}
