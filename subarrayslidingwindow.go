// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	plots := []int{3, 1, 4, 5, 1, 2, 3, 3}

	k := 8
	i := 0
	j := 0
	cs := 0

	for j < len(plots) {

		//take current element
		//go right and expand the window
		cs += plots[j]
		j++

		//exclude from left
		for cs > k && i < j {
			cs = cs - plots[i]
			i++
		}

		if cs == k {
			fmt.Println(i, j-1)
		}

	}

}
