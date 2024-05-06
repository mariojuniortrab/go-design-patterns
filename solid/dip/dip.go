package main

import "fmt"

// type Relationship int

// const (
// 	Parent Relationship = iota
// 	Child
// 	Sibling
// )

// type Person struct {
// 	name string
// }

// type Info struct {
// 	from         *Person
// 	relationship Relationship
// 	to           *Person
// }

// // low-level module
// type Relationships struct {
// 	relations []Info
// }

// func (r *Relationships) AddParentAndChild(parent, child *Person) {
// 	r.relations = append(r.relations, Info{parent, Parent, child})
// 	r.relations = append(r.relations, Info{child, Child, parent})
// }

// // high-level module
// type Research struct {
// 	//Break dip
// 	relationships Relationships
// }

// func (r *Research) Investigate(name string) {
// 	relations := r.relationships.relations

// 	for _, rel := range relations {
// 		if rel.from.name == name && rel.relationship == Parent {
// 			fmt.Println(name, " has a child called ", rel.to.name)
// 		}
// 	}
// }

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// low-level module
type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

type Relationships struct {
	relations []Info
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range r.relations {
		if v.from.name == name && v.relationship == Parent {
			result = append(result, r.relations[i].to)
		}
	}

	return result
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

// high-level module
type Research struct {
	browser RelationshipBrowser
}

func (r *Research) Investigate(name string) {

	for _, p := range r.browser.FindAllChildrenOf(name) {
		fmt.Println(name, " has a child called ", p.name)
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	r := Research{&relationships}
	r.Investigate("John")
}
