package chain

import (
	"fmt"
	"strings"
)

// 责任链模式
type ChainLogger interface {
	Next(string)
}

type FirstLogger struct {
	NextChain ChainLogger
}

func (f *FirstLogger) Next(s string) {
	fmt.Printf("First logger: %s\n", s)
	if f.NextChain != nil {
		f.NextChain.Next(s)
	}
}

type SecondLogger struct {
	NextChain ChainLogger
}

func (s *SecondLogger) Next(str string) {
	if strings.Contains(strings.ToLower(str), "hello") {
		fmt.Printf("Second logger: %s\n", str)
		if s.NextChain != nil {
			s.NextChain.Next(str)
		}
	}
	fmt.Printf("Finishing in second logging \n\n")
}
