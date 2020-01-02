package redis

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

const maxN = 1 << 16

func TestInsertAndSearch(t *testing.T) {
	list := NewSkiplist()
	fmt.Println(maxN)
	r := createStringObject("1")
	node := list.Insert(1.0, r)
	if list.tail != node {
		t.Error("node")
	}
	if list.Search(1.0) != r {
		t.Error("not search r")
	}
	if list.Search(1.1) != nil {
		t.Error("search")
	}

	skipListchann := make(chan *skiplistNode,maxN+maxN)
	for i := 0; i < maxN; i++ {
		pp := list.Insert(float64(i), createStringObject(strconv.Itoa(i)))
		skipListchann <- pp

	}
	wait:=make(chan bool)
	go func(i int) {
		for {
			select {
			case sdf := <-skipListchann:
				if list.Search(float64(i)).getString()==sdf.robj.getString(){
					fmt.Println(i)
				}
					i++
					if i==maxN{
						wait<-true
					}
			}
		}
	}(0)
	<-wait

	list = NewSkiplist()
	rList := rand.Perm(maxN)
	fmt.Println(rList[0:7])
	skipListchann1 := make(chan *skiplistNode,maxN+maxN)
	for _, e := range rList {
		pp:=list.Insert(float64(e), createStringObject(strconv.Itoa(e)))
		skipListchann1 <- pp
	}
	go func(i int) {
		for{
			select {
			case temp:=<- skipListchann1:
				if list.Search(float64(rList[i])).getString()==temp.robj.getString(){
					fmt.Println("random",rList[i])
				}
				i++
				if i==maxN{
					wait<-true
				}
			}
		}
	}(0)
	<-wait
	// Test at random positions in the list.

}

func TestDelete(t *testing.T) {
	list := NewSkiplist()

	a := createStringObject("a")
	if list.Delete(1, a) != 0 {
		t.Fail()
	}
	list.Insert(1.0, a)

	if list.Delete(1.0, a) != 1 {
		t.Fail()
	}
	if list.tail != nil {
		t.Fail()
	}
}