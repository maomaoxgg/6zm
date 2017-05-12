# 6zm
normal arithmetic/一些算法
package main

import (
	"sync"
	//"fmt"
	//"time"
	"fmt"
	"math/rand"
	"time"
)

var m1,m2,c int

var lock sync.Mutex
var ch,all chan int

type SAS struct {
	A string
}

func main() {

	var (
		arr3pokers [][]int
		arrPoker   []int
		a, b, c    int
	)
	for i := 0; i < 52; i++ {
		arrPoker = append(arrPoker, i+1)
	}
	fmt.Println(arrPoker)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for count := 0; count < 30; count++ {
		a = r.Intn(52)
		b = r.Intn(52)
		c = arrPoker[a]
		arrPoker[a] = arrPoker[b]
		arrPoker[b] = c
	}
	fmt.Println(arrPoker)
	arrPoker = arrPoker[:15]
	fmt.Println(arrPoker)
	for i := 0; i < 3; i++ {
		arr3pokers = append(arr3pokers, arrPoker[i*5:(i+1)*5])
	}
	fmt.Println(arr3pokers)
	for i:=0;i<len(arr3pokers);i++{
		idChangeNumber(arr3pokers[i])
		fmt.Println(arr3pokers[i])
	}
}

func idChangeNumber(a []int) []int{
	for i:=0;i<len(a);i++{
		num := a[i]%13
		switch num {
		case 0 :
			a[i] = 13
		default:
			a[i] = num
		}

	}
	return a
}

	//var stat int
	//c = 100000000
	//m1 = 0;m2 = 0
	//all = make(chan int,2)
	//timeout := make (chan bool, 1)
	//go func() {
	//	time.Sleep(5e9) //
	//	timeout <- true
	//}()
	//fmt.Printf("%s all started!\n",time.Now())
	//go func(count int){
	//	for m1 < count {
	//		m1 = testMutex(m1)
	//	}
	//	fmt.Printf("%s mutex finished!\n",time.Now())
	//	all <- 1
	//}(c)
	//go func(count int){
	//	for m2 < count {
	//		m2 = testChannel(m2)
	//	}
	//	fmt.Printf("%s channel finished!\n",time.Now())
	//	all <- 2
	//}(c)
	//select {
	//case stat = <- all:
	//	fmt.Printf("%d faster!\n",stat)
	//case <- timeout:
	//	fmt.Println("超时")
	//}
	//fmt.Printf("%s all finished!\n",time.Now())


func testMutex(m int) int{
	lock.Lock()
	m = m+1
	lock.Unlock()
	return  m
}

func testChannel(m int) int{
	ch = make(chan int,1)
	ch <- 1
	m = m+1
	<- ch
	return  m
}

