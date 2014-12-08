### Lazy loaded Filtering

Filter arrays/slices using lazy loaded functions.

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

See main_test.go for more examples.