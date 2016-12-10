package gomill

type Context struct {
	done chan struct{}
}

func NewContext() *Context {
	ctx := new(Context)
	ctx.done = make(chan struct{})
	return ctx
}

func (c *Context) Stop() {
	close(c.done)
}

func (c *Context) GetDone() <-chan struct{} {
	return c.done
}
