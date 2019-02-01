package zm

/*
Prompt
 Write an algorithm such that if an element in an MxN matrix is 0, its entire
 row and column are set to 0.
*/

//Ex. ("simple" case of one 0)
/*

	[1 2 3] 	   [1 0 3]
	[4 5 6]		   [4 0 6]
	[7 0 9]	->	   [0 0 0]
	[8 7 1]	       [8 0 1]
*/

import "fmt"

func isZero(num int) bool {
	return num == 0
}
func PrintMatrix(matrix [][]int) {
	for _, x := range matrix {
		fmt.Println(x)
	}
}

func zeroOut(cur_row, cur_col, max_row_len, max_col_len int, matr [][]int) [][]int {
	//Changing the current column to 0
	for r := 0; r < max_row_len; r++ {
		matr[r][cur_col] = 0
	}
	//Changing the current row to 0
	for c := 0; c < max_col_len; c++ {
		matr[cur_row][c] = 0
	}
	return matr
}

//TODO
//Handle the multiple zero case
//	WITHOUT it compounding upon itself

//Idea 1: catalog all the zero coordinates and then use my algorithm of clearing out crap

//This is most definitely brute force, but not sure how to make it more elegant
func iterateMatrix(matrix [][]int) bool {
	var max_row_len = len(matrix)
	var max_col_len = len(matrix[0])

	//Gets row
	//Have to use C-style for loop, since will complain if index was not used
	//Will cover the case where there is NO zeros
	for r := 0; r < max_row_len; r++ {
		//Gets column
		for c := 0; c < max_col_len; c++ {
			if isZero(matrix[r][c]) {
				PrintMatrix(zeroOut(r, c, max_row_len, max_col_len, matrix))
				return true
			} // if
		} // nested for

	} // outer for
	return false
} // func

func Example() {
	var matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 0, 9},
		{8, 7, 1}}

	//Brute force algo: iterate through every element. If spotted a zero, then need
	//to clear entire row AND the column

	if !(iterateMatrix(matrix)) {
		fmt.Println("No zeros in matrix!")
	}

}

/*
POST MORTEM
I cracked and took a peak at a Python solution (the official answer):
https://github.com/careercup/CtCI-6th-Edition-Python/blob/e6bc732588601d0a98e5b1bc44d83644b910978d/Chapter1/8_Zero%20Matrix/ZeroMatrix.py

My god, it's full of stars. But seriously, it is beautifully simple. It only concerns
itself with what it wants to change and then changes it.

Basically tracking all the rows that have zeros (which is a bit redundant with multiple
zeros, but covers the edge cases) and all
the columns that have zeros. Then iterating through all of the
rows and cols.

I sort of touched upon it when I coded this guy up, with the zeroOut() method, since
you only need to hold the row or col constant and iterate through the other

Take away:
Need to use a data structure or something to store relevant data. But it's mostly
pattern recognition to see that we hold the row/col constant and iterate through
the other half

I think it boils down to asking: what is changing, what is remaining the same, and
how to store/mark the things changing. General advice, I know, but it does help.

I need to ruminate on this for a bit. But for now? Keep on keeping on
Just need to keep making freethrows until they become second nature

*/
