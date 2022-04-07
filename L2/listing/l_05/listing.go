package main

type Printer interface {
	Print()
}

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	//fmt.Printf("%T %v\n", err, err)
	err = test()
	//fmt.Printf("%T %v\n", err, err)
	if err != nil {
		println("error")
		return
	}

	println("ok")
}
