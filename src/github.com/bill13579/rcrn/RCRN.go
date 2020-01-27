package rcrn

import (
	"sync"
	"runtime"
	"math"
)

var wg sync.WaitGroup = sync.WaitGroup{}
var start sync.WaitGroup = sync.WaitGroup{}
var r int = 1001

func race() int {
	var t int = 1
	for i := 1; i < r; i++ {
		t *= i
	}
	return t
}

func r1(bit *int8) {
	runtime.LockOSThread()
	start.Wait()
	race()
	(*bit) = 0
	runtime.UnlockOSThread()
	wg.Done()
}

func r2(bit *int8) {
	runtime.LockOSThread()
	start.Wait()
	race()
	(*bit) = 1
	runtime.UnlockOSThread()
	wg.Done()
}

func RandBits(n int) []int8 {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	var bits = make([]int8, n, n)
	for i := 0; i < n; i++ {
		var bit int8 = -1;
		wg.Add(2)
		start.Add(1)
		go r1(&bit)
		go r2(&bit)
		for j := 0; j < 100; j++ {
		}
		start.Done()
		wg.Wait()
		bits[i] = bit
	}
	return bits
}

func toByte(bits []int8) byte {
	var b byte = 0
	for i := 0; i < len(bits); i++ {
		b += byte(int(math.Pow(2, float64(i))) * int(bits[i]))
	}
	return b
}

func RandBytes(n int) []byte {
	var bytes = make([]byte, n, n)
	for i := 0; i < n; i++ {
		bytes[i] = toByte(RandBits(8))
	}
	return bytes
}

func Rand(n int) []byte {
	var bytes = make([]byte, n, n)
	for i := 0; i < n; i++ {
		bytes[i] = toByte(RandBits(8)) / 255
	}
	return bytes
}
