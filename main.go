package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func (Person) Write(p []byte) (int, error) {
	fmt.Print(string(p))
	return 0, nil
}
func (p *Person) String() string {
	return fmt.Sprintf("{Name:%v,Age:%v}", p.Name, p.Age)
}

func main() {
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	fmt.Printf("%#v",w)
}
