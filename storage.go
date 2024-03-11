package main

import "fmt"

type DBStore struct {
	db map[int]*Course
}

func NewDB() (*DBStore, error) {
	return &DBStore{
		db: make(map[int]*Course),
	}, nil
}

func (s *DBStore) CreateCourse(a *Course) error {

	s.db[a.CourseId] = a
	fmt.Printf("created course %+v\n", a)

	return nil

}

func (s *DBStore) UpdateCourse(a *Course) error {

	if _, ok := s.db[a.CourseId]; ok {
		s.db[a.CourseId] = a

	} else {
		return fmt.Errorf("course ID not found")
	}

	fmt.Printf("Updated course %+v\n", a)

	return nil

}
