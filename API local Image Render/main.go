package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"

	webp "github.com/deepteams/webp"
)

func imageHandler(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if len(data) == 0 {
		url := r.URL.Query().Get("url")
		if url == "" {
			url = "https://miro.medium.com/v2/resize:fit:640/format:webp/0*YISbBYJg5hkJGcQd.png"
		}

		resp, err := http.Get(url)
		if err != nil {
			http.Error(w, "Failed to fetch image", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			http.Error(w, "Invalid image", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
		io.Copy(w, resp.Body)
		return
	}

	img, format, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		img, err = webp.Decode(bytes.NewReader(data))
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid image: %v", err), http.StatusBadRequest)
			return
		}
		format = "webp"
	}

	if format == "png" {
		w.Header().Set("Content-Type", "image/png")
		png.Encode(w, img)
		return
	}

	if format == "webp" {
		w.Header().Set("Content-Type", "image/webp")
		err = webp.Encode(w, img, &webp.EncoderOptions{Quality: 90})
		if err != nil {
			http.Error(w, "Failed to encode WebP", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	jpeg.Encode(w, img, &jpeg.Options{Quality: 90})
}

func main() {

	http.HandleFunc("/", imageHandler)
	fmt.Println("Servidor rodando na porta :5080")

	err := http.ListenAndServe("192.168.100.9:5080", nil)
	if err != nil {
		fmt.Println("Erro na API:", err)
	}

}
