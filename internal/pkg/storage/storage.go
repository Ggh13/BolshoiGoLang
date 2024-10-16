package storage

import (
	"fmt"
	"strconv"

	"go.uber.org/zap"
)

type Value struct {
	s    string
	kind string
}

type Storage struct {
	inner  map[string]Value
	sql    map[string][]int
	logger *zap.Logger
}

func NewStorage() (Storage, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return Storage{}, err
	}

	defer logger.Sync()
	logger.Info("created new storage")

	return Storage{
		inner:  make(map[string]Value),
		sql:    make(map[string][]int),
		logger: logger,
	}, nil
}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}

func (r Storage) LPUSH(key string, val []int) (err string) {
	fmt.Println("Start LLLLPush")
	for i := range val {
		_, ok := r.sql[key]
		if !ok {
			r.sql[key] = []int{val[i]}

		} else {
			temp := reverseInts(r.sql[key])
			r.sql[key] = reverseInts(append(temp, val[i]))

		}
		fmt.Println(r.sql[key])

	}
	fmt.Println(r.sql[key])
	return ""

}

func (r Storage) RPUSH(key string, val []int) (err string) {
	fmt.Println("Start RRRRPush")
	for i := range val {
		_, ok := r.sql[key]
		if !ok {
			r.sql[key] = []int{val[i]}

		} else {

			r.sql[key] = append(r.sql[key], val[i])

		}
		fmt.Println(r.sql[key])

	}
	fmt.Println(r.sql[key])
	return ""

}

func (r Storage) LPOP(key string, val []int) (toPopV []int, err string) {

	toPop := []int{}

	if len(val) == 0 {

		toPop = append(toPop, r.sql[key][0])
		r.sql[key] = r.sql[key][1:]
	} else if len(val) == 1 {
		if val[0] >= len(r.sql[key]) {

			toPop = append(toPop, r.sql[key]...)
			r.sql[key] = []int{}
			return toPop, ""
		}

		toPop = r.sql[key][:val[0]]

		r.sql[key] = r.sql[key][val[0] : len(r.sql[key])-1]

	} else {

		toPop = append(toPop, r.sql[key][val[0]:val[1]+1]...)

		r.sql[key] = append(r.sql[key][0:val[0]], r.sql[key][val[1]+1:len(r.sql[key])]...)

	}

	return toPop, ""

}

func (r Storage) RADDTOSET(key string, val []int) (err string) {
	fmt.Println("Start RRRRPush")
	for i := range val {
		_, ok := r.sql[key]
		if !ok {
			r.sql[key] = []int{val[i]}

		} else {
			flag := false
			for _, v := range r.sql[key] {
				if v == val[i] {
					flag = true
					break
				}
			}
			if !flag {
				r.sql[key] = append(r.sql[key], val[i])
			}

		}

	}
	fmt.Println(r.sql[key])
	return ""
}
func (r Storage) LSET(key string, ind int, elem int) (err string) {
	_, ok := r.sql[key]
	if ok {

		if len(r.sql[key]) > ind {
			r.sql[key][ind] = elem

			return fmt.Sprintf("(integer) %d", elem)
		} else {
			return "Index out of range"
		}
	}
	return "Index out of range1"
}

func (r Storage) LGET(key string, ind int) (res int, err string) {
	_, ok := r.sql[key]
	if ok {
		fmt.Println("r.sql[key]")
		fmt.Println(len(r.sql[key]))
		if len(r.sql[key]) > ind {
			fmt.Println(r.sql[key])
			return r.sql[key][ind], ""
		} else {
			return -1, "Index out of range"
		}
	}
	return -1, "Index out of range1"
}

func (r Storage) Set(key, value string) {
	switch GetType(value) {
	case "D":
		r.inner[key] = Value{s: value, kind: "D"}
	case "Fl64":
		r.inner[key] = Value{s: value, kind: "Fl64"}
	case "S":
		r.inner[key] = Value{s: value, kind: "S"}
	}

	r.logger.Info("key set")
	r.logger.Sync()
}

func (r Storage) Get(key string) *string {
	res, ok := r.inner[key]
	if !ok {
		return nil
	}

	return &res.s
}

func GetType(value string) string {
	if _, err := strconv.Atoi(value); err == nil {
		return "D"
	}
	if _, err := strconv.ParseFloat(value, 64); err == nil {
		return "Fl64"
	}
	return "S"
}
