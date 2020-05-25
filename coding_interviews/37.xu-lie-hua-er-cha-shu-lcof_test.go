package codinginterviews

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCodec1(t *testing.T) {
	c := Constructor()
	root := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
	}
	serialize := c.serialize(root)
	fmt.Println(serialize)
	deserialize := c.deserialize(serialize)
	if !reflect.DeepEqual(root, deserialize) {
		fmt.Println("deserialize failed")
		return
	}
	fmt.Println("deserialize success!")
}

// func TestCodec2(t *testing.T) {
// 	c:=Constructor()
// 	root:=&TreeNode{
// 		Val: ,
// 	}
// }
