package generator

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	ch := Count(1, 50)

	for i := range ch {
		fmt.Println(i)
	}
}
