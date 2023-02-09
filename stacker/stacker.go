package main

import(
    "fmt"
    "errors"
)

type Stack []interface{}

func (stack Stack) Len() int {
    return len(stack)
}

func (stack *Stack) Push(x interface{}) {
    *stack = append(*stack, x)
}

func (stack Stack) Top() (interface{}, error) {
    if len(stack) == 0 {
      return nil, errors.New("Can't Top() on empty stack")
    }

    return stack[len(stack)-1], nil
}

func (stack *Stack) Pop() (interface{}, error) {
    theStack := *stack
    if len(theStack) == 0 {
      return nil, errors.New("Can't Pop() on empty stack")
    }

    x := theStack[len(theStack)-1]
    *stack = theStack[:len(theStack)-1]
    return x, nil
}

func main() {
    var haystack Stack
    haystack.Push("hay")
    haystack.Push(-15)
    haystack.Push([]string{"pin", "clip", "needle"})
    haystack.Push(81.52)

    for {
      item, err := haystack.Pop()
      if err != nil {
        break
      }

      fmt.Println(item)
    }
}
