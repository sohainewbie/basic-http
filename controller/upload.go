package controller

import (
	"basic-http/model"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

var AllowedExtension = map[string]bool{
	".jpeg": true,
	".png":  true,
	".jpg":  true,
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	authToken := "sohayUloader"
	var dataRender = map[string]interface{}{
		"authToken": authToken,
		"success":   "",
	}

	if r.Method == "POST" {
		r.Body = http.MaxBytesReader(w, r.Body, 8*1024*1024) // 8 Mb
		//The whole request body is parsed and up to a total of maxMemory bytes
		if err := r.ParseMultipartForm(1024); err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		auth := r.FormValue("auth")
		if len(auth) == 0 || auth != authToken {
			http.Error(w, "The auth token not match!", http.StatusForbidden)
			return
		}

		uploadedFile, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		defer uploadedFile.Close()

		dir, err := os.Getwd()
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		filename := handler.Filename
		contentType := filepath.Ext(filename)

		if !AllowedExtension[contentType] {
			http.Error(w, "Only format image is allowed!", http.StatusForbidden)
			return
		}

		fileLocation := filepath.Join(dir, "view/files", filename)
		targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, uploadedFile); err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		stat, _ := targetFile.Stat()
		size := stat.Size()
		userAgent := r.UserAgent()

		urlFile, err := model.InsertLogFiles(size, userAgent, filename, contentType)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		r.Method = "GET"
		dataRender["success"] = fmt.Sprintf("File success to upload, you can access %s", urlFile)
	}

	if r.Method == "GET" {
		var filepath = path.Join("view", "upload.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err = tmpl.Execute(w, dataRender); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

}
