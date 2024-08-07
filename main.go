package main

import "fmt"

func main() {
        var n int 

        fmt.Println("Enter the number: ")
        _, err := fmt.Scanln(&n)
        if err != nil{
                fmt.Println("Expected integer")
        }

        fibNumbers := fibonacci(n)

        for i, num := range fibNumbers{
                fmt.Printf("The fibonacci number %d: %d\n", i, num)
        }
}

func fibonacci(n int) []int {
        fib := make([]int, n+1)
        fib[0] = 0
        fib[1] = 1

        for i := 2; i <= n; i++ {
                fib[i] = fib[i-1] + fib[i-2]       
        }
        return fib
}

