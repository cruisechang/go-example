package interfaceMiddleware

//interface middleware example

type middleware interface {
	execute(string) string
	backExecute(string) string
}

type controller struct {
	ware []middleware
}

func NewController()*controller{
	return & controller{
		ware:[]middleware{},
	}
}
func (c *controller) Use(ware middleware) {
	c.ware = append(c.ware, ware)
}

func (c *controller) Run(args string) string {

	llen := len(c.ware)
	if llen == 0 {
		return args
	}

	//index 0
	ware := c.ware[0]
	res := ware.execute(args)

	if llen > 1 {

		//begin from index 1
		for i := 1; i < llen; i++ {
			ware = c.ware[i]
			res = ware.execute(res)
		}

		//begin from 倒數第二個
		for i := llen - 2; i > -1; i-- {
			ware = c.ware[i]
			res = ware.backExecute(res)
		}
	}

	return res

}
