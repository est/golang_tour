package main 
// without double quotes for namespace?

import "math"
import "fmt"
import (
    "math/rand"
    "math/cmplx"
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

var (
    ToBe   bool       = false
    MaxInt uint64     = 1<<64 - 1
    z      complex128 = cmplx.Sqrt(-5 + 12i) //wtf cmplx
)



func Sqrt(x float64) float64 {
    z := 1.0
    last := 0.0
    d := 0.0000000000000000000000000000001
    for z-last>d{
        last = z
        z = z - (z*z-x)/x/2
    }
    return z
}

func main() {
    fmt.Println("My favorite number is", rand.Intn(10))
    fmt.Printf("Now you have %g problems.",
        math.Nextafter(2, 3))
    fmt.Printf("Math of Pi is ", math.Pi) // Captical for exported
    fmt.Println(split(17)) // named results

    fmt.Println(i, c, python, java)

    //Short variable declarations 
    var i, j int = 1, 2
    k := 3
    c, python, java := true, false, "no!"
    fmt.Println(i, j, k, c, python, java)

    // Go's basic data types are
    // bool
    // string
    // int  int8  int16  int32  int64
    // uint uint8 uint16 uint32 uint64 uintptr
    // byte // alias for uint8
    // rune // alias for int32 or a Unicode code point
    // float32 float64
    // complex64 complex128


    // Constants cannot be declared using the := syntax.
    const f = "%T(%v)\n"
    fmt.Printf(f, ToBe, ToBe)
    fmt.Printf(f, MaxInt, MaxInt)
    fmt.Printf(f, z, z)


    // Go has only one looping construct, the for loop.
    // no (), requires {}
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)


    // if and for exercise
    // http://tour.golang.org/#25
    fmt.Println(Sqrt(2), math.Sqrt(2))


}