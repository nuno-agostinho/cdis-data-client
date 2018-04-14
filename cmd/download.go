package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/uc-cdis/cdis-data-client/jwt"
)

/* performing function of download data */
func RequestDownload(resp *http.Response) *http.Response {
	/*
		Download file from given url encoded in resp
	*/

	if resp == nil {
		return nil
	}

	msg := jwt.JsonMessage{}

	str := jwt.ResponseToString(resp)
	if strings.Contains(str, "Can't find a location for the data") {
		log.Fatalf("The provided uuid is not found!!!")
	}

	jwt.DecodeJsonFromString(str, &msg)
	if msg.Url == "" {
		fmt.Println("Can not get url from " + str)
		return nil
	}

	presignedDownloadURL := msg.Url
	fmt.Println("Downloading data ...")

	respDown, err := http.Get(presignedDownloadURL)
	if err != nil {
		panic(err)
	}

	return respDown
}

func DownloadAFile(uuid string) {
	/*
		Download a file from given an uuid
	*/

	request := new(jwt.Request)
	configure := new(jwt.Configure)
	function := new(jwt.Functions)

	function.Config = configure
	function.Request = request

	endPointPostfix := "/user/data/download/" + uuid

	respDown := function.DoRequestWithSignedHeader(RequestDownload, profile, "", endPointPostfix)

	if respDown == nil {
		fmt.Println("Error !!!")
	} else {
		out, err := os.Create(file_path)
		if err != nil {
			log.Fatalf(err.Error())
		}
		defer out.Close()
		defer respDown.Body.Close()
		_, err = io.Copy(out, respDown.Body)
		if err != nil {
			panic(err)
		}
	}
}

// represent to download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download a file from a UUID",
	Long: `Gets a presigned URL for a file from a UUID and then downloads the specified file.
Examples: ./cdis-data-client download --profile user1 --uuid 206dfaa6-bcf1-4bc9-b2d0-77179f0f48fc --file=~/Documents/file_to_download.json 
`,
	Run: func(cmd *cobra.Command, args []string) {
		if uuid_list_file == "" {
			if file_path == "" {
				log.Fatalf("Need to provide --file option !!!")
			}

			if uuid == "" {
				log.Fatalf("Need to provide --uuid option !!!")
			}
		} else {
			if save_dir == "" {
				log.Fatalf("Need to provide --out option !!!")
			}
		}

		if uuid_list_file == "" {
			DownloadAFile(uuid)
		} else {
			content, err := ioutil.ReadFile(uuid_list_file)
			if err != nil {
				panic(err)
			}
			lines := strings.Split(string(content), "\n")
			for i := 0; i < len(lines); i++ {
				words := strings.Split(lines[i], "\t")
				file_path = save_dir + "/" + uuid
				println("Download " + words[0])
				DownloadAFile(words[0])
			}
		}

		fmt.Println("Done!!!")

	},
}

func init() {
	RootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// putCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// putCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
