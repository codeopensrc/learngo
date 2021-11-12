package main

import (
    "fmt"
    "log"

    "rsc.io/quote"
    "example/greetings"
)


//Const can be character, string, bool, and num types
//Consts cannot use := (even in function scope)
const MyConstString = "important var"

// Declare over multiple lines
const (
    MyConstBool = true
    MyConstNum = 30
)

func main() {
    fmt.Println("Hello, World!")
    fmt.Println(quote.Go())

    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    names := []string{"Gladys", "Samantha", "Darrin"}

    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(messages)
}

// Return value types in function header
func more() (int, bool) {

    //Can use var
    var mynum = 3
    fmt.Println(mynum)

    //Or use := for shorthand
    //Can not be used outside a function
    mynewnum := 2
    fmt.Println(mynewnum)

    //Can delare types
    var istrue bool
    var num int
    fmt.Println(istrue, num)

    //Multi assignment
    var newnum, secondnum = 2, 3
    fmt.Println(newnum, secondnum)

    return num, istrue
}

//functions support named returned values known as naked returns
func example(str1 string, int1 int) (strreturn string, intreturn int) {
    strreturn = str1 + " more str"
    intreturn = int1 + 3
    return
}


//https://tour.golang.org/basics/11
func basictypes() {
    //bool
    boolstatement := true

    //string
    str := "random str"

    //int types
    //An int can store at maximum a 64-bit integer
    //int int8 int16 int32 int64
    //uint uint8 uint16 uint32 uint64 uintptr
    myint := 2
    fmt.Println(boolstatement, str, myint)

    //byte - alias for uint8
    //rune - alias for int32

    //float32 float64
    //complex64 complex128


    //Printing type with %T and value with %v
    fmt.Printf("Type: %T Value: %v\n", myint, myint)

    //zero values for types without initial values
    //0 for nums
    //false for bool
    //"" for strings
}


func typeconversion() {
    var myint = 32

    //to float
    f := float64(myint)

    //to uint
    u := uint(f)

    fmt.Println(myint, f, u)
}
