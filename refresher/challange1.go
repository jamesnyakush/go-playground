package main

import "fmt"

func main() {

	//coding := []string{"JavaScript", "Ruby", "python", "java", "Kotlin", "Pascal", "C#", "C++"}
	//coding = append(coding,"Coders")

	s := make([]string, 1, 4)
	s[0] = "filename"
	sFinal := []string{"test", "test1"}
	depString := append(s[:3], sFinal...)
	fmt.Println("dep_String is ", depString)

	//fmt.Println("appending", coding)
}
