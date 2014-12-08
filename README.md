### Streamy - Lazy loaded Filtering

Filter arrays/slices using lazy loaded functions.


#### GetFirst / GetLast
Use GetLast to get Last value passes the filter successfully or use GetFirst
to get First value that passes through provided filters

```go
package main

import (
	"fmt"
	"github.com/alihammad-gist/streamy"
)

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Get last even number which is also divisible by 4 [made up problem using 2 filters]
	l, ok := streamy.New(len(slice)).Filter(
		func(i int) bool {
			fmt.Println("Even check: ", i)
			if slice[i]%2 == 0 {
				return true
			}
			return false
		},
	).Filter(
		func(i int) bool {
			fmt.Println("DivisBy 4 check: ", i)
			if slice[i]%4 == 0 {
				return true
			}
			return false
		},
	).GetLast()

	fmt.Println("ok: ", ok, " - num: ", slice[l])
}
```

Output

```
Even check:  10
DivisBy 4 check:  10
Even check:  9
Even check:  8
DivisBy 4 check:  8
ok:  true  - num:  8
```

#### GetN(resultCount, order)
GetN requires you pass in 'n' (count of matched values). Iteration stops when resultCount is achieved, pass 0 to get all matched indexes. Second argument is order the slice should be evaluated in `streamy.OrderAsc` or `streamy.OrderDesc`. Note, returned list of matched indexes will follow the same order.

```go
package main

import (
	"fmt"
	"github.com/alihammad-gist/streamy"
)

func main() {
	slice :=
		[]int{
			0,
			1,
			2,
			3,
			4,
			5,
			6,
			7,
			8, 9, 10, 11, 12, 13, 14, 15}

	// Get First 2 odd numbers which are
	// greater thn 5
	res := streamy.New(len(slice)).Filter(
		func(i int) bool {
			fmt.Println("> 5 check:", i)
			if slice[i] > 5 {
				return true
			}
			return false
		},
	).Filter(
		func(i int) bool {
			fmt.Println("odd check:", i)
			if slice[i]%2 == 1 {
				return true
			}
			return false
		},
	).GetN(2, streamy.OrderAsc)

	for _, i := range res {
		fmt.Println("Match:", slice[i])
	}
}
```

Output
```
> 5 check: 0
> 5 check: 1
> 5 check: 2
> 5 check: 3
> 5 check: 4
> 5 check: 5
> 5 check: 6
odd check: 6
> 5 check: 7
odd check: 7
> 5 check: 8
odd check: 8
> 5 check: 9
odd check: 9
Match: 7
Match: 9
```

See main_test.go for more examples.