package main

import (
	"fmt"
	"io"
	"strings"
	"time"
)

// WordCount function
func WordCount(s string) (m map[string]int) {
	m = make(map[string]int)
	for _, value := range strings.Fields(s) {
		m[value]++
	}
	return
}

func swap(x, y interface{}) (interface{}, interface{}) {
	return y, x
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// Pic function
func Pic(dx, dy int) (r [][]uint8) {
	for i := 0; i < dy; i++ {
		r = append(r, make([]uint8, 0, 2))
		for j := 0; j < dx; j++ {
			r[i] = append(r[i], uint8((i+j)/2))
		}
	}
	return
}
func fibonacci() func() int {
	f1orig, f1, f2 := 0, 0, 1
	f := 0
	return func() int {
		switch f {
		case 0:
			f++
			return f1
		case 1:
			f++
			return f2
		default:
			f1orig = f1
			f1, f2 = f2, f1orig+f2
			return f1orig + f1

		}
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v",
		float64(e))
}

// Sqrt function
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(-1)
	}
	z := 1.1
	for i := 1; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		//fmt.Println(i, z)
	}
	return z, nil
}

type MyReader struct{}

func (r MyReader) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = 65
	}
	n = len(b)
	return
}

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r.r.Read(b)
	for i := range b {
		if (b[i] >= 'A' && b[i] < 'N') || (b[i] >= 'a' && b[i] < 'n') {
			b[i] += 13
		} else if (b[i] > 'M' && b[i] <= 'Z') || (b[i] > 'm' && b[i] <= 'z') {
			b[i] -= 13
		}
	}
	n = len(b)
	return
}

func main() {
	// fmt.Printf("Hello, world.\n")
	// fmt.Println(swap("hello", "world"))
	// fmt.Println(swap(20, 10))
	// pic.Show(Pic)
	// wc.Test(WordCount)
	// fmt.Println(Sqrt(2))
	// fmt.Println(Sqrt(-2))
	// f := fibonacci()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(f())
	// }
	// r := strings.NewReader("Hello, Reader!")

	/* b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	reader.Validate(MyReader{})

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r) */

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}

}
