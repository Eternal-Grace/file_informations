package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Item struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Server struct {
	*mux.Router
	historyItems []Item
}

func NewServer() *Server {
	s := &Server{
		Router:       mux.NewRouter(),
		historyItems: []Item{},
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/import", s.uploadFileInfoItem()).Methods("POST")
	s.HandleFunc("/", s.listFilesInfoHistory()).Methods("GET")
	s.HandleFunc("/history/{id}", s.displayFileInfoData()).Methods("GET")
	s.HandleFunc("/history/{id}", s.deleteFileInfoData()).Methods("DELETE")
}

func (s *Server) uploadFileInfoItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Multipart/form-data
		r.ParseMultipartForm(10 << 20) // Limit of 10Mb

		// retreive file from posted data
		file, handler, err := r.FormFile("image-file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		fmt.Printf("File name: %+v\n", handler.Filename)
		fmt.Printf("File size: %+v\n", handler.Size)
		fmt.Printf("File size: %+v\n", handler.Header)

		// write temporary file on server
		tempFile, err := ioutil.TempFile("temp", "tmp-*.png")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tempFile.Write(fileBytes)

		dir, err := os.Getwd()
		if err != nil {
			http.Error(w, err.Error(), http.StatusMethodNotAllowed)
			return
		}
		fmt.Fprintf(w, "test: %+v\n", dir)

		response := make(map[string]string)
		response["message"] = "Status Created"
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			//
		}
		w.Write(jsonResponse)
	}
}

func (s *Server) listFilesInfoHistory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.historyItems); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) displayFileInfoData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (s *Server) deleteFileInfoData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
