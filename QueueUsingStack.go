//Stacks & Queues

// My Initial approach was this
//Problem :-
//Implement a Queue data structure using two stacks.

//Solution :-
//Consider two stacks stack1 & stack2s2
//Enqueue Operation :: Simply push the element onto s1.
//Dequeue Operation :: Transfer all elements from s1 onto s2. Pop the top element from s2. Transfer remaining elements from s2 back to s1.

//Approach 2: Push all the elements to stack1 when Dequeue operation is called then pop all the element from stack1 to stack2.
//Then pop the top element from the stack2 and return
//After optimization This is the solution
//  1) If both stacks are empty then error.
//  2) If stack2 is empty
//       While stack1 is not empty, push everything from stack1 to stack2.
//  3) Pop the element from stack2 and return it

package main

import (
	"fmt"
)

type Stack struct {
	top  *Node
	size int
}

type Node struct {
	value interface{}
	next  *Node
}

func (s *Stack) Length() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Push(val interface{}) {
	s.top = &Node{val, s.top}
	s.size++
}

func (s *Stack) Peek() interface{} {
	return s.top.value
}

func (s *Stack) Pop() (val interface{}) {
	if s.size > 0 {
		val, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return ""
}

type Queue struct {
	mystack1 Stack
	mystack2 Stack
	next     *Queue
}

func NewQueue() *Queue {
	q := new(Queue)
	return q
}

func (q *Queue) Enqueue(e interface{}) *Queue {
	// push elements to stack1
	q.mystack1.Push(e)
	return q
}

func (q *Queue) Dequeue() (e interface{}) {
	// if both stack1 and stack2 are empty then log debug

	if q.mystack1.Length() == 0 && q.mystack2.Length() == 0 {
		fmt.Printf("Both stack empty ")
		return nil
	}

	//  1) If both stacks are empty then error.
	//  2) If stack2 is empty
	//       While stack1 is not empty, push everything from stack1 to stack2.
	//  3) Pop the element from stack2 and return it.

	if q.mystack2.Length() == 0 {
		for q.mystack1.Length() > 0 {
			q.mystack2.Push(q.mystack1.Pop())
		}
	}
	// store the value to return 
	var x = q.mystack2.Pop()
	
// Let stack1 will have the all values always 

	for q.mystack2.Length() > 0 {
		q.mystack1.Push(q.mystack2.Pop())
	}	
	
	return x 
}

// For testing the Queue operation.

func main() {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Dequeue()   // pop  value 1
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Dequeue() // pop value 2
	q.Enqueue(6)
	q.Enqueue(7)

	// pop all the remaining elements 
	// assuming 
	for n := 0; n < 5; n++ {
		fmt.Println(q.Dequeue())
	}

}

