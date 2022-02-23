package practice

import "fmt"

type Joseph struct{
	sumPeople int
	Debug int
}

func NewJoseph(sumPeople int)*Joseph{
	joseph := new(Joseph)
	joseph.sumPeople = sumPeople
	joseph.Debug = 1
	return joseph
}

//输出信息，用于debug
func (joseph *Joseph) Print(a ...interface{}) (n int, err error) {
	if joseph.Debug > 0{
		return fmt.Println(a)
	}
	return
}

func(joseph *Joseph) Recursive(n int,m int)int{
	joseph.Print("Recursive n: ",n)
	if n == 1 {
		joseph.Print("in last number ,start return...")
		return 0
	}

	location := joseph.Recursive(n - 1,m)
	joseph.Print("location:",location, " n:",n , " m:",m)
	lastLocation := (location +  m ) % ( n  )
	joseph.Print("lastLocation:",lastLocation)
	return lastLocation
}

func TestJoseph(){
	a := 1
	fmt.Print(" a = ",a," , left " , a << 1 , " right ", a >> 1, " ")
	b := -1
	fmt.Print(" b = ",b," , left " , b << 1 , " right ", b >> 1, " ")

	return
	n := 10
	m := 4
	joseph := NewJoseph(n)
	joseph.Recursive(n,m)
}