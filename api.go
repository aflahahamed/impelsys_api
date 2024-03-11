package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	CourseID int
)

type APIServer struct {
	listenaddr string
	store      *DBStore
}

func NewAPIServer(listenAddr string, store *DBStore) *APIServer {
	return &APIServer{
		listenaddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/course", makeHTPPHandlerFunc(s.handleAccount))

	log.Println("JSON api server running on port ", s.listenaddr)

	http.ListenAndServe(s.listenaddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "PATCH" {
		return s.handlePatchCourse(w, r)
	}
	if r.Method == "POST" {
		return s.handleSetCourse(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handlePatchCourse(w http.ResponseWriter, r *http.Request) error {
	updateCourseReq := CourseRequest{}
	if err := json.NewDecoder(r.Body).Decode(&updateCourseReq); err != nil {
		return err
	}
	course := NewCourse(updateCourseReq.CourseName, updateCourseReq.AuthorName, updateCourseReq.Website, updateCourseReq.CoursePrice, updateCourseReq.CourseId)

	err := s.store.UpdateCourse(course)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, course)
}

func (s *APIServer) handleSetCourse(w http.ResponseWriter, r *http.Request) error {
	createCourseReq := CourseRequest{}
	if err := json.NewDecoder(r.Body).Decode(&createCourseReq); err != nil {
		return err
	}
	CourseID += 1
	createCourseReq.CourseId = CourseID

	course := NewCourse(createCourseReq.CourseName, createCourseReq.AuthorName, createCourseReq.Website, createCourseReq.CoursePrice, createCourseReq.CourseId)
	s.store.CreateCourse(course)

	return WriteJSON(w, http.StatusOK, course)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTPPHandlerFunc(f apiFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			//Handle error
			WriteJSON(w, http.StatusBadRequest, APIError{
				Error: err.Error(),
			})
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	{
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(status)

		return json.NewEncoder(w).Encode(v)
	}
}
