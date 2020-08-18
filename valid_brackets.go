package main

import "fmt"

type Bracket struct {
	bracket byte
	next *Bracket
}

type BracketStack struct {
	bracket *Bracket
}

func (receiver *BracketStack) push(bracket byte) {
	if receiver.bracket==nil {
		receiver.bracket = &Bracket{bracket: bracket}
		return
	}
	temp := &Bracket{bracket: bracket, next: receiver.bracket}
	receiver.bracket = temp
	return
}

func (receiver *BracketStack) pop() byte{
	if receiver.bracket==nil{
		return '!'
	}
	temp := receiver.bracket
	receiver.bracket = temp.next
	return temp.bracket
}

func isValid(s string) bool {
	bracketStack := &BracketStack{}
	for i:=0;i<len(s);i++ {
		switch s[i] {
		case '(':
			bracketStack.push(byte(s[i]))
		case '[':
			bracketStack.push(byte(s[i]))
		case '{':
			bracketStack.push(byte(s[i]))
		case ')':
			ch := bracketStack.pop()
			if ch!='(' {
				return false
			}
		case ']':
			ch := bracketStack.pop()
			if ch!='[' {
				return false
			}
		case '}':
			ch := bracketStack.pop()
			if ch!='{' {
				return false
			}
		}
	}
	if bracketStack.pop()=='!' {
		return true
	}
	return false
}

func main() {
	bracketStr := "{[]}"
	fmt.Println("Validating the String: ",isValid(bracketStr))
}
