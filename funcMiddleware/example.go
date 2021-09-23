package funcMiddleware

import "fmt"

//
func doExample() {
	con:=NewController()
	con.Use(sampleFunc1)
	con.Use(sampleFunc2)
	con.Use(sampleFunc3)
	con.Use(sampleFunc4("param"))
	con.Run()
}

func sampleFunc1(con *controller){
	fmt.Println("before1")
	con.Next()
	fmt.Println("after1")
}
func sampleFunc2(con *controller){
	fmt.Println("before2")
	con.Next()
	fmt.Println("after2")
}
func sampleFunc3(con *controller){
	fmt.Println("before3")
	con.Next()
	fmt.Println("after3")
}
func sampleFunc4(param string)middleware{
	return func(con *controller){
		fmt.Println(param+":before4")
		con.Next()
		fmt.Println(param+":after4")
	}
}