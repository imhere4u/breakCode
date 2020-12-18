###timer 使用

-  定时器正常触发  d Duration后收到触发
```go
package main

import (
    "fmt"
    "time"
)

func main() {

	t := time.NewTimer(3*time.Second)

	fmt.Println(time.Now())         //2020-12-18 15:22:17
	select {
	case i := <- t.C :
		fmt.Println(time.Now())     //2020-12-18 15:22:20
		fmt.Println(i)              //2020-12-18 15:22:20
	}

}
```

- 关闭定时器
```go
package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.NewTimer(3*time.Second)

    //这时未到触发，t.C是空的
	if !t.Stop() {        //t.Stop return true
		fmt.Println(time.Now())
		o := <- t.C
		fmt.Println(o)
	}
}
```
----定时器并发的使用会有一个被阻塞，甚至后面reset得不到正确触发
```go
package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.NewTimer(3*time.Second)

	go s(t)         //两个中一个sub end 一个被阻塞
	go s(t)

	time.Sleep(2*time.Second)
	fmt.Println("main end")
}
func s(t *time.Timer) {
	if !t.Stop() {
		fmt.Println(time.Now())
		o := <- t.C
		fmt.Println(o)
	}
	fmt.Println("sub end")
}
```

- reset注意清空channel
```go
package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.NewTimer(3*time.Second)

	time.Sleep(5*time.Second)

	/*
	if !t.Stop() {
		fmt.Println(time.Now())
		o := <- t.C
		fmt.Println(o)
	}
	 */

	fmt.Println(time.Now())  //2020-12-18 15:47:45
	t.Reset(2*time.Second)
	select {
	case p := <- t.C :
		fmt.Println(time.Now()) //2020-12-18 15:47:45
        //这是p并不是预期的reset后2秒，而是第一次延时的结果，就是因为没有清空channel，原来的在channel中。
        //将上面注释的代码放开就可以得到正确的结果
		fmt.Println(p)          //2020-12-18 15:47:43
	}

	fmt.Println("main end")
}

```