package repository

import (
	"fibonacci"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetSequence(t *testing.T) {
	//Arrange
	testTable := []struct {
		input    *fibonacci.Input
		expected map[int64]string
	}{
		{input: &fibonacci.Input{Start: -1, End: 6}, expected: nil},
		{input: &fibonacci.Input{Start: 0, End: 10001}, expected: nil},
		{input: &fibonacci.Input{Start: 4, End: 7}, expected: map[int64]string{4: "3", 5: "5", 6: "8", 7: "13"}},
	}
	//Act
	for _, testCase := range testTable {
		rdb, _ := NewRedisClient(Config{
			Addr:     ":6379",
			DB:       0,
			Password: "",
		})
		repository := NewRepository(rdb)
		result := repository.GetSequence(testCase.input)
		//Assert
		if !cmp.Equal(result, testCase.expected) {
			t.Error("Incorrect result. Expect: ", testCase.expected, " got: ", result)
		}
	}

}
