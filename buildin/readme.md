
源码
buildin.go 里面
```
func make(t Type, size ...IntegerType) Type
```
# make
// The make built-in function allocates and initializes an object of type
// slice, map, or chan (only).
用于初始化 slice map chan 这三种类型的对象的初始化的一个函数。
// Like new, the first argument is a type, not a value.
和 new 函数 一样的是，第一个方法参数是一个类型而不是一个值。
// Unlike new, make's return type is the same as the type of its argument, not a pointer to it.
不一样的是，返回的类型和参数类型是一样的（集合、队列里面的类型要一致？）
// The specification of the result depends on the type:
返回结果的特性取决于类型，按类型来区分如下：
1. Slice: 切片，
//	Slice: The size specifies the length. The capacity of the slice is
//	equal to its length. A second integer argument may be provided to
//	specify a different capacity; it must be no smaller than the
//	length. For example, make([]int, 0, 10) allocates an underlying array
//	of size 10 and returns a slice of length 0 and capacity 10 that is
//	backed by this underlying array.
//	Map: An empty map is allocated with enough space to hold the
//	specified number of elements. The size may be omitted, in which case
//	a small starting size is allocated.
//	Channel: The channel's buffer is initialized with the specified
//	buffer capacity. If zero, or the size is omitted, the channel is
//	unbuffered.

# slice
类似 java 的 arrayList.
区别：添加元素超过初始容量时，会把之前的复制原来的到一个新的数组。

为什么这边设计呢？为什么和 java 的 arrayList 的实现逻辑不一样？是有什么优点？

1. 空间换时间：不影响之前的切片的使用。就像我和你一起买了一个储物间存放东西，可以存放10个，我有需求变多了。要100个。
但是你不想要这么多，那就分道扬镳吧。我走我的阳关道，你走你的独木桥.


