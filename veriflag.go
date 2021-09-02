package main

import (
	"fmt"
	"os"
)

//go:generate generate.bat
func main() {
	args := os.Args
	fmt.Println(args)
	//veriflags works like flags but lets you restrict the range of values
	//based on adding a function (inpt)->bool to the flag definition
	//also it's going to return concrete types + booleans instead of pointers
	//if a default is set, i'm not sure why using a pointer return is beneficial..?

	//1. define the flags & how to get them out - thinking to return a function that can be called after parse to get the value
	// getprice := veriflag.IntFlag("x",0,range(0,10), "price of a freddo in pence")
	// veriflag.Parse()
	// price := getprice()
	// could even do price := veriflag.IntFlag("x",0,range(0,10), "price of a freddo in pence")()
	getprice := IntFlag("x", 0, func(x int64) bool { return x > 0 && x < 10 }, "price of a freddo")
	fmt.Println(getprice())
	fmt.Println(IntFlag("freddos", 0, InListint64([]int64{1, 6, 10}), "number of freddos, only comes in certain pack sizes")())
	fmt.Println(IntFlag("fav", 0, InListint64([]int64{1, 6, 10}), "number of freddos, only comes in certain pack sizes")())
	fmt.Println(StringArg(1, "day", InListstring([]string{"MON", "TUES"}), "monday or tuesday")())

	//super cool feature would be to load in json/toml config but things like nested lists/maps will be hard to handle
	//For basic args however could be doable and handy
	// getprice := veriflag.IntConfig("x",0,range(0,10),"price of a freddo in pence", tomlconfigstr)
	// currencies := veriflag.MapConfig("crncy",default, keycond,valcond,descr, tomlconfigstr, map[string]mystruct)

	// fmt.Print(GenerateInList("int64"))
}
