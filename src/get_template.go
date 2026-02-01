package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetTemplate(uri string, filepath string) error {
	template_name := "basic-html.zip"

	// check file template is exist or not

	out, err := os.Create(filepath + template_name)
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(uri)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
