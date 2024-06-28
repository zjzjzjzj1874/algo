package middle

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	t.Run("#Test", func(t *testing.T) {
		lru := Constructor(2)
		var res int
		lru.Put(1, 0)
		lru.Put(2, 2)
		res = lru.Get(1)
		lru.Put(3, 3)
		res = lru.Get(2) // 返回 -1 (未找到)
		lru.Put(4, 4)    // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}

		res = lru.Get(1) // 返回 -1 (未找到)
		res = lru.Get(3) // 返回 3
		res = lru.Get(4) // 返回 4

		fmt.Println(res)
	})
}
