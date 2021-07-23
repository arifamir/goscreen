GoScreen package provides an easy way for clearing the screen and getting 
the size of the current terminal.


## Installation:

`$ go get -u github.com/arifamir/goscreen`

## Clearing the Screen:

You can clear the screen and move the cursor to the top-left corner of the screen. This is good enough to create an animated program (such as an always updating clock or a progress bar).

    package main

    import (
    "fmt"
    "time"
    "github.com/arifamir/goscreen"
    )

    func main() {

        goscreen.Clear()
    
        for {
            // Moves the cursor to the top-left position of the screen
            goscreen.MoveTopLeft()
    
            // Animate the time always in the same position
            fmt.Println(time.Now())
    
            time.Sleep(time.Second)
        }
    }
