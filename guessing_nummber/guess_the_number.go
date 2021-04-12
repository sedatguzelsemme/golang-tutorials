package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "time"
    "os"
    "strings"
    "strconv"
)


func main(){
    rand.Seed(time.Now().UnixNano())
    min := 1
    max := 100
    tryCount := 0
    reader := bufio.NewReader(os.Stdin)

    fmt.Println("Welcome to the Number Guessing Game!")
    fmt.Println("I'm thinking of a number between 1 and 100.")
    fmt.Println("Choose a difficulty. Type 'easy' or 'hard': ")
    level, _ := reader.ReadString('\n')
    level = strings.Replace(level, "\n", "", -1)

    if strings.Compare("easy", level) == 0{
        tryCount = 5
    }else if strings.Compare("hard", level) == 0{
        tryCount = 10
    }


    //fmt.Println("You have ", tryCount, "attempts remaining to guess the number")
    //userGuessedNumber, _ := reader.ReadString('\n')
    //userGuessedNumber = strings.Replace(userGuessedNumber, "\n", "", -1)

    keepPlaying := true
    for keepPlaying{
        randomNumber := rand.Intn(min +  max)
        //randomIntNumber, err := strconv.Atoi(randomNumber)
        notFound := true
        for tryCount > 0 && notFound {

            fmt.Println("You have ", tryCount, "attempts remaining to guess the number")
            userGuessedNumber, _ := reader.ReadString('\n')
            userGuessedNumber = strings.Replace(userGuessedNumber, "\n", "", -1)
            userGuessedInt, _ := strconv.Atoi(userGuessedNumber)

            if randomNumber == userGuessedInt{
                notFound = false
            }else if randomNumber < userGuessedInt{
                fmt.Println("Too high")
                tryCount -= 1
            }else if randomNumber > userGuessedInt{
                fmt.Println("Too low")
                tryCount -= 1
            }
            if tryCount <= 0{
                notFound = false
                fmt.Println("You lose")
            }
        }
        fmt.Println("Play again yes/no")
        playAgain, _ := reader.ReadString('\n')
        playAgain = strings.Replace(playAgain, "\n", "", -1)
        if strings.Compare("yes", playAgain) == 0{
            keepPlaying = true
            notFound = true
            tryCount = 5
        }else if strings.Compare("no", playAgain) == 0{
            keepPlaying = false
        }
        //keepPlaying = false
    }

}
