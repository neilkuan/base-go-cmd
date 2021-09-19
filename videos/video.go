package videos

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Video struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Imageurl    string `json:"Imageurl,omitempty"`
	Url         string `json:"url,omitempty"`
}

func GetVideos() (videos []Video) {
	fileBytes, err := ioutil.ReadFile("./videos.json")

	if err != nil {
		fmt.Println("./videos.json Not Found, Create one for you")
		ioutil.WriteFile("./videos.json", []byte("[]"), 0644)
	}

	_ = json.Unmarshal(fileBytes, &videos)
	return videos
}

func SaveVideo(videos []Video) {
	videoBytes, err := json.Marshal(videos)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./videos.json", videoBytes, 0644)
	if err != nil {
		panic(err)
	}
}
