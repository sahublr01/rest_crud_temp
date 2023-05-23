package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"rest_crud_temp/utils"
	"strconv"

	"github.com/kataras/golog"
)

type NewFlavors struct {
	Flavor  string `json:"flavor"`
	Type    string `json:"type"`
	URL     string `json:"url"`
	Bitrate string `json:"bitrate,omitempty"`
	Width   string `json:"width,omitempty"`
	Height  string `json:"height,omitempty"`
}
type APIResponse struct {
	Status  bool        `json:"status"`
	Error   int         `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type StreamData struct {
	ID           string        `json:"_id"`
	Kid          interface{}   `json:"_kid"`
	Name         string        `json:"name"`
	At           string        `json:"at"`
	Image        string        `json:"image"`
	Thumb        string        `json:"thumb"`
	Poster       string        `json:"poster"`
	CustomThumb  string        `json:"custom_thumb"`
	CustomPlayer string        `json:"custom_player"`
	CustomPoster string        `json:"custom_poster"`
	IsPrime      int           `json:"isPrime"`
	Duration     int           `json:"duration"`
	Length       int           `json:"length"`
	Vendor       int           `json:"vendor"`
	VendorName   interface{}   `json:"vendor_name"`
	Subtitle     []interface{} `json:"subtitle"`
	Ag           string        `json:"AG"`
	IsLive       int           `json:"isLive"`
	AudioOnly    int           `json:"audioOnly"`
	Cdn          struct {
		Ak string `json:"AK"`
	} `json:"cdn"`
	Imagereel struct {
		Rr int   `json:"rr"`
		Rc int   `json:"rc"`
		Tw int   `json:"tw"`
		Th int   `json:"th"`
		Tc []int `json:"tc"`
	} `json:"imagereel"`
	Flavors           []NewFlavors `json:"flavors"`
	CustomSocialImage interface{}  `json:"custom_social_image"`
	Tags              string       `json:"tags"`
	Istranscode       interface{}  `json:"istranscode"`
}

type Media struct {
	ID     string `json:"_id"`
	Source Source `json:"_source"`
}

type Cover struct {
	Front struct {
		Type    string `json:"type"`
		Videoid string `json:"videoid"`
		Src     struct {
			Poster string `json:"poster"`
			Thumb  string `json:"thumb"`
			Player string `json:"player"`
		} `json:"src"`
	} `json:"front"`
	Back struct {
		Type    string `json:"type"`
		Videoid string `json:"videoid"`
		Src     struct {
			Poster string `json:"poster"`
			Thumb  string `json:"thumb"`
			Player string `json:"player"`
		} `json:"src"`
	} `json:"back"`
}

type Source struct {
	St              string        `json:"_st"`
	Channel         string        `json:"channel"`
	IncludeProducts []interface{} `json:"includeProducts"`
	Stitle          string        `json:"stitle"`
	MediaType       string        `json:"media_type"`
	ID              int           `json:"id"`
	Contenttypes    string        `json:"contenttypes"`
	Images          struct {
		Thumb           string `json:"thumb"`
		Poster          string `json:"poster"`
		Player          string `json:"player"`
		CustomPlayer    string `json:"custom_player"`
		CustomPoster    string `json:"custom_poster"`
		CustomPosterall string `json:"custom_posterall"`
		CustomThumb     string `json:"custom_thumb"`
	} `json:"images"`
	FileStatus  int    `json:"file_status"`
	At          string `json:"at"`
	Status      int    `json:"status"`
	Vendor      int    `json:"vendor"`
	Audio       int    `json:"audio"`
	Product     string `json:"product"`
	Created     string `json:"created"`
	Description string `json:"description"`
	Video       int    `json:"video"`
	Title       string `json:"title"`
	Sid         string `json:"sid"`
	Duration    int    `json:"duration"`
}

type Markers struct {
	ID               string        `json:"_id"`
	Name             string        `json:"name"`
	At               string        `json:"at"`
	Image            string        `json:"image"`
	Thumb            string        `json:"thumb"`
	Poster           string        `json:"poster"`
	CustomThumb      string        `json:"custom_thumb"`
	CustomPlayer     string        `json:"custom_player"`
	CustomPoster     string        `json:"custom_poster"`
	Duration         int           `json:"duration"`
	Length           int           `json:"length"`
	Vendor           int           `json:"vendor"`
	VendorName       interface{}   `json:"vendor_name"`
	Subtitle         []interface{} `json:"subtitle"`
	Ag               string        `json:"AG"`
	Description      string        `json:"description"`
	Title            string        `json:"title"`
	IsLive           int           `json:"isLive"`
	Simulive         int           `json:"simulive"`
	AudioOnly        int           `json:"audioOnly"`
	SimulivePlaylist bool          `json:"simulivePlaylist"`
	StartTime        int64         `json:"startTime,omitempty"`
	EndTime          int64         `json:"endTime,omitempty"`
	ServerTime       string        `json:"serverTime"`
	Cdn              struct {
		Ak string `json:"AK"`
	} `json:"cdn"`
	Imagereel struct {
		Rr int   `json:"rr"`
		Rc int   `json:"rc"`
		Tw int   `json:"tw"`
		Th int   `json:"th"`
		Tc []int `json:"tc"`
	} `json:"imagereel"`
	Flavors           []Flavors   `json:"flavors"`
	Payload           []Source    `json:"payload"`
	CustomSocialImage interface{} `json:"custom_social_image"`
	Tags              string      `json:"tags"`
	Istranscode       interface{} `json:"istranscode"`
}

type Flavors struct {
	Flavor  string `json:"flavor"`
	Type    string `json:"type"`
	URL     string `json:"url"`
	Bitrate string `json:"bitrate,omitempty"`
	Width   string `json:"width,omitempty"`
	Height  string `json:"height,omitempty"`
}

type Index struct {
	Resolution string `json:"resolution"`
	Sid        string `json:"sid"`
	Type       string `json:"type"`
	// Playlist   *m3u8.MasterPlaylist `json:"playlist"`
	URL     *url.URL `json:"url"`
	Bitrate string   `json:"bitrate,omitempty"`
	Width   string   `json:"width,omitempty"`
	Height  string   `json:"height,omitempty"`
}

type SrtReq struct {
	Name    string `json:"name"`
	Default bool   `json:"dafault"`
	Code    string `json:"code"`
	Path    string `json:"path"`
}

func InitTelegraf(w http.ResponseWriter, req *http.Request) map[string]string {
	golog.Info("Init Telegraf Process...")
	golog.Info("here at getting data in InitTelegraf...")
	body, _ := ioutil.ReadAll(req.Body)
	golog.Info(string(body))
	return map[string]string{"data": string(body), "status": "true", "error": "0", "message": string(body)}
}

// func PostData(w http.ResponseWriter, req *http.Request) map[string]string {
// 	return nil
// }

func HandleSubtitleMaster(slikeId, path string, srt []SrtReq) (string, error) {
	golog.Infof("here HandleSubtitleMaster slikeId :: %s", slikeId)
	return "", nil
}

func Download(url, filepath string, getdata bool) ([]byte, string, error) {
	body, _, _, err := utils.SendHTTPRequest(nil, -1, url, nil, nil, nil, "GET", nil, 0, true, 5)
	if err != nil {
		return nil, "", err
	}
	path, _ := utils.SplitDirAndFilename(filepath)
	utils.CheckDirExists(path)
	err = ioutil.WriteFile(filepath, body, 0644)
	if err != nil {
		return nil, "", fmt.Errorf("error in writing file , data - %s error - %s", filepath, err)
	}
	if getdata {
		return body, filepath, nil
	}
	return nil, filepath, nil
}

type Todo struct {
	ID   int
	Task string
}

var todos []Todo

func PostData(w http.ResponseWriter, req *http.Request) ([]Todo, error) {
	var todo Todo

	err := json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		return nil, fmt.Errorf("error decoding Post Data %s", err)
	}

	todos = append(todos, todo)

	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// remove all the w.write
// gracefully error handling
// formatting of error fmt.errorf
// returing of error in every function

func DeleteData(w http.ResponseWriter, req *http.Request) error {
	query := req.URL.Query()

	id, err := strconv.Atoi(query.Get("id"))
	if err != nil {
		return fmt.Errorf("failed to convert id parameter to integer: %s", err)
	}

	for index, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:index], todos[index+1:]...)
			fmt.Printf("Success to delete todo")
			return nil
		}
	}

	return fmt.Errorf("Todo with ID %d not found", id)
}

func UpdateData(w http.ResponseWriter, req *http.Request) error {
	query := req.URL.Query()
	id, err := strconv.Atoi(query.Get("id"))

	if err != nil {
		return fmt.Errorf("failed to convert id parameter to integer: %s", err)
	}

	for index, todo := range todos {
		json.NewDecoder(req.Body).Decode(&todo)

		if todo.ID == id {
			todos[index].ID = todo.ID
			todos[index].Task = todo.Task
			fmt.Printf("Success to update todo")
		}
	}
	return nil
}

// func PostData(w http.ResponseWriter, req *http.Request) error {
// 	var todo Todo

// 	err := json.NewDecoder(req.Body).Decode(&todo)
// 	if err != nil {
// 		return fmt.Errorf("error decoding Post Data %s", err)
// 	}

// 	todos = append(todos, todo)

func GetData(w http.ResponseWriter, req *http.Request) ([]Todo, error) {

	return todos, nil
}

// new fucntion for writing a file for goroutine
// new fucntion for writing a file for goroutine using channels
// limit of go routines in windows as well as linux

// func handleRequest(w http.ResponseWriter, r *http.Request) {
// 	// Start a new Go routine to handle the request go
// 	processRequest(w, r)
// 	// Return an immediate response to the client
// 	fmt.Fprintf(w, "Request received and being processed concurrently.")
// }
