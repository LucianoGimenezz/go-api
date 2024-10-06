package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	// "path/filepath"
)

func HandlerImage(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")

	if err != nil {
		ResponseWithError(w, 500, err.Error())
		return
	}

	defer file.Close()
	// ext := filepath.Ext(header.Filename)
	fileName := strings.ReplaceAll(header.Filename, " ", "_")
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
