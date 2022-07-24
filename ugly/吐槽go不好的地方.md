

# 函数
## 函数的定义方式有多种，多就是复杂
通用语法:
func functionName([parametername parametertype]) [returntype]{
    // 函数体，具体的实现
}

具体可以使多个参数及返回值。
1. 最简单的，无参数及返回值
```go
func add(){
	
}
```
2. 无参数有返回值
```go
func add() age int{
	retuen 18;
}
```
3. 有参数无返回值
```go
func add(age int){
	age++
}
```

 但是看了很多github 上的代码，如下：
 func (c Cricle) accept(v Visitor) {
     v(c)
 }

就很懵了，怎么函数名前面还有()?
原来是有说法的：函数名称前面的括号是Go定义这些函数将在其上运行的对象的方式。来了一个新词“方式”，看完更加不明白，虽然字都认识。
ref:https://www.h5w3.com/183447.html.
1. 一个接收者，说明该方法属于哪个结构体. 类似 Java 类的方法，Go 没有类有结构体，即“接收者” --？“所有者” 是不是更精确更好理解呢


 
