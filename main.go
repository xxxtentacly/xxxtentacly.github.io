//Write a program that takes a string as input and checks whether it is a palindrome.

package main

import (
    "fmt"
    "strings"
)

func main() {
    var input string
    fmt.Println("Enter a string:")
    fmt.Scanln(&input)

   
    input = strings.ReplaceAll(input, " ", "")

    // проверка на полиндром
    isPalindrome := true
    for i := 0; i < len(input)/2; i++ {
        if input[i] != input[len(input)-i-1] {
            isPalindrome = false
            break
        }
    }

    //результат
    if isPalindrome {
        fmt.Println("The input string is a palindrome.")
    } else {
        fmt.Println("The input string is not a palindrome.")
    }
}
