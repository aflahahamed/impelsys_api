package main

type CourseRequest struct {
	CourseId    int    `json:"id"`
	CourseName  string `json:"couresename"`
	CoursePrice int    `json:"courseprice"`
	Fullname    string `json:"fullname"`
	Website     string `json:"website"`
}

type Course struct {
	CourseId    int     `json:"id"`
	CourseName  string  `json:"couresename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

type APIError struct {
	Error string `json:"error"`
}

func NewCourse(coursename, fullname, website string, courseprice, courseid int) *Course {

	return &Course{
		CourseId:    courseid,
		CourseName:  coursename,
		CoursePrice: courseprice,
		Author: &Author{
			Fullname: fullname,
			Website:  website,
		},
	}
}
