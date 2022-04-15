package studentUnion

import (
	"fmt"
	"os"
)

// add
// edit
// demo

func showMenu() {
	fmt.Println("welcome to student info manage system")
	fmt.Println("1. add student")
	fmt.Println("2. edit student info")
	fmt.Println("3. demo all student infos")
	fmt.Println("4. exit system")
}

func getInput() *student {
	var (
		id     int
		name   string
		class_ string
	)
	fmt.Println("input stu info")
	fmt.Println("pls input id of stu")
	fmt.Scanf("%d\n", &id)
	fmt.Println("pls input name of stu")
	fmt.Scanf("%d\n", &name)
	fmt.Println("pls input class of class")
	fmt.Scanf("%s\n", &class_)
	stu := NewStudent(id, name, class_)
	return stu

}

func main() {
	sm := newStudentMgr()
	for {
		showMenu()

		var input int
		fmt.Println("pls input the index of student")
		_, err := fmt.Scanf("%d\n", &input)
		if err != nil {
			return
		}
		fmt.Println("user input is:", input)
		switch input {
		case 1:
			stu := getInput()
			sm.addStudent(stu)
		case 2:
			stu := getInput()
			sm.modifyStudent(stu)
		case 3:
			sm.showStudent()
		case 4:
			os.Exit(0)
		}
	}
}
