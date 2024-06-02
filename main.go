package main

import (
	"fmt"
	"student-management/student"
)

func main() {
    // Добавляем студентов
    student.AddStudent(student.Student{ID: 1, Name: "Alice", Age: 20, Grade: "A"})
    student.AddStudent(student.Student{ID: 2, Name: "Bob", Age: 22, Grade: "B"})

    // Выводим список всех студентов
    students := student.ListStudents()
    fmt.Println("All students:", students)

    // Получаем студента по ID
    s, err := student.GetStudent(1)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Student 1:", s)
    }

    // Обновляем студента
    err = student.UpdateStudent(1, student.Student{ID: 1, Name: "Alice Smith", Age: 21, Grade: "A+"})
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Updated student 1:", student.ListStudents())
    }

    // Удаляем студента
    err = student.DeleteStudent(2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Deleted student 2. Remaining students:", student.ListStudents())
    }
}
