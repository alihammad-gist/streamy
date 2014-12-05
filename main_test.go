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
	} else {
		for _, idx := range m {
			log.Println(d[idx])
		}
	}
}
