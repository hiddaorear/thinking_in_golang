package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string

func fetchFromServer(query string) Result {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return Result(fmt.Sprintf("result for %q \n", query))
}

func main() {
	datas := []string{"data1", "data2", "data3"}
	results := make([]Result, len(datas))

	for _, data := range datas {
		res := fetchFromServer(data)
		results = append(results, res)
	}

	for _, r := range results {
		fmt.Println(r)
	}

}
