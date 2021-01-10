package main

import "fmt"

import "rsc.io/quote"

func main() {
    fmt.Println("")
    fmt.Println("-- Call Hello, returns a greeting.")
    fmt.Println(quote.Hello())
    fmt.Println("")
    fmt.Println("-- Call Go, returns a Go proverb.")
    fmt.Println(quote.Go())
    fmt.Println("")
    fmt.Println("-- Call Glass, returns a useful phrase for world travelers.")
    fmt.Println(quote.Glass())
    fmt.Println("")
    fmt.Println("-- Call Opt, returns an optimization truth.")
    fmt.Println(quote.Opt())
}

