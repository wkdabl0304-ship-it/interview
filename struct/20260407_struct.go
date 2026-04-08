package main

import "fmt"

// 组合（内嵌）：
// 1.子类中写匿名父类，那么子类继承父类所有属性和方法
// 2.子类可以重写属性或方法覆盖父类，如果要使用父类被遮蔽的属性或方法，需要显示调用
// 3.匿名不匿名本质是一样的，但是匿名更加方便（仅限于指针）

type Person struct {
	Name string
	Age  int
}

type Student struct {
	Person
	Name string
}

func (s *Student) String() string {
	return fmt.Sprintf("Student{Name:%s, Age:%d}", s.Name, s.Age)
}

func (s *Student) Print() {
	fmt.Printf("Student{Name:%s, Age:%d}", s.Name, s.Age)
	fmt.Println()
}

func main() {
	s1 := new(Student)
	s1.Name = "Jack"
	s1.Person.Name = "Jack2"
	s1.Age = 18
	fmt.Println(s1)
	s2 := Student{Name: "Tom", Person: Person{Age: 18}}
	fmt.Println(s2)
	s1.Print()
	s2.Print()
}
