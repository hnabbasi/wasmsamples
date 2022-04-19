package main

import (
	"syscall/js"
)

func main() {
	println("GO package loaded")

	c := make(chan struct{}, 0)

	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("square", js.FuncOf(square))
	js.Global().Set("fibGo", js.FuncOf(fib))
	js.Global().Set("fibGoSlow", js.FuncOf(fibSlow))

	<-c
}

func add(this js.Value, p []js.Value) interface{} {
	sum := 0
	for _, v := range p {
		sum = sum + v.Int()
	}
	return js.ValueOf(sum)
}

func square(this js.Value, p []js.Value) interface{} {
	val := js.ValueOf(p[0]).Int()
	return js.ValueOf(val * val)
}

func fib(this js.Value, p []js.Value) interface{} {
	val := js.ValueOf(p[0]).Int()
	cache := make(map[int]int)
	return calculateFib(val, cache)
}

func calculateFib(n int, cache map[int]int) int {
	if v, ok := cache[n]; ok {
		return v
	}

	if n <= 1 {
		cache[n] = n
		return n
	}

	result := calculateFib(n-1, cache) + calculateFib(n-2, cache)
	cache[n] = result
	return result
}

func fibSlow(this js.Value, p []js.Value) interface{} {
	return calculateFibSlow(js.ValueOf(p[0]).Int())
}

func calculateFibSlow(n int) int {
	if n <= 1 {
		return n
	}

	return calculateFibSlow(n-1) + calculateFibSlow(n-2)
}
