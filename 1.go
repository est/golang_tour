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



type Vertex struct {
    X, Y int
}

var (
    p = Vertex{1, 2}  // has type Vertex
    q = &Vertex{1, 2} // has type *Vertex
    r = Vertex{X: 1}  // Y:0 is implicit
    s = Vertex{}      // X:0 and Y:0
)

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

    // structs nigga
    fmt.Println(p, q, q.X, r, s)

    // The expression new(T) allocates a zeroed T value and returns a pointer to it.
    var t1 *Vertex = new(Vertex)
    // alternative syntax
    t2 := new(Vertex)
    fmt.Println(t1, t2)

    // arrays
    var a [2]string
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a[0], a[1])
    fmt.Println(a)

    // array slice
    p := []int{2, 3, 5, 7, 11, 13}
    fmt.Println("p ==", p)

    for i := 0; i < len(p); i++ {
        fmt.Printf("p[%d] == %d\n", i, p[i])
    }
    // []T is a slice with elements of type T.


    p2 := []int{2, 3, 5, 7, 11, 13}
    fmt.Println("p ==", p2)
    fmt.Println("p[1:4] ==", p2[1:4])
    // missing low index implies 0
    fmt.Println("p[:3] ==", p2[:3])
    // missing high index implies len(s)
    fmt.Println("p[4:] ==", p2[4:])

    // the `make` function
    a2 := make([]int, 5)
    fmt.Println("a", a2)
    b2 := make([]int, 0, 5)
    fmt.Println("b", b2)
    c2 := b2[:2]
    fmt.Println("c", c2)
    d2 := c2[2:5]
    fmt.Println("d", d2)
    // @ToDo: what is cap() ?

    // nil slice
    var z3 []int
    fmt.Println(z3, len(z3), cap(z3))
    if z3 == nil {
        fmt.Println("nil!")
    }

    // internal :
    // http://golang.org/doc/articles/slices_usage_and_internals.html

    // test a rest
    // http://tour.golang.org/#36



}