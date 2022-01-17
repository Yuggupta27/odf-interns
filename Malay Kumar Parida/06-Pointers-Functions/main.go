package main

import "fmt"

func main() {
	var a int = 42
	var b *int
	fmt.Println(a, b)
	b = &a
	fmt.Println(a, b)
	fmt.Println(&a, *b)
	a = 27
	fmt.Println(a, *b)
	*b = 14
	fmt.Println(a, *b)

	x := [3]int{1, 2, 3}
	y := &x[0]
	z := &x[1]
	fmt.Printf("%v %p %p\n", x, y, z)

	//Pointer arithmetic directly is not allowed in golang

	var ms *myStruct
	fmt.Println(ms)         //Pointers are initialized to nil
	ms = &myStruct{foo: 42} //pointer to object
	ms2 := new(myStruct)    //Using new function

	fmt.Println(ms, ms2)
	(*ms).foo = 100
	fmt.Println(ms.foo) //Can be used without dereferencing

	//Functions
	//Every go program starts with main function

	greetSomeone("Hello", "Malay")
	//We can have any number of parameters
	grretSomeone2("Hello", "Malay", "How are you?")

	//Pass by reference
	msg := "Hello"
	name := "Malay"
	greetSomeone3(&msg, name)
	fmt.Println(msg, name)

	s := sum(1, 2, 3, 4, 5) //function with return
	fmt.Println(s)

	s2 := sum2(1, 2)
	fmt.Println(*s2)

	d, err := divide(5.0, 0.0) //returning multiple values and error values
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	//Annonymous function
	func(){
		fmt.Println("Hello Golang")
	}()//Immediaely invoked


	//Functions as variables
	f:=func(){
		fmt.Println("Hello Golang")
	}
	f()

	//Methods
	g:=greeter{
		grettings: "Hello",
		name:"Go",
	}
	g.greet()
	
}

type greeter struct{
		grettings string
		name string
	}

	func(g greeter) greet(){
		fmt.Println(g.grettings,g.name)
	}

func divide(a, b float64) (float64, error) { //returning multpile values includinf error
	if b == 0.0 {
		return 0.0, fmt.Errorf("can Not devide by zero")
	}
	return a / b, nil
}

func sum2(val1, val2 int) *int {
	result := val1 + val2
	return &result //In go we can return pointer to local function variable
}

func sum(values ...int) int { //Catch parameters as a list
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result //return something
}

func greetSomeone(msg string, name string) {
	fmt.Println(msg, name)
}

//For mutiple same parameters only once type is sufficient
func grretSomeone2(msg, name, question string) {
	fmt.Println(msg, name, question)
}

//Pass By reference
func greetSomeone3(msg *string, name string) {
	*msg = "Holla"
	fmt.Println(*msg, name)
}

type myStruct struct {
	foo int
}
