package main

import "fmt"

type axis struct{
	cardinalDirection string
}

type cD struct{
	cardinalDirection string
}
type c struct{
	colour string
}


func main() {
	aList := []cD{cD {"north"},cD {"east"},cD {"south"},cD {"west"}}
	fmt.Println(aList)

}