package main

import "fmt"

//MAIN
func main() {
    
    //#######################################
    //VAR 
    rect1 := Rectangle{leftX:0, topY: 50, height: 10, width: 10}
    rect2 := Rectangle{0, 60, 20, 20}
    
    //#######################################
    //FUNC
    func (rect *Rectangle) area() float64 {
        return rect.width * rect.height
    }

    //#######################################
    //PRINT
    fmt.Println("RECTANGLE is ", rect1.width, "wide")
    fmt.Println("AREA is ", rect1.area())
  
}

//#######################################
//STRUCT
type Rectangle struct {
    
    //ATTRIBUTES
    leftX float64
    topY float64
    height float64
    width float64
 
}
