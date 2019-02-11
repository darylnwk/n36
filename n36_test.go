package n36_test

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/darylnwk/n36"
	"github.com/stretchr/testify/assert"
)

func ExampleN36_Iton() {
	n := n36.New(n36.CharRange62)

	fmt.Println(n.Iton(0))
	fmt.Println(n.Iton(1))
	fmt.Println(n.Iton(36))
	fmt.Println(n.Iton(62))
	fmt.Println(n.Iton(3))

	fmt.Println(n.Iton(412312313232346))
	// Output:
	// 1
	// A
	// 10
	// 3
	// 1T4ZlBAlI
}

func ExampleN36_Ntoi() {
	n := n36.New(n36.CharRange62)

	fmt.Println(n.Ntoi("0"))
	fmt.Println(n.Ntoi("1"))
	fmt.Println(n.Ntoi("A"))
	fmt.Println(n.Ntoi("1aA"))
	fmt.Println(n.Ntoi("10"))
	fmt.Println(n.Ntoi("3"))
	fmt.Println(n.Ntoi("1000000"))
	fmt.Println(n.Ntoi("---"))
	// Output:
	// 0 <nil>
	// 1 <nil>
	// 36 <nil>
	// 4500 <nil>
	// 62 <nil>
	// 3 <nil>
	// 56800235584 <nil>
	// 0 n36.Ntoi: character not part of charset
}

func TestN36_Random(t *testing.T) {
	n := n36.New(n36.CharRange62)

	assert.Len(t, n.Random(3), 3)
	assert.Regexp(t, "[0-9a-zA-Z]{20}", n.Random(20))
}

func benchN36Iton(b *testing.B, zeroes int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := n36.New(n36.CharRange62)
	x := int64(math.Pow10(zeroes))

	for i := 0; i < b.N; i++ {
		n.Iton(uint64(x + r.Int63n(x*10-x-1)))
	}
}

func BenchmarkN36_Iton1B(b *testing.B) { benchN36Iton(b, 9) }
func BenchmarkN36_Iton1T(b *testing.B) { benchN36Iton(b, 12) }
func BenchmarkN36_Iton1Q(b *testing.B) { benchN36Iton(b, 15) }
func BenchmarkN36_Iton1P(b *testing.B) { benchN36Iton(b, 18) }

func benchN36RandWithNtoi(b *testing.B, length int) {
	var d string
	n := n36.New(n36.CharRange62)

	b.Run("Random"+strconv.Itoa(length), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			d = n.Random(length)
		}
	})

	b.Run("Ntoi"+strconv.Itoa(length), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			n.Ntoi(d)
		}
	})
}

func BenchmarkN36_RandWithNtoi4(b *testing.B)   { benchN36RandWithNtoi(b, 4) }
func BenchmarkN36_RandWithNtoi8(b *testing.B)   { benchN36RandWithNtoi(b, 8) }
func BenchmarkN36_RandWithNtoi16(b *testing.B)  { benchN36RandWithNtoi(b, 16) }
func BenchmarkN36_RandWithNtoi32(b *testing.B)  { benchN36RandWithNtoi(b, 32) }
func BenchmarkN36_RandWithNtoi256(b *testing.B) { benchN36RandWithNtoi(b, 256) }
