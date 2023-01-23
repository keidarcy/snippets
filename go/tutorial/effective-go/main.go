package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // import for side effect
	"os"
	"sort"
)

// By convention, the global declarations to silence import errors should come right after the imports and be commented
var _ = sql.OpenDB

func main() {
	// methods()
	// interface1()
	// interface2()
	// blankIdentifier()
	embedding()
}

type Job struct {
	Command string
	*log.Logger
}

func NewJob(command string, logger *log.Logger) *Job {
	return &Job{command, logger}
}

func (job *Job) Printf(format string, args ...interface{}) {
	job.Logger.Printf("%q: %s", job.Command, fmt.Sprintf(format, args...))
}

func embedding() {
	job := &Job{"HELLO", log.New(os.Stderr, "Job: ", log.Ldate)}
	job.Printf("[bad job]")
}

type A struct {
	name string
}

func (a A) Marshaler() {
}

func (a *A) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.name)
}

func blankIdentifier() {
	var _ json.Marshaler = (*A)(nil)
	v := A{name: "bad"}
	v.Marshaler()

	// Interface checks
	var a interface{}
	if _, ok := a.(json.Marshaler); ok {
		fmt.Printf("value %v of type %T implements json.Marshaler\n", a, a)
	}
}

func interface2() {
	// ctr := &Counter{0}
	// http.Handle("/counter", ctr)

	// var ctr Counter = 0
	// http.Handle("/counter", &ctr)

	// var c Chan
	// http.Handle("/counter", c)

	http.Handle("/args", HandlerFunc(ArgServe))
}

// type Counter struct {
//     n int
// }
// func (ctr Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//     ctr.n++
//     fmt.Fprintf(w, "counter = %d\n", ctr.n)
// }

// type Counter int
// func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//     *ctr++
//     fmt.Fprintf(w, "counter = %d\n", *ctr)
// }

// type Chan chan *http.Request

// func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	ch <- req
// 	fmt.Fprint(w, "notification sent")
// }

// write a method for a function
// check http.HandlerFunc
type HandlerFunc func(http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	f(w, req)
}

func ArgServe(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, os.Args)
}

func interface1() {
	var s Sequence = []int{1, 2, 3}
	fmt.Printf("s: %v\n", s)
	fmt.Println()
}

type Sequence []int

func (a Sequence) Copy() Sequence {
	copy := make(Sequence, 0, len(a))
	return append(copy, a...)
}
func (a Sequence) Len() int           { return len(a) }
func (a Sequence) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Sequence) Less(i, j int) bool { return a[i] < a[j] }

// func (s Sequence) String() string {
// 	s = s.Copy()
// 	sort.Sort(s)
// 	str := "["
// 	for i, ele := range s {
// 		if i > 0 {
// 			str += " "
// 		}
// 		str += fmt.Sprint(ele)
// 	}
// 	str += "]"
// 	return str
// }

// func (s Sequence) String() string {
// 	s = s.Copy()
// 	sort.Sort(s)
// 	return fmt.Sprint([]int(s))
// }

func (s Sequence) String() string {
	s = s.Copy()
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}

func methods() {
	a := "Helo, world!"
	var b ByteSlice = []byte(a)
	// var b ByteSlice
	// b = []byte(a)
	_, error := b.Write([]byte(" Morning"))
	if error != nil {
		panic(error)
	}
	fmt.Printf("a: %v\n", string(a))
	fmt.Printf("b: %v\n", string(b))
	fmt.Fprintf(&b, "This hour has %d days\n", 7)
	// fmt.Fprint(os.Stdout, "Hello 2023", "\n")
}

type ByteSlice []byte

func (p *ByteSlice) Write(data []byte) (int, error) {
	slice := *p
	l := len(slice)
	// realocate array
	if cap(slice) < l+len(data) {
		newSlice := make([]byte, ((l + len(data)) * 2))
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	copy(slice[l:], data)
	*p = slice
	return len(data), nil
}

func Append(slice, data []byte) []byte {
	l := len(slice)
	// realocate array
	if cap(slice) < l+len(data) {
		newSlice := make([]byte, ((l + len(data)) * 2))
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : l+len(data)]
	copy(slice[l:], data)
	return slice
}
