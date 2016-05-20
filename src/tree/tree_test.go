package tree

import "testing"

func Test_GetDepth(t *testing.T) {
	a := GetDepth(RootNode)
	if a != 4 {
		t.Errorf("get depth failed")
	}
}

func Test_GetNodeNum(t *testing.T) {
	a := GetNodeNum(RootNode)
	if a != 6 {
		t.Errorf("get total count failed")
	}

}
