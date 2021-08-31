package smi_test

import (
	"fmt"
	"learngowithtests/smi"
	"testing"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangles", shape: smi.Rectangle{Width: 15, Height: 10}, want: 50},
		{name: "Circles", shape: smi.Circle{Radius: 12}, want: 75.39822368615503},
		{name: "Triangles", shape: smi.Triangle{SideA: 10, SideB: 6, SideC: 7}, want: 23},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.want {
				t.Errorf("got %g want %g", got, tt.want)
			}
		})
	}
}

func ExampleRectangle_Perimeter() {
	rectangle := smi.Rectangle{Width: 10, Height: 10}
	perimeter := rectangle.Perimeter()
	fmt.Println(perimeter)
	//Output: 40
}

func ExampleCircle_Perimeter() {
	circle := smi.Circle{Radius: 10}
	perimeter := circle.Perimeter()
	fmt.Println(perimeter)
	//Output: 62.83185307179586
}

func ExampleTriangle_Perimeter() {
	triangle := smi.Triangle{SideA: 10, SideB: 10, SideC: 10}
	perimeter := triangle.Perimeter()
	fmt.Println(perimeter)
	//Output: 30
}

func BenchmarkPerimeter(b *testing.B) {
	perimeterTests := []struct {
		name  string
		shape Shape
	}{
		{name: "Rectangle", shape: smi.Rectangle{Width: 15, Height: 10}},
		{name: "Circle", shape: smi.Circle{Radius: 10}},
		{name: "Triangle", shape: smi.Triangle{SideA: 3, SideB: 4, SideC: 5}},
	}

	for _, tt := range perimeterTests {
		for i := 0; i < b.N; i++ {
			tt.shape.Perimeter()
		}
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangles", shape: smi.Rectangle{Width: 15, Height: 13}, hasArea: 195},
		{name: "Circles", shape: smi.Circle{Radius: 14}, hasArea: 615.7521601035994},
		{name: "Triangles", shape: smi.Triangle{Base: 12, Height: 6}, hasArea: 36},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}

func ExampleRectangle_Area() {
	rectangle := smi.Rectangle{Width: 10, Height: 5}
	area := rectangle.Area()
	fmt.Println(area)
	//Output: 50

}

func ExampleCircle_Area() {
	circle := smi.Circle{Radius: 12}
	area := circle.Area()
	fmt.Println(area)
	//Output: 452.3893421169302

}

func ExampleTriangle_Area() {
	triangle := smi.Triangle{Base: 6, Height: 7}
	area := triangle.Area()
	fmt.Println(area)
	//Output: 21

}

func BenchmarkArea(b *testing.B) {
	benchmarkTests := []struct {
		name  string
		shape Shape
	}{
		{name: "Rectangle", shape: smi.Rectangle{Width: 15, Height: 10}},
		{name: "Circle", shape: smi.Circle{Radius: 10}},
		{name: "Triangle", shape: smi.Triangle{Base: 10, Height: 8}},
	}

	for _, tt := range benchmarkTests {
		for i := 0; i < b.N; i++ {
			tt.shape.Area()
		}
	}
}
