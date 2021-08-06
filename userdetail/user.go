package userdetail

import (
	"fmt"
)

type StudentInfo struct {
	Id          int
	Fname       string
	Lname       string
	IsActive    bool
	ContactInfo int
	//	DOB         time.Time
}

type StudentList struct {
	Data []StudentInfo
}

func (d StudentInfo) AddUser() StudentList {

	userInfos := StudentList{}

	userInfos.Data = append(userInfos.Data, d)

	return userInfos

}

func (d StudentInfo) GetUserAllDetail() {
	fmt.Println(d)

}

func (pointerToPerson *StudentInfo) UpdateUserName(newFirstName string) {
	(*pointerToPerson).Fname = newFirstName
}

func (d StudentInfo) GetUserUpdatedDetails() {
	fmt.Println("Updated User First Name")
	fmt.Println("First Name: " + d.Fname + " Last Name : " + d.Lname)

}
