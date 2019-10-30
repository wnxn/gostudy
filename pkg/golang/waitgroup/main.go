package main

import "sync"

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Add(3)
	wg.Done()
	wg.Done()
}
