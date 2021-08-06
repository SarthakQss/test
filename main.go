package main

import (
	"fmt"
	"bufio"
	"os"
	"github.com/sarthak/user/userdetail"
)

func main() {

	StudentInfo1 := userdetail.StudentInfo{
		Id:    1,
		Fname: "Sarthak",
		Lname: "Pruthi",
		IsActive  :  true,
		ContactInfo : 9671850070,
	}

	StudentInfo2 := userdetail.StudentInfo{
		Id:          2,
		Fname:       "Aman",
		Lname:       "Pruthi",
		 IsActive  :  true,
		 ContactInfo : 9671850070,
	}

	fmt.Println(StudentInfo1.AddUser())
	fmt.Println(StudentInfo2.AddUser())
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Update your First Name for first User")

	newUserName, _ := reader.ReadString('\n')
	StudentInfo1.Fname=newUserName;
	StudentInfo1.GetUserUpdatedDetails()

	reader1 := bufio.NewReader(os.Stdin)
	fmt.Println("Update your First Name for Second User")

	newUserName2, _ := reader1.ReadString('\n')
	StudentInfo2.Fname=newUserName2;
	StudentInfo2.GetUserUpdatedDetails()

	fmt.Println("Enter the User id you want to search for")

	var i int;
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println(err)
	 }

	
if i == StudentInfo1.Id {
	fmt.Println("user match with user id " ,i)
	StudentInfo1.GetUserAllDetail()
}else if i == StudentInfo2.Id {
	fmt.Println("user match with user id " ,i)
	StudentInfo2.GetUserAllDetail()
}else{
	fmt.Println("id is not match")
}


fmt.Println("Enter the User id you want to delete")




}
