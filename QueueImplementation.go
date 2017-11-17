// It is assumed that stack library functions are available

package main

import (
	"fmt"
	"path-to-stack-library/stack""
)

type Queue struct {
	mystack1 Stack
	mystack2 Stack
	next *Queue;
}

func NewQueue() *Queue {
	q := New(Queue)
	q.mystack1:= New()
	q.mystack2:= New()

	return q
}

func (q *Queue) Enqueue(e interface{}) *Queue {
	// push elements to stack1 
	q.mystack1.push(e, interface{})  
}

func (q *Queue) Dequeue() (e interface{}) {
	// if both stack1 and stack2 are empty then log debug
	
	if q.mystack1.len() == 0 && q.mystack2.len() == 0 	{
		fmt.Printf("Both stack empty ")
		return nil
	}
	
  1) If both stacks are empty then error.
  2) If stack2 is empty
       While stack1 is not empty, push everything from stack1 to stack2.
  3) Pop the element from stack2 and return it.
  
	if q.mystack2.len == 0 	{
		for q.mystack1.len > 0 {
			q.mystack2.push(q.mystack1.pop())
		}
	}	

	val = q.mystack2.pop()
	return val
}

func main() {
	q := NewQueue()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	
	for q.Size() > 0 {
		fmt.Println(q.Dequeue())
	}

}