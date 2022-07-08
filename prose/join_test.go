package prose

import (
	"fmt"
	"testing"
)

func TestOneElement2(t *testing.T) {
	// Arrange
	//
	list := []string{"apple"}
	want := "apple"

	// Act
	//
	got := JoinWithCommas(list)

	// Assert
	//
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func TestTwoElements2(t *testing.T) {
	// Arrange
	//
	list := []string{"apple", "orange"}
	want := "apple and orange"

	// Act
	//
	got := JoinWithCommas(list)

	// Assert
	//
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func TestThreeElements2(t *testing.T) {
	// Arrange
	//
	list := []string{"apple", "orange", "pear"}
	want := "apple, orange, and pear"

	// Act
	//
	got := JoinWithCommas(list)

	// Assert
	//
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
}
