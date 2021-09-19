package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/neilkuan/base-go-cmd/videos"
)

func main() {

	// get videos from getvideo function.
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	// input for `videos get` command.
	getAll := getCmd.Bool("all", false, "Get all videos")
	getID := getCmd.String("id", "", "Youtube video ID")

	// add videos from getvideo function.
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	// input for `videos add` command.
	addID := addCmd.String("id", "", "Youtube video ID")
	addTitle := addCmd.String("title", "", "Youtube video Title")
	addUrl := addCmd.String("url", "", "Youtube video URL")
	addImageUrl := addCmd.String("imageurl", "", "Youtube video Image URL")
	addDesc := addCmd.String("desc", "", "Youtube video description")

	if len(os.Args) < 2 {
		fmt.Println("Usage: base-go-cmd [command] [arguments]")
		fmt.Println("Example: base-go-cmd add/get subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		HandleGet(getCmd, getAll, getID)
	case "add":
		HandleAdd(addCmd, addID, addTitle, addUrl, addImageUrl, addDesc)
	default:
		os.Exit(1)
	}
}

func HandleGet(getCmd *flag.FlagSet, all *bool, id *string) {
	// parse args...
	getCmd.Parse(os.Args[2:])
	if !*all && *id == "" {
		fmt.Println("Usage: base-go-cmd get --all for all videos or base-go-cmd get --id [ID]")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	// get all videos
	if *all {
		v := videos.GetVideos()

		fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n")
		for _, video := range v {
			fmt.Printf("%s \t %s \t %s \t %s \t %s \n", video.Id, video.Title, video.Url, video.Imageurl, video.Description)
		}
		return
	}

	// get id videos
	if *id != "" {
		v := videos.GetVideos()
		id := *id

		for _, video := range v {
			if video.Id == id {
				fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n")
				fmt.Printf("%s \t %s \t %s \t %s \t %s \n", video.Id, video.Title, video.Url, video.Imageurl, video.Description)
			}
		}
		return
	}

}

func ValidateVideo(addCmd *flag.FlagSet, id *string, title *string, url *string, imageurl *string, desc *string) {
	addCmd.Parse(os.Args[2:])
	if *id == "" || *title == "" || *url == "" || *imageurl == "" || *desc == "" {
		fmt.Println("Usage: base-go-cmd add --id [ID] --title [Title] --url [URL] --imageurl [ImageURL] --desc [Description]")
		addCmd.PrintDefaults()
		os.Exit(1)
	}
}

func HandleAdd(addCmd *flag.FlagSet, id *string, title *string, url *string, imageurl *string, desc *string) {
	ValidateVideo(addCmd, id, title, url, imageurl, desc)

	v := videos.Video{
		Id:          *id,
		Title:       *title,
		Url:         *url,
		Imageurl:    *imageurl,
		Description: *desc,
	}
	videoList := videos.GetVideos()
	videoList = append(videoList, v)
	videos.SaveVideo(videoList)
}
