package main

import (
	"awesomeProject/model"
	"fmt"
	"sync"
)

func search(id int) {
	adminObject2 := model.GetAdminById(id)
	fmt.Println(adminObject2)
}

func main() {
	arrIds := []int{2, 49}
	search(1)

	var wg sync.WaitGroup
	wg.Add(len(arrIds))

	for i, id := range arrIds {
		fmt.Println(i)
		go func(i1 int, id1 int) {
			defer wg.Done()
			fmt.Println("id=", id1)
			search(id1)
		}(i, id)
	}

	wg.Wait()

	fmt.Println("可以继续别的逻辑了")

}
