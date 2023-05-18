package main

import (
	"fmt"
	"strconv"
	"strings"
)

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func main() {

	l1 := &LinkNode{
		Val: 1,
		Next: &LinkNode{
			Val: 22,
			Next: &LinkNode{
				Val: 33,
				Next: &LinkNode{
					Val: 11,
				},
			},
		},
	}
	l1.print()
	r := reverse(l1)
	r.print()

}

func (l LinkNode) print() {
	s := []string{}
	for {
		if l.Next == nil {
			i := strconv.Itoa(l.Val)
			s = append(s, i)
			break
		}
		i := strconv.Itoa(l.Val)
		s = append(s, i)
		l = *l.Next
	}
	res := strings.Join(s, " -> ")
	fmt.Println(res)
}

func reverse(head *LinkNode) *LinkNode {
	var prev, curr, next *LinkNode
	curr, next = head, head

	for next != nil {
		next = next.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
