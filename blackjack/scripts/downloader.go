package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func downloadFile(url, filePath string) error {
	out, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer out.Close()

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	_, err = io.Copy(out, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	baseURL := "https://chicodeza.com/wordpress/wp-content/uploads/"
	baseTargetFile := "torannpu-illust"
	baseExt := ".png"
	outputDir := "/home/suhrkawase/work/projects/tourofgo/blackjack/assets"
	max := 52

	for i := 1; i <= max; i++ {
		url := baseURL + baseTargetFile + strconv.Itoa(i) + baseExt
		fileName := fmt.Sprintf("%s/card%d.png", outputDir, i)

		fmt.Printf("Downloading: %s\n", url)
		err := downloadFile(url, fileName)
		fmt.Println("Done")

		if err != nil {
			panic(err)
		}

		time.Sleep(3000 * time.Millisecond)
	}
}
