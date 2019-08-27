package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	UPLOAD_DIR = "C:/Users/123/Desktop/untitled1/src/demo/photoweb/uploads"
	BASE_DIR   = "C:/Users/123/Desktop/untitled1/src/demo/photoweb"
)

func uploadHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		if err := renderHtml(w, BASE_DIR+"/"+"upload.html", nil); err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
	}
	if req.Method == "POST" {
		f, h, err := req.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		defer t.Close()
		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		http.Redirect(w, req, "/view?id="+filename,
			http.StatusFound)
	}
}
func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}
func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")

	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}
func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images
	if err = renderHtml(w, BASE_DIR+"/"+"list.html", locals); err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
	}
}
func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) (err error) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		return
	}
	return t.Execute(w, locals)
}
func deleteHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "DELETE" {
		err := os.Remove(UPLOAD_DIR + "/" + req.FormValue("imgid"))
		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
		} else {
			io.WriteString(w, "aaaa")
		}
	}
}
func test(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Println("Read failed:", err)
		}
		fmt.Println(req.Header)
		fmt.Println(string(b))
		defer req.Body.Close()
		return

	}
}
func main() {
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/work_sheet", test)
	http.HandleFunc("/view", viewHandler)
	http.HandleFunc("/upload", uploadHandler)
	err := http.ListenAndServe(":7999", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
