package streamy_test

import (
	"github.com/alihammad-gist/streamy"
	"log"
	"testing"
)

func TestGetFirst(t *testing.T) {
	d := []int{1, 2, 3, 4, 5, 6, 7, 6}
	stream := streamy.New(len(d)).Filter(func(i int) bool {
		if d[i]%2 == 0 {
			return true
		}
		return false
	}).Filter(func(i int) bool {
		if d[i] == 2 {
			return false
		}
		return true
	})

	if val, ok := stream.GetFirst(); !ok {
		t.Log(val)
		t.Fail()
	} else {
		if d[val] != 4 {
			t.Log(d[val])
			t.Fail()
		}
	}
}

func TestGetAll(t *testing.T) {
	d := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	m := streamy.New(len(d)).Filter(func(i int) bool {
		if d[i]%2 == 1 {
			return true
		}
		return false
	}).Filter(func(i int) bool {
		if d[i] == 5 {
			return false
		}
		return true
	}).GetAll()

	if len(m) != 4 {
		t.Fail()
	}
}

func TestMap(t *testing.T) {
	d := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	var mp []int
	streamy.New(len(d)).Filter(
		filterEven(&d),
	).Filter(
		func(i int) bool {
			mp = append(mp, d[i]*10)
			return true
		},
	).GetAll()
}

func filterEven(data *[]int) streamy.Filter {
	return func(i int) bool {
		if (*data)[i]%2 == 0 {
			return true
		}
		return false
	}
}

func TestGetLast(t *testing.T) {
	d := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 21}
	n, ok := streamy.New(len(d)).Filter(filterEven(&d)).GetLast()
	if d[n] != 8 {
		log.Fatal("wrong number")
		t.Fail()
	}
	if !ok {
		log.Fatal("Returned nothing")
		t.Fail()
	}
}

func TestGetN(t *testing.T) {
	d := []int{1, 44, 8, 2, 3, 4, 5, 7}
	evens := streamy.New(len(d)).Filter(filterEven(&d)).GetN(2, streamy.OrderDesc)
	if len(evens) != 2 {
		log.Println("Wrong number of evens returned", len(evens))
		t.Fail()
	}
	if d[evens[0]] != 4 {
		log.Println("Wrong order")
		t.Fail()
	}
	if d[evens[1]] != 2 {
		log.Println("Wrong order")
		t.Fail()
	}

	d = []int{1, 4}
	evens = streamy.New(len(d)).Filter(filterEven(&d)).GetN(0, streamy.OrderAsc)
	if len(evens) != 1 {
		log.Println("Wrong number of evens returned")
		t.Fail()
	}
	if d[evens[0]] != 4 {
		log.Println("Wrong number")
		t.Fail()
	}
}
