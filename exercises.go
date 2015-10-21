//**********************

package main

import "fmt"

type Vertex struct{
	x int
	y int
}

func (v *Vertex) Reset() {
	v.x = 0
	v.y = 0
}

func Whatever(v Vertex) int { // Normal method
	return v.x * v.y
}

func main() {
	var v Vertex
	// W is V
	v.x = 2
	v.y = 3
	fmt.Println(Whatever(v))	// Normal method
	
}

//************ POINTERS
package main

import "fmt"

type Vertex struct{
	x int
	y int
}

func (v *Vertex) Reset() {
	v.x = 0
	v.y = 0
}


func main() {
	var v Vertex
	// W is V
	w:=v
	w.x = 15
	w.y = 15

	fmt.Println(v, w) // {0 0} {15 15}
	fmt.Printf("Type v: %T, value %d || Type w: %T, value %d \n",v, v, w, w) // {0 0} {15 15}
	
	z:=&v
	z.x = 3
	z.y = 3
	
	fmt.Printf("Type z: %T, value: %d || Type v: %T, value: %d \n",z, z, v, v) // {0 0} {15 15}
	fmt.Printf("Type z: %T, value through pointer: %d || Type v: %T, value: %d \n",z, *z, v, v) // {0 0} {15 15}	
	
}



// Methods with pointer receivers
// https://golang.org/doc/faq#methods_on_values_or_pointers

import "fmt"

type Mutatable struct {
    a int
    b int
}

func (m Mutatable) StayTheSame() { // Passing it by value
    m.a = 5
    m.b = 7
}

func (m *Mutatable) Mutate() { // Passing it by reference
    m.a = 5
    m.b = 7
}

func main() {
    m := &Mutatable{0, 0}
	fmt.Println(*m) // {0, 0}
    m.StayTheSame()
    fmt.Println(*m) // {0, 0}
    m.Mutate()
    fmt.Println(*m) // {5, 7}
}


// exercise-reader.go https://tour.golang.org/methods/11

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(b []byte) (n int, err error) {
	b[0] = 'A'
    return 1, nil  // My reader reads one byte the A :)
}
// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}

/******************************************************/

// rot-reader https://tour.golang.org/methods/12

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(p []byte) (n int, err error) {
	
    n,err = rot.r.Read(p)
	
	for i := 0; i < len(p); i++ {
		if (p[i] >= 'A' && p[i] < 'N') || (p[i] >='a' && p[i] < 'n') {
			p[i] += 13
		} else if (p[i] > 'M' && p[i] <= 'Z') || (p[i] > 'm' && p[i] <= 'z'){
			p[i] -= 13
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
