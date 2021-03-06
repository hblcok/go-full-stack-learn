package method

import (
	"fmt"
	"testing"
)

type Point struct {
	X, Y int
}

func (p Point) Distance(q Point) {
	fmt.Println(p.X, q.Y)
}

func (p *Point) ptrDistance(q Point) {
	fmt.Println(p.X, q.Y)
}

// TestMethod 方法变量和方法表达式 study
func TestMethod(t *testing.T) {
	p := Point{
		X: 1,
		Y: 2,
	}
	q := Point{
		X: 3,
		Y: 4,
	}
	// 调用方式1：普通调用
	p.Distance(q) // 1 2

	// 调用方式2：使用方法变量
	distance := p.Distance
	distance(q) // 1 4
	// 1 4
}
