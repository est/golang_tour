package main 
// without double quotes for namespace?

import "math"
import "fmt"
import (
    "math/rand"
)


// consective parameter x, y int
func add(x int, y int) int {
    return x + y
}


// named results 
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

// test vars
var i int
var c, python, java bool



func main() {
    fmt.Println("My favorite number is", rand.Intn(10))
    fmt.Printf("Now you have %g problems.",
        math.Nextafter(2, 3))
    fmt.Printf("Math of Pi is ", math.Pi) // Captical for exported
    fmt.Println(split(17)[0]) // named results

    fmt.Println(i, c, python, java)

    

}