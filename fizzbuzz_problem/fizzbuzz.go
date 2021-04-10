package main


import "fmt"

/*
FizzBuzz Problem
    In this problem, you should display a list from 1 to 100, one on each line, with the following expectations:
    Numbers divisible by 3 should appear as 'fizz' instead of number;
    Numbers divisible by 5 should appear as 'buzz' instead of number;
    Numbers divisible by 3 and 5 should appear as' fizzbuzz 'instead of number'.
*/


func main(){

    for i:=1; i<101; i++{
        if i%5 == 0 && i%3==0{
            fmt.Println("fizzbuzz")
        }
        if i%5 == 0{
            fmt.Println("buzz")
        }
         if i%3 == 0{
            fmt.Println("fizz")
        }

    }
}


