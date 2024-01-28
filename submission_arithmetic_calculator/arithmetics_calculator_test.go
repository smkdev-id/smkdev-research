// ! DON'T CHANGE ANYTHING HERE OR THE RESULT WILL NOT DEFINED

// ! DON'T CHANGE ANYTHING HERE OR THE RESULT WILL NOT DEFINED

package main

import (
	"testing"
)

func TestArrangement(t *testing.T) {
	actual := ArithmeticArranger([]string{"3 + 855", "3801 - 2", "45 + 43", "123 + 49"})
	expected := "    3      3801      45      123\n+ 855    -    2    + 43    +  49\n-----    ------    ----    -----"
	if actual != expected {
		t.Errorf("Expected different output when calling ArithmeticArranger() with test input")
	}

	actual = ArithmeticArranger([]string{"11 + 4", "3801 - 2999", "1 + 2", "123 + 49", "1 - 9380"})
	expected = "  11      3801      1      123         1\n+  4    - 2999    + 2    +  49    - 9380\n----    ------    ---    -----    ------"
	if actual != expected {
		t.Errorf("Expected different output when calling ArithmeticArranger() with test input")
	}
}

func TestTooManyProblems(t *testing.T) {
	actual := ArithmeticArranger([]string{"44 + 815", "909 - 2", "45 + 43", "123 + 49", "888 + 40", "653 + 87"})
	expected := "Error: Too many problems."
	if actual != expected {
		t.Errorf("Expected calling ArithmeticArranger() with more than five problems to return 'Error: Too many problems.'")
	}
}

func TestIncorrectOperator(t *testing.T) {
	actual := ArithmeticArranger([]string{"3 / 855", "3801 - 2", "45 + 43", "123 + 49"})
	expected := "Error: Operator must be '+' or '-'."
	if actual != expected {
		t.Errorf("Expected calling ArithmeticArranger() with a problem that uses the '/' operator to return 'Error: Operator must be '+' or '-'.")
	}
}

func TestTooManyDigits(t *testing.T) {
	actual := ArithmeticArranger([]string{"24 + 85215", "3801 - 2", "45 + 43", "123 + 49"})
	expected := "Error: Numbers cannot be more than four digits."
	if actual != expected {
		t.Errorf("Expected calling ArithmeticArranger() with a problem that has a number over 4 digits long to return 'Error: Numbers cannot be more than four digits.'")
	}
}

func TestOnlyDigits(t *testing.T) {
	actual := ArithmeticArranger([]string{"98 + 3g5", "3801 - 2", "45 + 43", "123 + 49"})
	expected := "Error: Numbers must only contain digits."
	if actual != expected {
		t.Errorf("Expected calling ArithmeticArranger() with a problem that contains a letter character in the number to return 'Error: Numbers must only contain digits.'")
	}
}

func TestSolutions(t *testing.T) {
	actual := ArithmeticArranger([]string{"32 - 698", "1 - 3801", "45 + 43", "123 + 49"}, true)
	expected := "   32         1      45      123\n- 698    - 3801    + 43    +  49\n-----    ------    ----    -----\n -666     -3800      88      172"
	if actual != expected {
		t.Errorf("Expected solutions to be correctly displayed in output when calling ArithmeticArranger() with arithmetic problems and a second argument of `True`.")
	}
}

// ! DON'T CHANGE ANYTHING HERE OR THE RESULT WILL NOT DEFINED

// ! DON'T CHANGE ANYTHING HERE OR THE RESULT WILL NOT DEFINED
