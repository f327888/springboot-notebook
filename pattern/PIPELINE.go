package main

import (
	"fmt"
	"math"
	"sync"
)

// 需一个 echo()函数，其会把一个整型数组放到一个Channel中，并返回这个Channel
func echo(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// 平方函数
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n // 队列里面的元素
		}
		close(out)
	}()
	return out
}

//过滤奇数函数
func odd(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if n%2 != 0 {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

// 求和函数
func sum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum = 0
		for n := range in {
			sum += n
		}
		out <- sum // 求和
		close(out)
	}()
	return out
}

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for n := range sum(sq(odd(echo(nums)))) {
		fmt.Println(n) // 165
	}
	// 类似于我们执行了Unix/Linux命令： echo $nums | sq | sum

	for n := range pipeline(nums, echo, odd, sq, sum) {
		fmt.Println(n) // 165
	}

	// Fan in 大数组里面的质数分段求和
	nums3 := makeRange(1, 1000)
	in := echo(nums3)

	const nProcessor = 4
	var chans [nProcessor]<-chan int
	for i := range chans {
		chans[i] = sum(prime(in))
	}
	for n := range sum(merge(chans[:])) {
		fmt.Println(n)
	}
}

func merge(cs []<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func prime(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if is_prime(n) {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

func is_prime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

// 定义一个函数类型
type EchoFunc func([]int) <-chan int
type PipeFunc func(<-chan int) <-chan int

// 代理函数
func pipeline(nums []int, echo EchoFunc, pipeFns ...PipeFunc) <-chan int {
	ch := echo(nums)
	for i := range pipeFns {
		ch = pipeFns[i](ch)
	}
	return ch
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + 1
	}
	return a
}
