package main

import (
	"fmt"
)

func main(){
	inchan := make(chan int)
	abort := make(chan int)
	//var b int = 5
	//var a int


	//go ab(outchan)
	go over(inchan, abort)


	//select {
	//	case a = <- inchan:
	//		fmt.Printf("Received a : %d \n", a)
	//	case outchan <- b :
	//		fmt.Printf("Sent b : %d \n", b)
	//	default:
	//		fmt.Println("Default")
	//}

	for {
		select {
		case i:= <- inchan:
			fmt.Println(i)
		case <- abort:
			fmt.Println("stopped")
			return
		}
	}

}

func over(ch chan int, ab chan int) {
	for i := 0; i<6 ; i++ {
		ch <- i
	}
	ab <- 1
}

func ab (i chan int){
	fmt.Println(i)
}

func cd (i chan int){
	i <- 4
}
