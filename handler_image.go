package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Data struct {
	FileName string `json:"fileName"`
}

func HandlerImage(w http.ResponseWriter, r *http.Request) {

	// Extract data from the body

	contentType := strings.Split(r.Header.Get("Content-Type"), ";")[0]

	var fileName string = ""

	switch contentType {
	case "application/json":
		var data Data

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			ResponseWithError(w, 500, err.Error())
			return
		}

		fileName = data.FileName
	case "multipart/form-data":

		err := r.ParseMultipartForm(10 << 20)

		if err != nil {
			ResponseWithError(w, 500, err.Error())
		}

		fileName = r.FormValue("fileName")
	}

	file, header, err := r.FormFile("file")

	if err != nil {
		ResponseWithError(w, 500, err.Error())
		return
	}

	defer file.Close()
	ext := filepath.Ext(header.Filename)

	if fileName == "" {
		fileName = strings.ReplaceAll(header.Filename, " ", "_")
	} else {
		fileName = fmt.Sprintf("%s%s", fileName, ext)
	}

	outFile, err := os.Create(fileName)

	if err != nil {
		ResponseWithError(w, 500, err.Error())
		return
	}

	defer outFile.Close()

	_, err = io.Copy(outFile, file)

	if err != nil {
		ResponseWithError(w, 500, err.Error())
		return
	}

	fmt.Fprintln(w, "File uploaded successfully!")
}
