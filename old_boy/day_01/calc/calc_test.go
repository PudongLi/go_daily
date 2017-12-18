package calc

import (
	"testing"
)

func TestAdd(t *testing.T)  {
	sum := Add(5, 6)
	if sum != 11{
	t.Fatalf("error,sum:%v expected:11", sum)
	}

}
func TestSub(t *testing.T) {
	diff := Sub(10, 9)
	if diff != 1{
		t.Fatalf("error,diff:%v expected:1", diff)
	}
}

