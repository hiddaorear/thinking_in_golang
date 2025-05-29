package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Result string

func fetchFromServer(query string) Result {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return Result(fmt.Sprintf("result for %q\n", query))
}

type result struct {
	m map[int]Result
	sync.Mutex
}

func main() {
	datas := []string{"data1", "data2", "data3"}

	var results result
	results.m = make(map[int]Result)

	var wg sync.WaitGroup
	for i := range datas {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			res := fetchFromServer(datas[i])
			results.Lock()
			results.m[i] = res
			results.Unlock()
		}(i)
	}
	wg.Wait()

	for _, r := range results.m {
		fmt.Println(r)
	}

}
