package main

import "fmt"

type Student struct {
	Num      int
	Name     string
	Chinaese int
	English  int
	Math     int
}

func NewStudent(num int, name string) (*Student, error) {
	if num < 1 || len(name) < 1 {
		return nil, fmt.Errorf("num name empty")
	}
	stu := new(Student)
	stu.Num = num
	stu.Name = name
	return stu, nil
}

func (this *Student) GetAve() (int, error) {
	score := this.Chinaese + this.English + this.Math
	if score == 0 {
		return 0, fmt.Errorf("score is 0")
	}
	return score / 3, nil
}

func main() {
	fmt.Println("Hello")
}
