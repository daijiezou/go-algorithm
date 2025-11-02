package main

import "testing"

func TestLRUCache_Basic(t *testing.T) {
	cache := New(2)

	// 测试基本的 Put 和 Get
	cache.Put(1, 1)
	cache.Put(2, 2)

	if val := cache.Get(1); val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}

	// 添加新元素，应该淘汰 key=2
	cache.Put(3, 3)

	if val := cache.Get(2); val != -1 {
		t.Errorf("Expected -1 (key 2 should be evicted), got %d", val)
	}

	if val := cache.Get(3); val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}
}

func TestLRUCache_Update(t *testing.T) {
	cache := New(2)

	// 测试更新已存在的 key
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(1, 10) // 更新 key=1 的值

	if val := cache.Get(1); val != 10 {
		t.Errorf("Expected 10, got %d", val)
	}

	// 添加新元素，应该淘汰 key=2（因为 key=1 刚被访问过）
	cache.Put(3, 3)

	if val := cache.Get(2); val != -1 {
		t.Errorf("Expected -1 (key 2 should be evicted), got %d", val)
	}
}

func TestLRUCache_GetUpdatesOrder(t *testing.T) {
	cache := New(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	// 访问 key=1，使其成为最近使用
	cache.Get(1)

	// 添加新元素，应该淘汰 key=2
	cache.Put(3, 3)

	if val := cache.Get(1); val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}

	if val := cache.Get(2); val != -1 {
		t.Errorf("Expected -1 (key 2 should be evicted), got %d", val)
	}

	if val := cache.Get(3); val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}
}

func TestLRUCache_SingleCapacity(t *testing.T) {
	cache := New(1)

	cache.Put(1, 1)
	if val := cache.Get(1); val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}

	cache.Put(2, 2)
	if val := cache.Get(1); val != -1 {
		t.Errorf("Expected -1 (key 1 should be evicted), got %d", val)
	}

	if val := cache.Get(2); val != 2 {
		t.Errorf("Expected 2, got %d", val)
	}
}

func TestLRUCache_ComplexScenario(t *testing.T) {
	cache := New(3)

	// 复杂场景测试
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4) // 淘汰 key=1

	if val := cache.Get(1); val != -1 {
		t.Errorf("Expected -1, got %d", val)
	}

	if val := cache.Get(2); val != 2 {
		t.Errorf("Expected 2, got %d", val)
	}

	cache.Put(5, 5) // 淘汰 key=3（因为 key=2 刚被访问）

	if val := cache.Get(3); val != -1 {
		t.Errorf("Expected -1, got %d", val)
	}

	if val := cache.Get(4); val != 4 {
		t.Errorf("Expected 4, got %d", val)
	}

	if val := cache.Get(5); val != 5 {
		t.Errorf("Expected 5, got %d", val)
	}
}

func TestLRUCache_MultipleUpdates(t *testing.T) {
	cache := New(2)

	cache.Put(1, 1)
	cache.Put(1, 10)
	cache.Put(1, 100)

	if val := cache.Get(1); val != 100 {
		t.Errorf("Expected 100, got %d", val)
	}

	// 验证 cache 中只有一个 key=1 的节点
	if len(cache.Cache) != 1 {
		t.Errorf("Expected cache size 1, got %d", len(cache.Cache))
	}
}

func TestLRUCache_GetNonExistent(t *testing.T) {
	cache := New(2)

	if val := cache.Get(1); val != -1 {
		t.Errorf("Expected -1 for non-existent key, got %d", val)
	}

	cache.Put(1, 1)

	if val := cache.Get(2); val != -1 {
		t.Errorf("Expected -1 for non-existent key, got %d", val)
	}
}
