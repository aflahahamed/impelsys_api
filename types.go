package main

type CourseRequest struct {
	CourseId    int    `json:"id"`
	CourseName  string `json:"couresename"`
	CoursePrice int    `json:"courseprice"`
	AuthorName  string `json:"authorname"`
	Website     string `json:"authorwebsite"`
}

type Course struct {
	CourseId    int     `json:"id"`
	CourseName  string  `json:"couresename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	AuthorName    string `json:"authorname"`
	AuthorWebsite string `json:"authorwebsite"`
}

type APIError struct {
	Error string `json:"error"`
}

func NewCourse(coursename, authorname, website string, courseprice, courseid int) *Course {

	return &Course{
		CourseId:    courseid,
		CourseName:  coursename,
		CoursePrice: courseprice,
		Author: &Author{
			AuthorName:    authorname,
			AuthorWebsite: website,
		},
	}
}
