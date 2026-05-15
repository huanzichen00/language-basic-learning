package main

import (
	"errors"
	"fmt"
)

type Student struct {
	ID    int
	Name  string
	Age   int
	Score int
}

func (s Student) IsPassed() bool {
	return s.Score >= 60
}

func (s Student) PrintInfo() {
	fmt.Println("ID:", s.ID)
	fmt.Println("Name:", s.Name)
	fmt.Println("Age:", s.Age)
	fmt.Println("Score:", s.Score)
}

func (s Student) GrowUpWrong() {
	s.Age++
}

func (s *Student) GrowUp() {
	s.Age++
}

func findStudentByID(students []Student, id int) (Student, error) {
	for _, student := range students {
		if student.ID == id {
			return student, nil
		}
	}

	return Student{}, errors.New("student not found")
}

func updateStudentScore(students []Student, id int, newScore int) error {
	// index, value 省略了 value
	for i := range students {
		if students[i].ID == id {
			students[i].Score = newScore
			return nil
		}
	}
	return errors.New("student not found")
}

func main() {
	students := []Student{
		{
			ID:    1,
			Name:  "Tom",
			Age:   20,
			Score: 90,
		},
		{
			ID:    2,
			Name:  "Rose",
			Age:   19,
			Score: 96,
		},
	}

	err := updateStudentScore(students, 2, 100)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("students:", students)

	err = updateStudentScore(students, 99, 80)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("students:", students)
}
