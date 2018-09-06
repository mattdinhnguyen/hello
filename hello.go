package main

import (
	"fmt"
	"io"
	"strings"
	"sync"

	"golang.org/x/tour/tree"
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

// ErrNegativeSqrt type
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

// MyReader type
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

// Walk function
func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func walk(t *tree.Tree, ch chan int) {
	if t != nil {
		walk(t.Left, ch)
		ch <- t.Value
		walk(t.Right, ch)
	}
}

// Same function
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}
	return true
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

// Cache type
type Cache struct {
	visited map[string]bool
	mux     sync.Mutex
}

// Fetcher interface
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan response, cache *Cache) {
	defer close(ch)
	if depth <= 0 {
		return
	}
	cache.mux.Lock()
	if cache.visited[url] {
		cache.mux.Unlock()
		return
	}
	cache.visited[url] = true
	cache.mux.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	ch <- response{url, body}
	result := make([]chan response, len(urls))
	for i, u := range urls {
		result[i] = make(chan response)
		go Crawl(u, depth-1, fetcher, result[i], cache)
	}

	for i := range result {
		for resp := range result[i] {
			ch <- resp
		}
	}
	return
}

type response struct {
	url  string
	body string
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
	io.Copy(os.Stdout, &r)

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
	ch := make(chan int)
	go Walk(tree.New(2), ch)
	for v := range ch {
		fmt.Print(v)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1001; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey")) */
	var ch = make(chan response)
	var cache = &Cache{visited: make(map[string]bool)}
	go Crawl("https://golang.org/", 4, fetcher, ch, cache)
	for resp := range ch {
		fmt.Printf("found: %s %q\n", resp.url, resp.body)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
