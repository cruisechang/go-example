package interfaceMiddleware

import "fmt"

func doExample() {
	con:=NewController()
	con.Use(&sample1{})
	con.Use(&sample2{})
	con.Use(&sample3{})
	con.Run("controller:")
}

type sample1 struct{}
func (s1 *sample1)execute(param string)string{
	res:=param+" execute1"
	fmt.Println(res)
	return res
}
func (s1 *sample1)backExecute(param string)string{
	res:=param+" backExecute1"
	fmt.Println(res)
	return res
}


type sample2 struct{}
func (s1 *sample2)execute(param string)string{
	res:=param+" execute2"
	fmt.Println(res)
	return res
}
func (s1 *sample2)backExecute(param string)string{
	res:=param+" backExecute2"
	fmt.Println(res)
	return res
}

type sample3 struct{}
func (s1 *sample3)execute(param string)string{
	res:=param+" execute3"
	fmt.Println(res)
	return res
}
func (s1 *sample3)backExecute(param string)string{
	res:=param+" backExecute3"
	fmt.Println(res)
	return res
}


