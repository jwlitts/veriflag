# veriflag

veriflag adds verification to command line named and positional arguments

Basic Usage:

```

import veriflag

func main(){

    getprice := IntFlag("x", 0, func(x int64) bool { return x > 0 && x < 10 }, "price of a freddo")
	fmt.Println(getprice())
	fmt.Println(IntFlag("freddos", 0, InListint64([]int64{1, 6, 10}), "number of freddos, only comes in certain pack sizes")())
	fmt.Println(IntFlag("fav", 0, InListint64([]int64{1, 6, 10}), "number of freddos, only comes in certain pack sizes")())
	fmt.Println(StringArg(1, "day", InListstring([]string{"MON", "TUES"}), "monday or tuesday")())
}

```
