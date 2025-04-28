# Go 语言标准库 `container/list` 包中 `List` 和 `Element` 结构体的常用方法及函数的详细列举和说明

### **一、全局函数**
#### 1. `list.New() *List`
- **作用**：创建并初始化一个新的空链表。
- **参数**：无。
- **返回值**：指向新链表的指针（`*List`）。
- **示例**：
  ```go
  l := list.New()
  ```

---

### **二、List 结构体的方法**
#### 1. `Init() *List`
- **作用**：初始化或重置链表（清空链表并重置为初始状态）。
- **参数**：无。
- **返回值**：指向当前链表的指针（`*List`）。
- **示例**：
  ```go
  l.Init() // 清空链表并重置
  ```

#### 2. `Len() int`
- **作用**：返回链表中元素的数量。
- **参数**：无。
- **返回值**：链表的长度（`int`）。
- **时间复杂度**：O(1)。
- **示例**：
  ```go
  length := l.Len()
  ```

#### 3. `Front() *Element`
- **作用**：返回链表的第一个元素（头元素）。
- **参数**：无。
- **返回值**：指向第一个元素的指针（`*Element`），若链表为空则返回 `nil`。
- **示例**：
  ```go
  first := l.Front()
  ```

#### 4. `Back() *Element`
- **作用**：返回链表的最后一个元素（尾元素）。
- **参数**：无。
- **返回值**：指向最后一个元素的指针（`*Element`），若链表为空则返回 `nil`。
- **示例**：
  ```go
  last := l.Back()
  ```

#### 5. `PushFront(v interface{}) *Element`
- **作用**：在链表头部插入一个新元素。
- **参数**：`v`（要插入的值，类型为 `interface{}`）。
- **返回值**：指向新插入元素的指针（`*Element`）。
- **示例**：
  ```go
  e := l.PushFront(10)
  ```

#### 6. `PushBack(v interface{}) *Element`
- **作用**：在链表尾部插入一个新元素。
- **参数**：`v`（要插入的值，类型为 `interface{}`）。
- **返回值**：指向新插入元素的指针（`*Element`）。
- **示例**：
  ```go
  e := l.PushBack("hello")
  ```

#### 7. `InsertAfter(v interface{}, mark *Element) *Element`
- **作用**：在指定元素 `mark` 之后插入新元素。
- **参数**：
  - `v`：要插入的值。
  - `mark`：标记元素（新元素将插入到其后）。
- **返回值**：指向新插入元素的指针（`*Element`）。
- **示例**：
  ```go
  markElement := l.Front()
  e := l.InsertAfter(20, markElement)
  ```

#### 8. `InsertBefore(v interface{}, mark *Element) *Element`
- **作用**：在指定元素 `mark` 之前插入新元素。
- **参数**：
  - `v`：要插入的值。
  - `mark`：标记元素（新元素将插入到其前）。
- **返回值**：指向新插入元素的指针（`*Element`）。
- **示例**：
  ```go
  markElement := l.Back()
  e := l.InsertBefore(30, markElement)
  ```

#### 9. `Remove(e *Element) interface{`
- **作用**：删除指定元素 `e`，并返回其存储的值。
- **参数**：`e`（要删除的元素指针）。
- **返回值**：被删除元素的值（类型为 `interface{}`）。
- **注意**：删除后，`e` 不再属于链表，后续操作需谨慎。
- **示例**：
  ```go
  value := l.Remove(e)
  ```

#### 10. `MoveToFront(e *Element)`
- **作用**：将元素 `e` 移动到链表头部。
- **参数**：`e`（要移动的元素指针）。
- **返回值**：无（`void`）。
- **示例**：
  ```go
  l.MoveToFront(e)
  ```

#### 11. `MoveToBack(e *Element)`
- **作用**：将元素 `e` 移动到链表尾部。
- **参数**：`e`（要移动的元素指针）。
- **返回值**：无（`void`）。
- **示例**：
  ```go
  l.MoveToBack(e)
  ```

#### 12. `MoveAfter(e, mark *Element)`
- **作用**：将元素 `e` 移动到元素 `mark` 之后。
- **参数**：
  - `e`：要移动的元素。
  - `mark`：目标位置的标记元素。
- **返回值**：无（`void`）。
- **示例**：
  ```go
  l.MoveAfter(e, markElement)
  ```

#### 13. `MoveBefore(e, mark *Element)`
- **作用**：将元素 `e` 移动到元素 `mark` 之前。
- **参数**：
  - `e`：要移动的元素。
  - `mark`：目标位置的标记元素。
- **返回值**：无（`void`）。
- **示例**：
  ```go
  l.MoveBefore(e, markElement)
  ```

---

### **三、Element 结构体的方法**
#### 1. `Next() *Element`
- **作用**：返回当前元素的下一个元素。
- **参数**：无。
- **返回值**：指向下一个元素的指针（`*Element`），若当前是最后一个元素则返回 `nil`。
- **示例**：
  ```go
  nextElement := e.Next()
  ```

#### 2. `Prev() *Element`
- **作用**：返回当前元素的上一个元素。
- **参数**：无。
- **返回值**：指向上一个元素的指针（`*Element`），若当前是第一个元素则返回 `nil`。
- **示例**：
  ```go
  prevElement := e.Prev()
  ```

#### 3. `Value`
- **作用**：获取或设置当前元素的值。
- **类型**：`interface{}`。
- **示例**：
  ```go
  value := e.Value       // 获取值
  e.Value = "new value"  // 设置值
  ```

---

### **四、其他说明**
1. **链表特性**：
   - `container/list` 实现的是**双向循环链表**，每个元素（`Element`）包含 `prev` 和 `next` 指针，指向前后元素。
   - 链表操作（插入、删除）的时间复杂度均为 **O(1)**（若已知目标位置）。

2. **注意事项**：
   - 删除元素后，`Element` 指针不再有效，后续操作需通过链表的 `Front()`、`Back()` 或遍历获取新指针。
   - 遍历链表时，通常从 `Front()` 开始，通过 `Next()` 逐个访问元素：
     ```go
     for e := l.Front(); e != nil; e = e.Next() {
         fmt.Println(e.Value)
     }
     ```

---

### **五、完整示例**
```go
package main

import (
    "container/list"
    "fmt"
)

func main() {
    l := list.New()

    // 插入元素
    e1 := l.PushFront(100)
    e2 := l.PushBack("hello")
    l.InsertAfter(200, e1) // 在 e1 后插入 200

    // 遍历链表
    for e := l.Front(); e != nil; e = e.Next() {
        fmt.Printf("Value: %v\n", e.Value)
    }

    // 删除元素
    value := l.Remove(e2)
    fmt.Printf("Removed value: %v\n", value)

    // 移动元素
    l.MoveToFront(e1)
    fmt.Println("After moving e1 to front:")
    for e := l.Front(); e != nil; e = e.Next() {
        fmt.Printf("Value: %v\n", e.Value)
    }
}
```

---

### **六、常见用途**
- **栈**：通过 `PushBack` 和 `Remove(Back())` 实现。
- **队列**：通过 `PushBack` 和 `Remove(Front())` 实现。
- **缓存**：根据元素的访问顺序维护链表，淘汰最久未使用的元素。
