package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Fetch() (*PolicyDocument, error) {
	fmt.Println("Downloading policies...")
	resp, err := http.Get(DownloadUrl)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	body = body[len(RemovePrefix):] // It's used for editor config

	var doc PolicyDocument
	err = json.Unmarshal(body, &doc)

	if err != nil {
		return nil, err
	}

	return &doc, nil
}
