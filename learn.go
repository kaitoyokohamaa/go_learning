package main
import (
	"fmt"
	// "reflect"
)


// func sayHello() {
// 	fmt.Println("Hello World!")
// }

// func cal(x, y int) int{
// 	return (x + y)
// }

// func cal(x, y int) (int, int){
// 	return (x + y)
// }

// type Student struct {
// 	name string
// }

// type User struct {
// 	gender string
// 	age int
// }

// type User struct {
// 	name string
// }

// func (u User) cal(weight,height float64) (myBMI float64){
// 	myBMI =((weight / height / height )* 10000)
// 	return
// }

// func (s Student) avg(math, english float64) (avgResult float64) {
// 	avgResult = ((math + english) / 2)
// 	return
// }

type Student struct {
	name string
}

func (s Student) calAvg(data []float64) (avgResult float64) {
	sum := 0.0
	for i :=0; i< len(data); i++ {
		sum += data[i]
	}
	avgResult = sum /float64(len(data))
	return avgResult
}

func (s Student) judge(avg float64) (judgeResult string) {
	if avg >= 60 {
		judgeResult = "passed"
	} else {
		judgeResult = "failed"
	}
	return
}

func main(){
	a001 := Student{"kaito"}
	data := []float64{70, 65, 50, 90, 30}
	var avg float64 = a001.calAvg(data)
	result := a001.judge(avg)
	fmt.Println(avg)
	fmt.Println(a001.name + " " + result)
	// u := User{"kaito"}
	// fmt.Println(u.cal(52,160))
	// a0001 := Student{"sato"}
	// fmt.Println(a0001.avg(80, 70))
	// s := Student{"sato", 80, 70}
	// g :=User{"male", 20}
	// fmt.Println(g)
	// var s Student
	// s.name = "sato"
	// s.math = 80
	// s.english = 70

	// fmt.Println(s)
	// var string_a string="Helo, World!"
	// num01 :=123
	// var num02 int =1234567890
	// num03:=1.23
	// var num04 float64=1.23456789
	// a:=10
	// b:=1
	// var num_bool bool = a > b
	// fmt.Println(num_bool)
	// fmt.Println(reflect.TypeOf(num_bool))
	// a := [...]string{"sato", "suzuki", "takahashi"}
	// a[1] = "tanaka"
	// fmt.Println(a[0])
	// fmt.Println(a[1])
	// fmt.Println(a[2])
	// a := [2][2]string{{"sato", "suzuki"}, {"takahashi", "tanaka"}}
	// fmt.Println(a[0][0])
	// fmt.Println(a[0][1])
	// fmt.Println(a[1][0])
	// fmt.Println(a[1][1])

	// x := 10
	// y := 2

	// fmt.Println(x + y)
	// fmt.Println(x - y)
	// fmt.Println(x * y)
	// fmt.Println(x / y)
	// fmt.Println(x % y)


	// a :=12
	// if a>5 && 10<a{
	// 	fmt.Println("pass")
	// }
	// for i := 0; i<= 4; i++ {
	// 	if i == 3{
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// for i := 0; i <= 2; i++ {
	// 	for j :=0; j <= 2; j++ {
	// 		fmt.Println(i, "-", j)
	// 	}
	// }

	// arr := [...]int{2, 4, 6, 8, 10}

	// for i := 0; i <= 4; i++ {
	// 	fmt.Println(arr[i])
	// }

	// for i := 0; i <= 10; i++{
	// 	fmt.Println(i)
	// }
		// result01, result02 := cal(6, 3)
		// fmt.Println(result01, result02)
	
	// func(greeting string) {
	// 	fmt.Println(greeting)
	// }("Good evening")
	// result :=cal(10, 5)
	// fmt.Println(result)
}