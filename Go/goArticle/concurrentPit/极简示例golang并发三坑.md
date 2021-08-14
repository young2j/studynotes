#! https://zhuanlan.zhihu.com/p/399905251
我们经常会遇到一些耗时的任务,然后又需要拿到任务处理后的结果作进一步处理,  在go语言中首先想到的莫过于`goroutine`加等待组`wg`的方式来并发处理加快效率。尽管在go中写并发程序已经足够简单了，但对一部分人来说往往一不注意就会掉进坑里。本文通过一个简单的例子梳理几个踩坑点。

# 简单的例子

业务场景中，我们往往需要将多个任务或者一个任务拆分，然后分别作复杂的逻辑处理。假如我们有1到10个数字代表了这10个任务，然后需要分别对这10个数字乘以2，以此代表逻辑处理。

# 坑点一

猜猜下面的程序会输出什么？

```go
var (
	wg   = sync.WaitGroup{}
	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
)

func task(num int) int {
	return num * 2
}

func main() {
	wg.Add(len(nums))
	for _, num := range nums {
		go func() {
			defer wg.Done()
			res := task(num)
			fmt.Println(res)
		}()
	}
	wg.Wait()
}
```

如果你的结论是"以不确定的顺序输出`2,4,6,8,10,12,14,16,18,20`"，那么恭喜你入坑了！运行代码的实际结果为:

```shell
20
20
20
20
20
20
20
20
20
20
```

因为`for循环中的goroutine在实际运行的时候，循环已经执行完毕了，num的值为循环后的最后一个值20`。解决这个问题也很简单，在很多语言中也是如此，通过闭包的方式让每一个`go func()`独自保存其`num`值。

> 这种最基本的坑，虽然可以运行，但编辑器往往会有提示的

修正代码如下，便可以不确定的顺序输出`2,4,6,8,10,12,14,16,18,20`:

```go
var (
	wg   = sync.WaitGroup{}
	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
)

func task(num int) int {
	return num * 2
}

func main() {
	wg.Add(len(nums))
	for _, num := range nums {
		go func(num int) {
			defer wg.Done()
			res := task(num)
			fmt.Println(res)
		}(num)
	}
	wg.Wait()
}
```

# 坑点二

在上面的例子中，有些有着严格内存管理要求的小伙伴，可能会不假思索的改成这样:

```go
var (
	wg   = sync.WaitGroup{}
	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
)

func task(num int) int {
	return num * 2
}

func main() {
	wg.Add(len(nums))
	for _, num := range nums {
		go func(num *int) {
			defer wg.Done()
			res := task(*num)
			fmt.Println(res)
		}(&num)
	}
	wg.Wait()
}
```

这又会输出什么结果呢？不一定，但大部分值都为20:

```shell
20
20
20
20
20
8
20
20
20
20
```

同样是闭包，为什么传指针就不行了呢？恰恰是因为闭包，`go func()里保存了同一个内存地址`，即`&num`在`for循环`中指向的是同一个内存地址，但该地址上存储的值在for中不断发生变化，`goroutine`实际执行时，`&num`上的值基本都已经变成了最后一个值20.所以这时候不能传递指针。

> 可能有些人会说，是不是傻，一个简单的int类型，没事去传个指针干啥？这都能入坑。在实际情况中，可能确实没有人会这么做，怪就怪在本文的例子太过简单，如若这里的num不是int类型呢？现实写代码的时候，这里往往可能是一个复杂的结构体，比如orm中定义的model，我想肯定会有人传递&model的！

# 坑点三

紧接着上面的例子，假如我们想将处理后的结果保存起来，很自然的写出了如下代码:

```go
var (
	wg   = sync.WaitGroup{}
	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	results  []int
)

func task(num int) int {
	return num * 2
}

func main() {
	wg.Add(len(nums))
	for _, num := range nums {
		go func(num int) {
			defer wg.Done()
			res := task(num)
			results = append(results, res)
		}(num)
	}
	wg.Wait()
	fmt.Println(results)
}
```

`results`又会输出什么结果呢？如果你的结论是`results中包含顺序不定的2,4,6,8,10,12,14,16,18,20`，那么恭喜你又入坑了！正常情况下，`len(results)`的值应该为10，但上面代码多运行几次的结果表明，`len(results)`的值几乎都是小于10的。因为在go中，`切片slice`类型是非并发安全的，也就是说`results`中的某一个位置在同一时刻插入了多个值，最终造成了数据丢失。解决的办法可以通过加锁的方式:

```go
var (
	wg   = sync.WaitGroup{}
	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	results  []int
	lock  = sync.RWMutex{}
)

func task(num int) int {
	return num * 2
}

func main() {
	wg.Add(len(nums))
	for _, num := range nums {
		go func(num int) {
			defer wg.Done()
			res := task(num)
			lock.Lock()
			results = append(results, res)
			lock.Unlock()
		}(num)
	}
	wg.Wait()
	fmt.Println(results)
}
```

> 类似`slice`类型，go中`map`类型也是非并发安全的，在并发场景中我们可以使用sync.Map代替。

# 小结

就像那单细胞生物，越是简单反而越让人头疼！上面三个踩坑点其实都非常简单，大部分的人可能都是跟我一样的心态："这么简单的问题我是不会入坑的！"。然而，自认为go代码我已经写得很熟练了，却不曾想最近在业务代码中被疯狂打脸，浪费很多时间！可能就关注代码本身而言，我想很少有人会掉进坑里。道理都懂，但现实中当我们的思路总是关注于复杂的业务逻辑如何组织代码实现时，一长串一长串的代码往往会让我们疏于这些细节。谨以此文为诫。

