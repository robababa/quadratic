package quadratic

import (
	"math"
	"sort"
)

type Solution interface {
	Always() bool
	Never() bool
	Sometimes() bool
	Values() []int
}

type Never struct {}

func (_ Never) Always() bool {
	return false
}

func (_ Never) Never() bool {
	return true
}

func (_ Never) Sometimes() bool {
	return false
}

func (_ Never) Values() []int {
	var answer []int
	return answer
}

type Always struct {}

func (_ Always) Always() bool {
	return true
}

func (_ Always) Never() bool {
	return false
}

func (_ Always) Sometimes() bool {
	return false
}

// return an empty array. The caller needs to check Always() first
func (_ Always) Values() []int {
	var answer []int
	return answer
}

type Sometimes struct {
	solutionValues []int
}

func (_ Sometimes) Always() bool {
	return false
}

func (_ Sometimes) Never() bool {
	return false
}

func (_ Sometimes) Sometimes() bool {
	return true
}

func (s Sometimes) Values() []int {
	return s.solutionValues
}

func CombineSolutions(solutions ...Solution) Solution {
	// if we have no solutions, return Never
	if len(solutions) == 0 {return Never{}}

	// check for Never among the solutions
	for _, sol := range solutions {
		if sol.Never() {return Never{}}
	}
	// check for Always
	alwaysForAll := true
	for _, sol := range solutions {
		if !sol.Always() {
			alwaysForAll = false
			break
		}
	}
	if alwaysForAll {return Always{}}

	// we have work to do. Identify the candidates, sort them, and see if one works
	var candidates []int
	for _, sol := range solutions {
		if sol.Sometimes() {
			candidates = append(candidates, sol.Values()...)
			}
	}
	sort.Ints(candidates)

CandidateLoop:
	for _, c := range candidates {
		SolutionsLoop:
		for _, sol := range solutions {
			// if this Solution is valid for all positive integers, then great, go to the next Solution!
			if sol.Always() {continue SolutionsLoop}
			for _, val := range sol.Values() {
				// if the candidate is a value in the Solution, then great, go to the next Solution!
				if val == c {continue SolutionsLoop}
			}
			// didn't find the candidate in this Solution, go to the next candidate
			continue CandidateLoop
		}
		// this candidate worked for all solutions, so it is the lowest positive integer that works.  Return it!
		return Sometimes{solutionValues: []int{c}}
	}
	// no candidates worked
	return Never{}
}

// returns the positive integer Solution to the equation bx + c = 0, if one exists
// if it is Always true for all positive x, then it returns AlwaysZero
// if it is true for no positive integers x, then it returns NoSolutionValue
func LinearPositiveIntegerSolution(b int, c int) Solution {
	switch {
	// 0 = 0 is Always true
	case b == 0 && c == 0: return Always{}
		// c == 0 is Always false when c != 0
	case b == 0 && c != 0: return Never{}
		// bx = 0 for b != 0 is true iff x is 0
	case b != 0 && c == 0: return Sometimes{solutionValues: []int{0}}
	default: {
		candidate := int(-1 * c / b)
		if candidate > 0 && b * candidate + c == 0 {
			return Sometimes{solutionValues: []int{candidate}}
		} else {
			return Never{}
		}
	}
	}
}

// returns the positive integer solutions to the quadratic equation
// ax^2 + bx + c = 0
// if there are two solutions, this function returns both of them
// if there is only one positive integer Solution, it returns that Solution and NoSolutionValue
// if there are no positive integer solutions, the function returns NoSolutionValue, NoSolutionValue
func QuadraticPositiveIntegerSolutions(a int, b int, c int) Solution {
	// first, get some edge cases out of the way when a = 0
	if a == 0 {
		return LinearPositiveIntegerSolution(b, c)
	}

	// at this point, we have a real quadratic equation
	discriminant := b * b - 4 * a * c
	// imaginary solutions don't work here
	if discriminant < 0 {
		return Never{}
	}

	// the discriminant is nonnegative.  However, if its root is not an integer, we have no solutions
	discriminantRoot := int(math.Sqrt(float64(discriminant)))
	if discriminantRoot * discriminantRoot != discriminant {
		return Never{}
	}

	// put our potential solutions in ascending order
	couldBe := [2]int{
		((-1 * b) - discriminantRoot) / (2 * a),
			((-1 * b) + discriminantRoot) / (2 * a),
			}
	var reallyWorks []int

	// if the possible solutions don't actually work, set them to NoSolutionValue
	if couldBe[0] >= 0 && a * couldBe[0] * couldBe[0] + b * couldBe[0] + c == 0 {
		reallyWorks = append(reallyWorks, couldBe[0])
	}

	if couldBe[1] >= 0 && couldBe[1] != couldBe[0] && a * couldBe[1] * couldBe[1] + b * couldBe[1] + c == 0 {
		reallyWorks = append(reallyWorks, couldBe[1])
	}

	if len(reallyWorks) == 0 {
		return Never{}
	}

	// reallyWorks has one or two elements, so there are some Solution Values
	return Sometimes{reallyWorks}
}

