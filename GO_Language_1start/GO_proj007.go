package main

import "fmt"

//MAIN
func main() {
    
    //#######################################
    //VAR 
    numSlice := []int {5,4,3,2,1}
    numSlice2 := numSlice[3:5]
    
    //VAR ARRAY WITH NO VALUES 
    numSlice3 := make([]int, 5, 10)

    //#######################################
    //PRINT
    fmt.Println("numSlice2[2] = ",numSlice2[2])
    fmt.Println("numSlice[:2] = ",numSlice[:2])
    fmt.Println("numSlice[2:] = ",numSlice[2:])
    
    //COPY STRING
    copy(numSlice3, numSlice)
    fmt.Println("numSlice3[0] = ", numSlice3[0])
    
    //APPEND STRING
    numSlice3 = append(numSlice3, 0, -1)
    fmt.Println("numSlice3[6] = ", numSlice3[6])
}