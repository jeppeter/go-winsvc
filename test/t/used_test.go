package used

import "testing"

func TestAdd(t *testing.T) {
	c := Add(1, 2)
	if c != 3 {
		t.Error("not right Add")
	}

}

func TestSub(t *testing.T) {
	c := Sub(10, 5)
	if c != 5 {
		t.Error("not right Sub")
	}
}

func TestAddFailed(t *testing.T) {

	c := Add(10, 20)
	if c != 10 {
		t.Errorf("Add(10,20) = %d", c)
	}
}

func TestSubFailed(t *testing.T) {
	c := Sub(100, 20)
	if c != 20 {
		t.Errorf("Sub(100,20) = %d", c)
	}
}
