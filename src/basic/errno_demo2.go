package main

import (
	"fmt"
	"math"
)

func main() {
	radius := -3.8
	area, err := circleArea(radius)
	if err != nil {
		//err.Error()
		fmt.Println(err)
		return
	}
	fmt.Println("圆形的面积是：", area)
	fmt.Println("errno demo")
}

// 定义一个结构体，表示错误的类型
type areaError struct {
	msg    string
	redius float64
}

// 实现error接口，就是实现error方法
func (e *areaError) Error() string {
	return fmt.Sprintf("error:半径是， %.2f,%s", e.redius, e.msg)
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"半径是非法的", radius}
	}
	return math.Pi * radius * radius, nil
}
