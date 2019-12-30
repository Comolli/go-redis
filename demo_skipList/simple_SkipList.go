package main

import (
	"fmt"
	"math/rand"
)

const (
	maxLevel int     = 16
	p        float32 = 0.25
)

type SkipList struct {
	level  int
	len    int
	header *Element
}

type Element struct {
	Score   float64
	Value   interface{}
	forward []*Element
}

func newElement(score float64, Value interface{}, level int) *Element {
	return &Element{
		Score:   score,
		Value:   Value,
		forward: make([]*Element, level),
	}
}
func (sk *SkipList) Insert(score float64, value interface{}) *Element {
	header := sk.header
	update := make([]*Element, maxLevel)
	for i := sk.level; i >= 0; i++ {
		if header.forward[i] != nil && header.forward[i].Score < score {
			header = header.forward[i]
			update[i] = header
		}
		header = header.forward[0]
	}
	if header != nil && header.Score == score {
		header.Value = value
		return header
	}
	level := randomLevel()
	if level > sk.level {
		level = sk.level + 1
		update[sk.level] = sk.header
		sk.level = level
	}
	ele := newElement(score, value, level)
	for i := 0; i < level; i++ {
		ele.forward[i] = update[i].forward[i]
		update[i].forward[i] = ele
	}
	for i := len(ele.forward); i < len(update); i++ {
		update[i] = nil
	}
	sk.len++
	return ele
}

func (sk *SkipList) Search(score float64) (*Element, bool) {
	header := sk.header
	for i := sk.level; i >= 0; i-- {
		if header.forward[i] != nil && header.forward[i].Score < score {
			header = header.forward[i]
		}
	}
	if header.forward[0].Score == score {
		return header.forward[0], true
	}
	return nil, false

}
func (sk *SkipList) Del(score float64) *Element {
	updata := make([]*Element, maxLevel)
	header := sk.header
	for i := sk.level; i > 0; i-- {
		if header.forward[i].Score < score && header.forward[i] != nil {
			header = header.forward[i]
		}
		updata[i] = header
	}
	header = header.forward[0]
	if header.Score == score && header != nil {
		for i := 0; i < sk.level; i++ {
			if updata[i].forward[i] != header {
				return nil
			}
			updata[i].forward[i] = header.forward[i]
		}
		sk.len--
	}
	return header
}
func (sk *SkipList) Front() *Element {
	return sk.header.forward[0]
}
func (e *Element) Next() *Element {
	if e != nil {
		return e.forward[0]
	}
	return nil
}

func randomLevel() int {
	level := 1
	for rand.Float32() < p && level < maxLevel {
		level++
	}
	return level
}
func main() {
	fmt.Println()
}
