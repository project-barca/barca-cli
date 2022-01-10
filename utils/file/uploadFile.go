package file

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func UploadToURL(uri string, fileName string) {
	if len(os.Args) != 4 {
		os.Exit(0)
	}

	uploadURL := uri
	fileToUpload := fileName

	//sanity check
	fmt.Println(uploadURL)
	fmt.Println(fileToUpload)

	file, err := os.Open(fileToUpload)
	if err != nil {
		os.Exit(-1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()

	var fileBody bytes.Buffer
	writer := multipart.NewWriter(&fileBody)

	filePart, err := writer.CreateFormFile("fileUploadName", fileInfo.Name())
	if err != nil {
		os.Exit(-1)
	}

	// remember we are using mime - multipart
	_, err = io.Copy(filePart, file)
	if err != nil {
		os.Exit(-1)
	}

	// populate our header with simple data
	_ = writer.WriteField("title", "Data collected on 09th Jan 2022")
	// remember to close writer
	err = writer.Close()
	if err != nil {
		os.Exit(-1)
	}

	request, err := http.NewRequest("POST", uploadURL, &fileBody)
	if err != nil {
		fmt.Println("POST ERROR : ", err)
		os.Exit(-1)
	}
	// set the header with the proper content type for the fileBody's boundary
	request.Header.Set("Content-Type", writer.FormDataContentType())

	// with 10 seconds timeout
	client := &http.Client{Timeout: time.Second * 10}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Client POST error : ", err)
		os.Exit(-1)
	}

	defer response.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading body of response.", err)
		os.Exit(-1)
	}
	fmt.Println("Output: ", string(body))
}
