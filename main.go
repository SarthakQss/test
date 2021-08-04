package main

import "fmt"

type User interface {
	getUserInfo() string
	
}

type StudentInfo struct{
	Fname string
	LName string
}
type TeacherInfo struct{
	Fname string
	
}

func main() {
	SI := StudentInfo{
		Fname: "Sarthak",
		LName: "Pruthi",
	}
	TI := TeacherInfo{
		Fname: "Qss Technosoft",
	}

	getUserInfo(SI)
	getUserInfo(TI)
}

func getUserInfo(b User) {
	fmt.Println(b.getUserInfo())
}

func (StudentInfo) getUserInfo() string {
	
	return "Student"
}

func (TeacherInfo) getUserInfo() string {
	return "Teacher"
}
