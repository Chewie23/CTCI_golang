/*
Prompt:
Assumeyou have a method isSubstringwhich checks if one word is a substring of
another. Given two strings, sl and s2, write code to check if s2 is a rotation
of sl using only one call to isSubstring (e.g.,"waterbottle" is a rotation of"erbottlewat").
*/

//OOF. I need sometime to parse this. Because diving into coding is a bad plan

/*
I had sometime to parse this. There are all kinds of CORRECT edge cases involved with this
The base case is:
 String_1 = original unaltered word
 String_2 = rotated original word

Ex.
str_1 = "odyssey"
str_2 = "sseyody"

The edge case is having rotations of the original word
Ex.
str_1 = "sseyody"
str_2 = "eyodyss"


Then there are the WRONG cases of different strings
1) same length, different letters
2) different length, same letters
3) different length, different letters

The hints (which I should not have looked at before trying) suggest concantenating
the str_2 to find out where str_1 is.

So. There are two paths here:
1) Assume that the correct edge case of two different rotations of the same string is
not in the picture
	Easy to do, since str_1 would be the unaltered string

2) Involve the correct edge case. Which is gonna suck. Since yeah.


*/

package sr

import "fmt"

func Example() {
	fmt.Println("vim-go")
}
