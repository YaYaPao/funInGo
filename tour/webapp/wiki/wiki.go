package wiki

import (
	"fmt"
	"os"
)

// Creating a data structure with load and save methods
// Using the net/http package to build web applications
// Using the html/template package to process HTML templates
// Using the regexp package to validate user input
// Using closures

type Page struct {
	Title string
	Body  []byte
}

// 保存文件
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

// 读取文件
func LoadFile(t string) (*Page, error) {
	filename := t + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: filename, Body: body}, nil
}

func main() {
	p := &Page{Title: "HelloWeb", Body: []byte("Produced by Go!")}
	err := p.save()
	if err != nil {
		fmt.Println("Save failed!")
	}
	p2, _ := LoadFile(p.Title)
	fmt.Println(string(p2.Body))
}
