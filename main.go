package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-7235250462978-7237773415044-NiNd2EZeOD0WGGkg89LjYpTu")
	os.Setenv("CHANNEL_ID", "C076QJFA06A")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelID := os.Getenv("CHANNEL_ID")
	fileArr := []string{"A.pdf", "B.pdf"}

	for _, fileName := range fileArr {
		fileContent, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", fileName, err)
			continue
		}

		params := slack.FileUploadParameters{
			Content:     string(fileContent),
			Filename:    fileName,
			Channels:    []string{channelID},
			Filetype:    "auto",
			Title:       fileName,
			InitialComment: "File uploaded by bot",
		}

		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("Error uploading file %s: %v\n", fileName, err)
			continue
		}
		fmt.Printf("File uploaded successfully: %s\n", file.Name)
	}
}
