package main

// package pour l'exécution de print
import (
	"context"
	"fmt"
	"time"
)

// type "done" qui permettra de savoir si une goroutine est terminée ou non
// on définit alors un channel
type Done chan bool

/**

 */
func main() {
	// str := "Hello, World"
	// c := make(chan int)

	// done := make(Done)
	// ctx := context.Background()
	// var delay time.Duration = 500 * time.Millisecond
	// go printThis(ctx, str, done, delay)

	// i := 0
	// // fonction faite en parallèle dans un autre thread
	// go func() {
	// 	print(i)
	// 	i++
	// 	// envoi dans le thread principal (dans le channel) la variable i
	// 	c <- i
	// 	done <- false
	// }()
	// time.Sleep(delay)

	// fmt.Println(<-c)

	// if <-done {
	// 	fmt.Println("printed!")
	// } else {
	// 	fmt.Println("not printed")
	// }

	// myNewString := exampleFunctionReturn("printed")
	// fmt.Println(myNewString)

}

func printThis(ctx context.Context, stringToPrint string, done Done, delay time.Duration) {
	select {
	case <-time.After(delay):
		fmt.Println(stringToPrint)
	}
}

func exampleFunctionReturn(firstString string) string {
	return firstString + "xxxx"
}
