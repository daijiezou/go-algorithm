package aim_offer

import (
	"fmt"
	"testing"
)

// 辅助函数：创建测试链表
func createTestList() *RandomListNode {
	// 创建节点
	node1 := &RandomListNode{Label: 1}
	node2 := &RandomListNode{Label: 2}
	node3 := &RandomListNode{Label: 3}
	node4 := &RandomListNode{Label: 4}

	// 设置Next指针
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	// 设置Random指针
	node1.Random = node3  // 1 -> 3
	node2.Random = node1  // 2 -> 1
	node3.Random = node4  // 3 -> 4
	node4.Random = node2  // 4 -> 2

	return node1
}

// 辅助函数：打印链表
func printList(head *RandomListNode, name string) {
	fmt.Printf("%s链表:\n", name)
	cur := head
	for cur != nil {
		randomLabel := "nil"
		if cur.Random != nil {
			randomLabel = fmt.Sprintf("%d", cur.Random.Label)
		}
		fmt.Printf("节点%d -> Next: ", cur.Label)
		if cur.Next != nil {
			fmt.Printf("%d", cur.Next.Label)
		} else {
			fmt.Printf("nil")
		}
		fmt.Printf(", Random: %s\n", randomLabel)
		cur = cur.Next
	}
	fmt.Println()
}

// 辅助函数：验证两个链表是否相同（值相同但地址不同）
func verifyClone(original, cloned *RandomListNode) bool {
	cur1, cur2 := original, cloned
	
	for cur1 != nil && cur2 != nil {
		// 检查值是否相同
		if cur1.Label != cur2.Label {
			fmt.Printf("值不匹配: 原始=%d, 克隆=%d\n", cur1.Label, cur2.Label)
			return false
		}
		
		// 检查地址是否不同（确保是深拷贝）
		if cur1 == cur2 {
			fmt.Printf("地址相同，不是深拷贝: %p\n", cur1)
			return false
		}
		
		// 检查Random指针
		if (cur1.Random == nil) != (cur2.Random == nil) {
			fmt.Println("Random指针状态不匹配")
			return false
		}
		
		if cur1.Random != nil && cur2.Random != nil {
			if cur1.Random.Label != cur2.Random.Label {
				fmt.Printf("Random指向的值不匹配: 原始=%d, 克隆=%d\n", 
					cur1.Random.Label, cur2.Random.Label)
				return false
			}
		}
		
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	
	// 检查长度是否相同
	if cur1 != nil || cur2 != nil {
		fmt.Println("链表长度不匹配")
		return false
	}
	
	return true
}

func TestClone(t *testing.T) {
	// 测试空链表
	if Clone(nil) != nil {
		t.Error("空链表测试失败")
	}
	
	// 测试复杂链表
	original := createTestList()
	cloned := Clone(original)
	
	fmt.Println("=== 复杂链表克隆测试 ===")
	printList(original, "原始")
	printList(cloned, "克隆")
	
	if !verifyClone(original, cloned) {
		t.Error("复杂链表克隆验证失败")
	} else {
		fmt.Println("✅ 克隆成功！所有测试通过")
	}
}

// 演示函数
func ExampleClone() {
	// 创建一个简单的测试链表
	node1 := &RandomListNode{Label: 7}
	node2 := &RandomListNode{Label: 13}
	node3 := &RandomListNode{Label: 11}
	
	node1.Next = node2
	node2.Next = node3
	node1.Random = node3  // 7 -> 11
	node2.Random = node1  // 13 -> 7
	node3.Random = node1  // 11 -> 7
	
	cloned := Clone(node1)
	
	fmt.Println("原始链表和克隆链表的地址比较:")
	fmt.Printf("原始node1地址: %p, 克隆node1地址: %p\n", node1, cloned)
	fmt.Printf("原始node2地址: %p, 克隆node2地址: %p\n", node2, cloned.Next)
	fmt.Printf("原始node3地址: %p, 克隆node3地址: %p\n", node3, cloned.Next.Next)
}
