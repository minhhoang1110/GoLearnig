package main

import "fmt"

type Person interface {
	getInfomation() string
}
type Student struct {
	name string
	age  int
}
type Teacher struct {
	name string
	age  int
}

func (student Student) getInfomation() string {
	return "I'm " + student.name + " ," + fmt.Sprint(student.age) + " year old. I'm a student."
}
func (teacher Teacher) getInfomation() string {
	return "I'm " + teacher.name + " ," + fmt.Sprint(teacher.age) + " year old. I'm a teacher."
}
func main() {
	student := Student{"Ben", 12}
	teacher := Teacher{"Henry", 35}
	fmt.Println(student.getInfomation())
	fmt.Println(teacher.getInfomation())
}
