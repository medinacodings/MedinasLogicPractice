package main

import (
	"MedinasLogicPractice/Logic3"
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	for i := 1; i <= 5; i++ {
		n := 9 // Desired size of the triangle

		// Call the Number3 Function from Logic3 package
		result := Logic3.Number14(n)

		// Print the result
		for i := 0; i < len(result); i++ {
			fmt.Println(result[i])
		}
	}
}
