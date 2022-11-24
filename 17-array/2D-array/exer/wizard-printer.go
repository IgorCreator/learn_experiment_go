package main

import "fmt"

func main() {
	scientists := [...][3]string{
		{"Albert", "Einstein", "emc2"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
	}

	fmt.Println("First Name      Last Name       Nickname")
	fmt.Println("==================================================")
	for i := range scientists {
		fmt.Printf("%-15s %-15s %-15s\n", scientists[i][0], scientists[i][1], scientists[i][2])
	}
}
