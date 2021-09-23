package funcMiddleware

//func middleware example

type middleware func(con *controller)

type controller struct {
	wares []middleware
	wareIndex int
	//container for middleware
}

func NewController()*controller{
	return & controller{
		wares: []middleware{},
		wareIndex: 0,
	}
}
func (c *controller) Use(mw middleware) {
	c.wares=append(c.wares,mw)
}
func(c *controller)Run(){
	c.wareIndex=0
	if len(c.wares)>0{
		c.wares[c.wareIndex](c)
	}
}
func (c *controller) Next() {
	if c.wareIndex<(len(c.wares)-1){
		c.wareIndex++
		c.wares[c.wareIndex](c)
	}
}
