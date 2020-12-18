## context

- Deadline() (deadline time.Time, ok bool)
  > 返回设置的deadline，没有设置ok为false
- Done() <-chan struct{}
  > 返回一个channel，如果已经取消了返回关闭的，如果不能被取消则返回nil
  >> WithCancel类型的调用cancel后关闭
  >> WithDeadline类型的deadline到期后关闭
  >> WithTimeout类型的timeout时间过了关闭
- Err() error
  > 当Done的channel没有关闭时返回nil
  > Done的channel关闭后返回error解释为什么关闭
  > var Canceled = errors.New("context canceled") 取消后返回的 error
  > var DeadlineExceeded error = deadlineExceededError{} deadline超过后返回的错误，它有一个方法Temporary 返回true
- Value(key interface{}) interface{}
  > 返回对应key设置的值，如果没有设置返回nil
  > key的值一般设置在全局变量中


### type emptyCtx int
  > 根context，不能被取消，其实就是什么都没有 这个类型的有两个
  > var (
  >  	background = new(emptyCtx)   //Background() 返回这个
  >  	todo       = new(emptyCtx)   //TODO() 返回这个
  >  ) //这两个在技术上一样，只是语义上不同

### type CancelFunc func()
  > 取消的方法，调用过后再调用什么都不干

### func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
  > 返回一个带子context，cancel可以终止这个context

