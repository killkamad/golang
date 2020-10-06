package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func task1() {
	var number float64
	fmt.Print("Введите float: ")
	fmt.Scan(&number)
	var output = int64(number)
	fmt.Println(output)
}

func task2() {
	fmt.Print("Введите строку: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	text = strings.ToLower(text)
	start := strings.HasPrefix(text, "b")
	end := strings.HasSuffix(text, "c")
	if start == true && end == true {
		between := strings.Contains(text, "z")
		if between == true {
			fmt.Println("Match")
		} else {
			fmt.Println("Try again")
		}
	} else {
		fmt.Println("Try again")
	}

}

func main() {

	for true {
	   task1()
	   task2()
	}

	//check := [][]int{{15, 11}, {22, 55}, {33, 5}, {11, 4}}
	////fmt.Println(check[0])
	//for i := range check {
	//	fmt.Println(check[i][0], []int{15, 11}[0])
	//}
	//aa := 1
	//bb := 2
	//output := fmt.Sprintf("%d%d", aa, bb)
	//fmt.Println("dadada ", output)
}

//fmt.Println("hello world")
//var x int = 5
//fmt.Println(x)
//var y float64 = 22.1
//var a int32 = y
//fmt.Println(y)
//
//test:= "This is string"
//fmt.Println(test)
//
//var six int16 = 7
//var two int32 = 2
//
//two = int32(six)
//fmt.Println(two)
//fmt.Println(&two)
