package solutions

import (
	ds "dsa-go/ds"
	"strconv"
	"strings"
	"unicode"
)

type Array struct{}

type Matrix struct{}

type Stack struct{}

func (s *Stack) HammingWeight(n int) (int, string) {
	count := 0
	stack := []int{}
	for n != 0 {
		remainder := n % 2
		if remainder == 1 {
			count += 1
		}
		stack = append(stack, remainder)
		n /= 2
	}
	var binary strings.Builder
	binary.Grow(len(stack)) // Preallocate memory for the builder
	for i := len(stack) - 1; i > -1; i-- {
		binary.WriteString(strconv.Itoa(stack[i]))
	}
	return count, binary.String()
}

func (s *Stack) IsValid(s1 string) bool {
	stack := ds.NewStack[byte]()
	parenMap := map[byte]byte{
		']': '[',
		')': '(',
		'}': '{',
	}
	for i := 0; i < len(s1); i++ {
		token := s1[i]
		if token == '(' || token == '{' || token == '[' {
			stack.Push(token)
		}
		if token == ')' || token == '}' || token == ']' {
			if stack.IsEmpty() {
				return false
			} else {
				openParen := parenMap[token]
				stackOpen := stack.Pop().Value()
				if openParen != stackOpen {
					return false
				}
			}
		}
	}
	return stack.IsEmpty()
}

func (s *Stack) IsBalancedParens(str string) bool {
	// If we're an empty string, we're technically "balanced"
	if len(str) == 0 {
		return true
	}
	// Initialize the dynamic slice-stack
	stack := []byte{}
	// Initialize the dictionary
	dict := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}
	for i := 0; i < len(str); i++ {
		char := str[i]
		if char == '{' || char == '(' || char == '[' {
			// If we're an opening paren, we push
			// to the top of the stack
			stack = append(stack, char)
		} else if char == '}' || char == ')' || char == ']' {
			// The char is closing bracket. First, "peek"
			// if the stack is empty. If its empty, we're
			// unbalanced
			if len(stack) == 0 {
				return false
			} else {
				// otherwise, peek the top of the stack.
				// pull the matching opening paren from
				// the dict
				top := stack[len(stack)-1]
				open := dict[char]
				// If they're not the same opening paren,
				// we aren't balanced
				if top != open {
					return false
				}
				// Pop the top
				stack = stack[:len(stack)-1]
			}
		}
	}
	// The stack has to be empty
	return len(stack) == 0
}

func (s *Stack) ReverseString(str string) string {
	bytes := []byte(str)
	n := len(bytes)
	stack := ds.NewStack[byte]()
	for i := 0; i < n; i++ {
		stack.Push(bytes[i])
	}
	for i := 0; i < n; i++ {
		_byte := stack.Pop().Value()
		bytes[i] = _byte
	}
	return string(bytes)
}

func (s *Stack) ReverseStr(str string) string {
	// Initialize the stack to store chars as we make our pass
	stack := []byte{}
	for i := len(str) - 1; i > -1; i-- {
		stack = append(stack, str[i])
	}
	return string(stack)
}

func (s *Stack) DecimalToBinary(num int) string {
	if num == 0 {
		return "0"
	}
	stack := ds.NewStack[int]()
	for num > 0 {
		stack.Push(num % 2)
		num /= 2
	}
	var binary strings.Builder
	binary.Grow(stack.Size()) // Preallocate memory for the builder
	for i, n := 0, stack.Size(); i < n; i++ {
		binary.WriteString(strconv.Itoa(stack.Pop().Value()))
	}
	return binary.String()
}

func (s *Stack) NextLargerElement(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	stack := ds.NewStack[int]()
	index := n - 1

	for index > -1 {
		num := arr[index]
		if stack.IsEmpty() {
			result[index] = -1
			stack.Push(num)
			index -= 1
		} else if stack.Peek().Value() > num {
			result[index] = stack.Peek().Value()
			stack.Push(num)
			index -= 1
		} else {
			stack.Pop()
		}
	}
	return result
}

func (s *Stack) SortStack(stack *ds.Stack[int]) *ds.Stack[int] {
	sortedStack := ds.NewStack[int]()
	for !stack.IsEmpty() {
		inputStackTopNum := stack.Pop().Value()
		if sortedStack.IsEmpty() {
			sortedStack.Push(inputStackTopNum)
		} else if inputStackTopNum >= sortedStack.Peek().Value() {
			sortedStack.Push(inputStackTopNum)
		} else {
			for !sortedStack.IsEmpty() && inputStackTopNum < sortedStack.Peek().Value() {
				stack.Push(sortedStack.Pop().Value())
			}
			sortedStack.Push(inputStackTopNum)
		}
	}
	return sortedStack
}

// simplifyPath simplifies the given file path
func (s *Stack) SimplifyPathByteByByte(path string) string {
	stack := ds.NewStack[byte]()
	n := len(path)

	for i := 0; i < n; i++ {
		char := path[i]
		isNextEnd := i+1 == n

		if stack.IsEmpty() {
			stack.Push(char)
		} else if char == '/' {
			if stack.Peek().Value() == '.' {
				stack.Pop()
			} else if stack.Peek().Value() != '/' {
				stack.Push(char)
			}
		} else if char == '.' {
			if stack.Peek().Value() == '.' {
				stack.Pop()
				if stack.Size() > 1 {
					stack.Pop()
					// we need to pop until we reach /, which brings us up a dir
					for !stack.IsEmpty() && stack.Peek().Value() != '/' {
						stack.Pop()
					}
				}
			} else if !isNextEnd {
				stack.Push(char)
			}
		} else {
			stack.Push(char)
		}
	}
	if stack.Size() > 1 && stack.Peek().Value() == '/' {
		stack.Pop()
	}
	// convert stack to string
	bytes := make([]byte, stack.Size())
	for i := len(bytes) - 1; i >= 0; i-- {
		bytes[i] = stack.Pop().Value()
	}
	return string(bytes)
}

func (s *Stack) SimplifyPath(path string) string {
	stack := []string{}
	for _, part := range strings.Split(path, "/") {
		if part != "" && part != "." {
			if part != ".." {
				stack = append(stack, part)
			} else if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}

		}
	}
	return "/" + strings.Join(stack, "/")
}

// removeDuplicates removes all duplicates from the string s
func (s *Stack) RemoveDuplicates(str string) string {
	stack := ds.NewStack[byte]()
	for i := len(str) - 1; i > -1; i-- {
		char := str[i]

		// check if stack is empty
		if stack.IsEmpty() {
			stack.Push(char)
		} else if stack.Peek().Value() == char {
			stack.Pop()
		} else {
			stack.Push(char)
		}
	}

	var bytes []byte = make([]byte, stack.Size())
	for i, n := 0, stack.Size(); i < n; i++ {
		bytes[i] = stack.Pop().Value()
	}
	return string(bytes)
}

func (s *Stack) RemoveDuplicatesOptimized(str string) string {
	stack := []byte{}
	for i, n := 0, len(str); i < n; i++ {
		char := str[i]
		if len(stack) == 0 {
			stack = append(stack, char)
			continue
		}
		// Peek and compare char
		if stack[len(stack)-1] == char {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}
	return string(stack)
}

func (s *Stack) RemoveStars(str string) string {
	stack := []byte{}
	for i := 0; i < len(str); i++ {
		char := str[i]
		if len(stack) > 0 && char == '*' {
			stack = stack[:len(stack)-1]
		} else if char != '*' {
			stack = append(stack, char)
		}
	}
	return string(stack)
}

func (s *Stack) MakeGood(str string) string {
	stack := []rune{}
	for _, char := range str {
		if len(stack) == 0 || char == stack[len(stack)-1] {
			stack = append(stack, char)
		} else if unicode.ToLower(char) == unicode.ToLower(stack[len(stack)-1]) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}
	return string(stack)
}

func (s *Stack) NextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := []int{}      // worst space o(n)
	hash := map[int]int{} // worst space o(n)

	// first, populate the hash where each number points to
	// a number that is either bigger than it or -1.
	// The strategy is to start from the end, and work to the front

	for i := len(nums2) - 1; i > -1; {
		num := nums2[i]
		// check if the stack is empty. if it is, set the num
		// in the hash to point to -1, then push the num
		if len(stack) == 0 {
			hash[num] = -1
			stack = append(stack, num)
			i -= 1
		} else if stack[len(stack)-1] > num {
			hash[num] = stack[len(stack)-1]
			stack = append(stack, num)
			i -= 1
		} else {
			// pop
			stack = stack[:len(stack)-1]
		}
	}
	// now we have our hash, we pass through nums1, updating
	// each cell with its value from the map
	for i := 0; i < len(nums1); i++ {
		nums1[i] = hash[nums1[i]]
	}

	return nums1
}

func (s *Stack) NextGreaterElements(nums []int) []int {
	n := len(nums)
	stack := []int{}
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = -1
	}

	for i := 2*n - 1; i > -1; i-- {
		index := i % n
		num := nums[index]

		for len(stack) > 0 && num >= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1] // pop
		}

		// if the stack is not empty at this point, the top is
		// is the index from nums represeting the next greater element
		if len(stack) > 0 {
			result[index] = nums[stack[len(stack)-1]]
		}

		stack = append(stack, index)
	}
	return result
}

func (s *Stack) NextGreaterElementsByValue(nums []int) []int {
	n := len(nums)
	stack := []int{}
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = -1
	}

	for i := 2*n - 1; i > -1; i-- {
		index := i % n
		num := nums[index]

		for len(stack) > 0 && num >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			result[index] = stack[len(stack)-1]
		}

		stack = append(stack, num)
	}
	return result
}
