package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	var router *mux.Router

	router.HandleFunc(
		"/",
		func(rw http.ResponseWriter, r *http.Request) {
			location := os.Getenv("DATA")
			file, err := os.Open(location)
			if err != nil {
				log.Fatalf("provided DATA isn't a proper location | %s", err.Error())
			}
			defer file.Close()
			stats, err := file.Stat()
			if err != nil {
				log.Fatalf("failed stating DATA input | %s", err.Error())
			}
			if !stats.IsDir() {
				log.Fatal("provided DATA location isn't a directory")
			}
			workfile := fmt.Sprintf("%s/data.txt", location)
			if _, err := os.Stat(workfile); err != nil { // le fichier n'existe pas
				f, err := os.Create(workfile)
				if err != nil {
					log.Fatalf("failed creating file on %s | %s ", workfile, err.Error())
				}
				if _, err := f.Write([]byte("1")); err != nil {
					log.Fatalf("failed writing to file on %s | %s", workfile, err.Error())
				}
			} else { // le fichier existe
				data, err := os.ReadFile(workfile)
				if err != nil {
					log.Fatalf("failed reading file on %s | %s ", workfile, err.Error())
				}
				strdata := string(data)
				visits, err := strconv.Atoi(strdata)
				if err != nil {
					log.Fatalf("failed reading file data, not a number | %s", err.Error())
				}
				if err := os.WriteFile(workfile, []byte(strconv.Itoa(visits+1)), 0644); err != nil {
					log.Fatalf("failed writing to file on %s | %s", workfile, err.Error())
				}
			}
		},
	)

	if err := http.ListenAndServe("0.0.0.0:80", router); err != nil {
		log.Fatalf("failed starting server | %s", err.Error())
	}

}
