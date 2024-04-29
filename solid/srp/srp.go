package main

import (
	"fmt"
	"os"
	"strings"
)

type Recordable interface {
	ToString() string
}

type Journal struct {
	entries []string
}

func (j Journal) ToString() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) addEntry(text string) {
	entry := fmt.Sprintf("%d) %s", len(j.entries)+1, text)
	j.entries = append(j.entries, entry)
}

func (j *Journal) removeEntry(index int) {
	// ...
}

// separation of concerns
// func (j Journal) Save(filename string) {
// 	os.WriteFile(filename, []byte(j.toString()), 0644)
// }

// func (j *Journal) Load(filename string) {

// }

// func (j *Journal) LoadFromWeb(url *url.URL) {

// }

type Persistence struct{}

func (p Persistence) SaveToAFile(r Recordable, filename string) {
	os.WriteFile(filename, []byte(r.ToString()), 0644)
}

func main() {
	j := Journal{}

	j.addEntry("entry 1")
	j.addEntry("entry 2")
	j.addEntry("entry 3")

	p := Persistence{}

	p.SaveToAFile(j, "text.txt")
}
