package main

import (
	"BolshiGoLang/internal/pkg/storage"
	"fmt"
)

func main() {
	s, err := storage.NewStorage()
	if err != nil {
		fmt.Println("Something broke(")
		return
	}
	mas := []int{1, 2, 3, 9}
	_ = s.LPUSH("key1", mas)

	ind := []int{1, 2}
	fmt.Println(s.LPOP("key1", ind))

	mas = []int{1, 2, 3, 9}
	_ = s.LPUSH("key1", mas)

	ind = []int{2}
	fmt.Println(s.LPOP("key1", ind))

	mas = []int{1, 2, 3, 9}
	_ = s.LPUSH("key1", mas)

	ind = []int{}
	fmt.Println(s.LPOP("key1", ind))

	mas = []int{1, 2, 3, 9, 7}
	_ = s.RADDTOSET("key1", mas)

	mas = []int{1, 2, 3, 9, 7}
	tr := s.LSET("key1", 1, 45)
	fmt.Println(tr)
	ans, _ := s.LGET("key1", 1)
	fmt.Println(ans)
	s.Set("0", "mamba")
	s.Set("1", "22")
	s.Set("2", "22.2")
	fmt.Println(*s.Get("0"), storage.GetType(*s.Get("0")))
	fmt.Println(*s.Get("1"), storage.GetType(*s.Get("1")))
	fmt.Println(*s.Get("2"), storage.GetType(*s.Get("2")))
}
