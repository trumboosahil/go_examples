package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var dbdata = []string{"db1", "db2", "db3", "db4", "db5"}
var wg = sync.WaitGroup{}

func main() {
	//printMe()
	testConcurrency()
}

func testConcurrency() {
	t0 := time.Now()
	for i := 0; i < len(dbdata); i++ {
		wg.Add(1)
		go dbcall(i)
	}
	wg.Wait()
	fmt.Printf("\nTime passes id %v", time.Since(t0))
}

func dbcall(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Microsecond)
	fmt.Println(i)
	fmt.Printf("\nthe db call return is %v index at %v ", dbdata[i], i)
	wg.Done()
}

func printMe() {
	var result, rem, err = numerator(6, 4)
	if err != nil {
		fmt.Println(err.Error())
	} else if rem == 0 {
		fmt.Printf("The result of integer division is %v", result)
	} else {
		fmt.Printf("The result of non-integer dividion is %v with reminder %v", result, rem)
	}
}

func PrintString() {
	var mtstr string = "resسume"
	for i, key := range mtstr {
		fmt.Printf("%v %v", i, key)
		fmt.Println()
	}
}
func PrintStringRunes() {
	var mtstr = []rune("resسume")
	for i, key := range mtstr {
		fmt.Printf("%v %v", i, key)
		fmt.Println()
	}
}

func numerator(num1 int, num2 int) (int, int, error) {
	var err error
	if num2 == 0 {
		err = errors.New("denominator is zero")
		return num1, num2, err
	}
	var result int = num1 / num2
	var rem int = num1 % num2
	return result, rem, err
}
func keyvalue() {
	var mymap map[string]uint8 = make(map[string]uint8)
	mymap["hello"] = 1
	mymap["hello1"] = 100
	fmt.Println(mymap["hello"])
	for key, value := range mymap {
		fmt.Printf("key : %v\n", key)
		fmt.Printf("Value : %v\n", value)

	}

}
func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	var i int = 0
	for ; i < n; i++ {
		slice = append(slice, 1)
	}
	print(i)
	return time.Since(t0)
}
