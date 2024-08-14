package main

import "fmt"

func main() {
        var n int 

        fmt.Println("Enter the number: ")
        _, err := fmt.Scanln(&n)
        if err != nil{
                fmt.Println("Expected integer")
                return 
        }

        if n < 0 {
                fmt.Println("Number must be non negative")
                return
        }

        fibNumbers := fibonacci(n)

        for i, num := range fibNumbers{
                fmt.Printf("The fibonacci number %d: %d\n", i, num)
        }
}

func fibonacci(n int) []int {
        if n == 0{
                return []int{0}
        } else if n == 1{
                return []int{0,1}
        }
        fib := make([]int, n+1)
        fib[0] = 0
        fib[1] = 1

        for i := 2; i <= n; i++ {
                fib[i] = fib[i-1] + fib[i-2]       
        }
        return fib
}
        fib[1] = 1

        for i := 2; i <= n; i++ {
                fib[i] = fib[i-1] + fib[i-2]       
        }
        return fib
}

