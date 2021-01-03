package main

import (
	//"sync"
	"fmt"
	"sort"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sync"
)


//Please enter an array of integers seperated by , :
//3,4,34,43,65,23,65,73,2,64,74,
//[2 64 74]
//[23 65 73]
//[34 43 65]
//[3 4]
//
//Sorted out Array : [2 3 4 23 34 43 64 65 65 73 74]

func main(){


	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter an array of integers seperated by , :")

	text, _ := reader.ReadString('\n')

	userInputStr := strings.Split(text, ",")
	var userNumInR []int
	var num int
	for _,str := range userInputStr {
		num, _ = strconv.Atoi(str)
		userNumInR = append(userNumInR,num)
	}
	userNumIn := userNumInR[0:len(userNumInR)]

	lenUserNum := len(userNumIn) - 1
	var (
		arr0 []int
		arr1 []int
		arr2 []int
		arr3 []int
	)

	arr0 = userNumIn[0:(lenUserNum/4)]
	arr1 = userNumIn[(lenUserNum/4):(lenUserNum/2)]
	arr2 = userNumIn[(lenUserNum/2):(3*lenUserNum/4)]
	arr3 = userNumIn[(3*lenUserNum/4):(lenUserNum)]

	var syn sync.WaitGroup
	//ch := make(chan []int, 10)
	syn.Add(4)
	go SortNum(arr0, &syn)
	go SortNum(arr1, &syn)
	go SortNum(arr2, &syn)
	go SortNum(arr3, &syn)


	syn.Wait()

	sort.Ints(userNumIn)
	fmt.Printf("\nSorted out Array : %v", userNumIn[1:])

}

func SortNum( sli []int, syn * sync.WaitGroup)  {
	sort.Ints(sli)
	fmt.Println(sli)
	syn.Done()
}

