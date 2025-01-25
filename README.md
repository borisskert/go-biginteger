# go-biginteger

This is a simple implementation of a big integer in Go. It is a simple implementation and is not optimized for performance. It is meant to be used as a learning tool.

## Usage

### Add dependency

```bash
go get github.com/borisskert/go-biginteger
```

### Example

```go
package main

import (
	"fmt"
	"github.com/borisskert/go-biginteger"
)

func main() {
	a, _ := biginteger.Of("123456789012345678901234567890")
	b, _ := biginteger.Of("987654321098765432109876543210")

	fmt.Println(a.Add(*b))      // 1111111110111111111011111111100
	fmt.Println(a.Subtract(*b)) // -864197532086419753208641975320
	fmt.Println(a.Multiply(*b)) // 121932631137021795226185032733622923332237463801111263526900
	fmt.Println(b.Divide(*a))   // 8

	fmt.Println(b.Modulo(*a))              // 9000000000900000000090
	fmt.Println(a.Power(biginteger.Two())) // 15241578753238836750495351562536198787501905199875019052100

	fmt.Println(a.CompareTo(*b))                // -1
	fmt.Println(a.IsEqualTo(biginteger.Zero())) // false
	fmt.Println(a.IsGreaterThan(*b))            // false
	fmt.Println(a.IsLessThan(*b))               // true

	fmt.Println(a.ShiftRight(2)) // 30864197253086419725308641972
	fmt.Println(b.ShiftLeft(2))  // 3950617284395061728439506172840

	fmt.Println(a.BitLength()) // 97
	fmt.Println(b.Digits())    // 30

	fmt.Println(a.String()) // "123456789012345678901234567890"
	fmt.Println(b.Negate()) // -987654321098765432109876543210
}
```

## License

MIT
