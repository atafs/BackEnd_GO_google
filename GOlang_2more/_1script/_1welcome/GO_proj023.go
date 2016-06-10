package main

//#######################################
//IMPORT
import (
    "fmt"
)

//#######################################
//TYPE
type Office int

//#######################################
//CONSTANT
const (
	Boston Office = iota
	NewYork
)

//#######################################
//FUNC
func (o *Office) String() string {
	return "Google, " + officePlace[0]
}

//#######################################
//VAR global
var officePlace = [5] string {"Boston","NewYork","Agrela","Ana","Batalha"}    

//#######################################
//MAIN
func main() {
 	
    //***********************************
    //PRINT	
	fmt.Printf("Hello World Americo Tomas - %s\n", officePlace[0])   

	fmt.Printf("Hello World Americo Tomas - %s\n", Boston)    
	
}

