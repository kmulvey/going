package voting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVoteOne(t *testing.T) {
	names := [4]string{"Toshi", "Jun", "Teru", "Chihiro"}
	votes := [4]string{"Jun", "Chihiro", "Toshi", "Toshi"}
	assert.Equal(t, "Toshi", tally(names[:], votes[:]), "Toshi should win")
}
func TestVoteTwo(t *testing.T) {
	names := [4]string{"Toshi", "Jun", "Teru", "Chihiro"}
	votes := [4]string{"Teru", "Teru", "Jun", "Jun"}
	assert.Equal(t, "", tally(names[:], votes[:]), "they should be equal")
}
func TestVoteThree(t *testing.T) {
	names := [4]string{"Toshi", "Jun", "Teru", "Chihiro"}
	votes := [4]string{"Toshi", "Toshi", "Jun", "Jun"}
	assert.Equal(t, "Jun", tally(names[:], votes[:]), "they should be equal")
}
func TestVoteFour(t *testing.T) {
	names := [2]string{"Toshi", "Jun"}
	votes := [2]string{"Toshi", "Jun"}
	assert.Equal(t, "", tally(names[:], votes[:]), "they should be equal")
}
func TestVoteFive(t *testing.T) {
	names := [4]string{"PhpLove", "phplove", "phpLove", "Phplove"}
	votes := [4]string{"phpLove", "phpLove", "phpLove", "PhpLove"}
	assert.Equal(t, "phpLove", tally(names[:], votes[:]), "they should be equal")
}
