Go语言可以对任何类型添加方法，给一种类型添加方法就像给结构体添加方法一样，因为结构体也是一种类型。

### 为基本类型添加方法

在Go语言中，使用 type 关键字可以定义出新的自定义类型，之后就可以为自定义类型添加各种方法了。我们习惯于使用面向过程的方式判断一个值是否为 0，例如：

if  v == 0 {
    // v等于0
}

如果将 v 当做整型对象，那么判断 v 值就可以增加一个 IsZero() 方法，通过这个方法就可以判断 v 值是否为 0，例如：

if  v.IsZero() {
    // v等于0
}

为基本类型添加方法的详细实现流程如下：
+ 案例
  1.go