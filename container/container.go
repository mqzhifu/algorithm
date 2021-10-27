package container

import (
	"math/rand"
	"time"
)

//获取一个随机数：int
func GetRandIntNum(max int) int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max) + 1
}
