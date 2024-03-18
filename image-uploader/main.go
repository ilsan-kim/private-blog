package main

import (
	"flag"
	"fmt"
	"github.com/ilsan-kim/private-blog/image-uploader/config"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func uploadFileHandler(config config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`
                <form action="/upload" method="post" enctype="multipart/form-data">
                    thumbnail: <input type="file" name="uploadfile"> </br>
					md: <input type="file" name="mdfile"> </br>
                    <input type="submit" value="Upload">
                </form>
            `))
			return
		}

		if r.Method == "POST" {
			// Parse the multipart form, 10 << 20 specifies a maximum
			// upload of 10 MB files.
			r.ParseMultipartForm(10 << 20)

			file, handler, err := r.FormFile("uploadfile")
			if err != nil {
				fmt.Println("Error Retrieving the File")
				fmt.Println(err)
				return
			}
			defer file.Close()

			mdfile, mdHandler, err := r.FormFile("mdfile")
			if err != nil {
				fmt.Println("Error Retrieving the File")
				fmt.Println(err)
				return
			}
			defer file.Close()

			mdContent, err := io.ReadAll(mdfile)
			if err != nil {
				return
			}
			fmt.Printf("Uploaded File: %+v\n", handler.Filename)
			fmt.Printf("Uploaded File: %+v\n", mdHandler.Filename)

			// Create a new file in the static directory
			dst, err := os.Create(filepath.Join(config.ThumbnailUploadPath, filepath.Base(fmt.Sprintf("%s_thumbnail", mdHandler.Filename))))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer dst.Close()

			// Copy the uploaded file to the created file on the filesystem
			_, err = io.Copy(dst, file)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			dst2, err := os.Create(filepath.Join(config.MDFileUploadPath, filepath.Base(mdHandler.Filename)))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer dst2.Close()

			_, err = dst2.Write(mdContent)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			fmt.Fprintf(w, "Successfully Uploaded File\n")
		}
	}
}

func main() {
	configPath := flag.String("c", "./config.json", "config path")
	flag.Parse()

	log.Println(*configPath)
	conf := config.MustLoadConfig(*configPath)

	http.Handle("/", ipBlockMiddleware(uploadFileHandler(conf), conf.ImageUploadFrom))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(conf.ThumbnailUploadPath))))

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func ipBlockMiddleware(next http.HandlerFunc, allowedIPs []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.Header.Get("X-Forwarded-For")
		clientIP := strings.Split(ip, ",")[0]
		pass := false

		for _, allowedIP := range allowedIPs {
			if allowedIP == "*" {
				next.ServeHTTP(w, r)
				return
			}

			if clientIP == allowedIP {
				pass = true
			}
		}

		if !pass {
			http.Error(w, "Access Denied", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
		return
	}
}
