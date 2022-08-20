package actor

import (
	"fmt"
	"testing"
)

func TestUint64ToId(t *testing.T) {
	for i := uint64(100000000); i < 100000100; i++ {
		fmt.Println(uint64ToId(i))
	}
	fmt.Println("=======")
	for i := uint64(0); i < 10; i++ {
		fmt.Println(uint64ToId(i))
	}
}
