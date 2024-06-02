package student

import (
	"errors"
)
type Student struct {
    ID    int
    Name  string
    Age   int
    Grade string
}



var students = make(map[int]Student)

func AddStudent(s Student) {
    students[s.ID] = s
}

func GetStudent(id int) (Student, error) {
    if s, ok := students[id]; ok {
        return s, nil
    }
    return Student{}, errors.New("student not found")
}

func UpdateStudent(id int, s Student) error {
    if _, ok := students[id]; ok {
        students[id] = s
        return nil
    }
    return errors.New("student not found")
}

func DeleteStudent(id int) error {
    if _, ok := students[id]; ok {
        delete(students, id)
        return nil
    }
    return errors.New("student not found")
}

func ListStudents() []Student {
    var list []Student
    for _, s := range students {
        list = append(list, s)
    }
    return list
}
