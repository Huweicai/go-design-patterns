package creational

type obj struct {
	name string
}

func (o *obj) Run() {
	println(o.name)
}

type pool chan *obj

func NewPool(size int) *pool {
	p := make(pool, size)
	for i := 0; i < size; i++ {
		p <- new(obj)
	}
	return &p
}

func (p pool) Get() *obj {
	return <-p
}

func (p pool) Return(o *obj) {
	p <- o
}
