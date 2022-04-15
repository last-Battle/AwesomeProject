package studentUnion

import "fmt"

type student struct {
	id    int
	name  string
	class string
}

func NewStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

type studentMgr struct {
	allStudents []*student
}

func newStudentMgr() *studentMgr {
	return &studentMgr{
		allStudents: make([]*student, 0, 100),
	}
}

func (s *studentMgr) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}

func (s *studentMgr) modifyStudent(newStu *student) {
	for i, v := range s.allStudents {
		if newStu.id == v.id {
			s.allStudents[i] = newStu
			return
		}
	}
	fmt.Println("error")
}

func (s *studentMgr) showStudent() {
	for _, v := range s.allStudents {
		fmt.Println("student id %d name %s class %s\n", v.id, v.name, v.class)
	}

}
