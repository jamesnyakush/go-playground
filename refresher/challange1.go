package main

import "fmt"

func main() {
	
	//s := make([]string, 3)
	// i wanna append  Coders in the middle of this array
	coding := []string{"JavaScript", "Ruby", "python", "java", "Kotlin", "Pascal", "C#", "C++"}
	coding = append(coding, "Coders")

	fmt.Println("appending", coding)
}
