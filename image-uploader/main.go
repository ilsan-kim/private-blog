package main

import (
	"flag"
	"fmt"
	"github.com/google/uuid"
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
                    Image: <input type="file" name="imageFile"> </br>
                    MD: <input type="file" name="mdfile"> </br>
                    <input type="submit" value="Upload">
                </form>
            `))
			return
		}

		if r.Method == "POST" {
			// Parse the multipart form
			r.ParseMultipartForm(10 << 20) // 10 MB

			// Handle MD file upload
			mdfile, mdHandler, err := r.FormFile("mdfile")
			if err == nil {
				defer mdfile.Close()

				// Save MD file
				dstPath := filepath.Join(config.MDFileUploadPath, filepath.Base(mdHandler.Filename))
				saveFile(mdfile, dstPath)
				fmt.Printf("Uploaded MD File: %+v\n", mdHandler.Filename)
			}

			// Handle Image file upload
			imageFile, imageHandler, err := r.FormFile("imageFile")
			if err == nil {
				defer imageFile.Close()

				// Generate UUID for image filename
				newFilename := uuid.New().String() + filepath.Ext(imageHandler.Filename)

				// Save image file
				dstPath := filepath.Join(config.ThumbnailUploadPath, newFilename)
				saveFile(imageFile, dstPath)
				fmt.Printf("Uploaded Image File: %+v\n", newFilename)
			}

			fmt.Fprintf(w, "Successfully Uploaded Files\n")
		}
	}
}

func saveFile(src io.Reader, dstPath string) {
	dst, err := os.Create(dstPath)
	if err != nil {
		log.Printf("Error creating file %s: %v", dstPath, err)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		log.Printf("Error saving file %s: %v", dstPath, err)
		return
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
