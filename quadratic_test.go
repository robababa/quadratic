package quadratic

import "fmt"

func linearHelper(b, c int) {
	sol := LinearPositiveIntegerSolution(b, c)
	fmt.Println(b, ",", c, "Always(), Never(), Values() are", sol.Always(), sol.Never(), sol.Values())
}

func ExampleLinearPositiveIntegerSolution() {
	// 2x - 6 = 0 : x = 3 is a Solution
	linearHelper(2, -6)

	// 2x - 5 = 0 : x = 2.5 is not an integer, so there are no solutions
	linearHelper(2, -5)

	// 2x + 6 = 0 : x = -3 is not positive, so there are no solutions
	linearHelper(2, 6)

	// 2x = 0 : x = 0 is a Solution
	linearHelper(2, 0)

	// 0x = 0 : Always is the Solution
	linearHelper(0, 0)

	// Output:
	// 2 , -6 Always(), Never(), Values() are false false [3]
	// 2 , -5 Always(), Never(), Values() are false true []
	// 2 , 6 Always(), Never(), Values() are false true []
	// 2 , 0 Always(), Never(), Values() are false false [0]
	// 0 , 0 Always(), Never(), Values() are true false []
}

func quadraticHelper(a, b, c int) {
	sol := QuadraticPositiveIntegerSolutions(a, b, c)
	fmt.Println(a, ",", b, ",", c, "Always(), Never(), Values() are", sol.Always(), sol.Never(), sol.Values())
}

func ExampleQuadraticPositiveIntegerSolutions() {
	// 1*x^2 - 4x + 3 = 0 : x = 1, 3
	quadraticHelper(1, -4, 3)
	// x^2 - 4 = 0 : x = 2
	quadraticHelper(1, 0, -4)
	// x^2 - 4x + 4 = 0 : x = 2 is a double Solution, but should only appear once
	quadraticHelper(1, -4, 4)
	// x^2 + 4x + 3 = 0 : no solutions
	quadraticHelper(1, 4, 3)

	fmt.Println()
	fmt.Println("a = 0:")
	// 0x^2 + x - 7 = 0 : x = 7
	quadraticHelper(0, 1, -7)

	fmt.Println()
	fmt.Println("a = b = 0, c nonzero")
	// 2 = 0 : no solutions
	quadraticHelper(0, 0, 2)

	fmt.Println()
	fmt.Println("a = b = c = 0")
	// 0 = 0 : Always
	quadraticHelper(0, 0, 0)
	// Output:
	// 1 , -4 , 3 Always(), Never(), Values() are false false [1 3]
	// 1 , 0 , -4 Always(), Never(), Values() are false false [2]
	// 1 , -4 , 4 Always(), Never(), Values() are false false [2]
	// 1 , 4 , 3 Always(), Never(), Values() are false true []
	//
	// a = 0:
	// 0 , 1 , -7 Always(), Never(), Values() are false false [7]
	//
	// a = b = 0, c nonzero
	// 0 , 0 , 2 Always(), Never(), Values() are false true []
	//
	// a = b = c = 0
	// 0 , 0 , 0 Always(), Never(), Values() are true false []
}

func ExampleCombineSolutions() {
	fmt.Println("CombineSolutions(Always{}, Always{}).Always() =",
		CombineSolutions(Always{}, Always{}).Always())
	fmt.Println("CombineSolutions(Always{}, Never{}).Never() =",
		CombineSolutions(Always{}, Never{}).Never())
	fmt.Println("CombineSolutions(Always{}, Sometimes{solutionValues: []int{5, 4}}).Values() =",
		CombineSolutions(Always{}, Sometimes{solutionValues: []int{5, 4}}).Values())
	fmt.Println("CombineSolutions(Never{}, Sometimes{solutionValues: []int{5, 4}}).Never() =",
		CombineSolutions(Never{}, Sometimes{solutionValues: []int{5, 4}}).Never())
	fmt.Println("CombineSolutions(Sometimes{solutionValues: []int{3, 9, 7}}, Sometimes{solutionValues: []int{5, 7, 4}}).Values() =",
		CombineSolutions(Sometimes{solutionValues: []int{3, 9, 7}}, Sometimes{solutionValues: []int{5, 7, 4}}).Values())
	// Output:
	// CombineSolutions(Always{}, Always{}).Always() = true
	// CombineSolutions(Always{}, Never{}).Never() = true
	// CombineSolutions(Always{}, Sometimes{solutionValues: []int{5, 4}}).Values() = [4]
	// CombineSolutions(Never{}, Sometimes{solutionValues: []int{5, 4}}).Never() = true
	// CombineSolutions(Sometimes{solutionValues: []int{3, 9, 7}}, Sometimes{solutionValues: []int{5, 7, 4}}).Values() = [7]
}