package main

import (
	"fmt"
	"slices"
)

func main() {
	var i1 Item
	fmt.Println(i1)
	fmt.Printf("i1: %#v\n", i1)

	i2 := Item{1, 2}
	fmt.Printf("i2: %#v\n", i2)

	i3 := Item{
		Y: 10,
		// X: 20
	}
	fmt.Printf("i3: %#v\n", i3)
	fmt.Println(NewItem(10, 20))
	fmt.Println(NewItem(10, -20))

	i3.Move(100, 200)
	fmt.Printf("i3 (move): %#v\n", i3)

	p1 := Player{
		Name: "Parzival",
		Item: Item{500, 300},
	}

	fmt.Printf("p1: %#v\n", p1)
	fmt.Printf("p1.X: %#v\n", p1.X)

	ms := []mover{
		&i1,
		&p1,
		&i2,
	}

	moveAll(ms, 0, 0)
	for _, m := range ms {
		fmt.Println(m)
	}

	k := Jade
	fmt.Println("k:", k)

	john := Player{Name: "John"}
	fmt.Printf("FoundKey returned %v. Person: %#v\n", john.FoundKey(Jade), john)
	mike := Player{Name: "Mike"}
	fmt.Printf("FoundKey returned %v. Person: %#v\n", mike.FoundKey(42), mike)
}

func (k Key) String() string {
	switch k {
	case Jade:
		return "jade"
	case Copper:
		return "copper"
	case Crystal:
		return "cystal"
	}

	return fmt.Sprintf("<Key %d>", k)
}

// Go's version of enum
const (
	Jade Key = iota + 1
	Copper
	Crystal
	invalidKey
)

type Key byte

func moveAll(ms []mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

type mover interface {
	Move(x, y int)
}

type Player struct {
	Keys []Key
	Name string
	Item
}

func (p *Player) FoundKey(k Key) error {
	if k == 0 || k >= invalidKey {
		return fmt.Errorf("invalid key: %#v", k)
	}

	if !slices.Contains(p.Keys, k) {
		p.Keys = append(p.Keys, k)
	}

	return nil
}

// Item is an item in the game
type Item struct {
	X int
	Y int
}

func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}

// NewItem returns a pointer to a new Item
func NewItem(x, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of bounds %d/%d", x, y, maxX, maxY)
	}

	i := Item{
		X: x,
		Y: y,
	}

	return &i, nil
}

const (
	maxX = 1000
	maxY = 600
)
