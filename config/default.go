package config

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"rest_crud_temp/utils"

	"github.com/kataras/golog"
	"github.com/spf13/viper"
)

// build info
var (
	Build_Version = "dev"
	Build_Time    = ""
	Build_User    = "dev"
	Build_Commit  = ""
)

func GetBuildInfo() []byte {
	payload, _ := json.Marshal(map[string]string{"version": Build_Version, "user": Build_User, "time": Build_Time, "commit": Build_Commit})
	return payload
}

var Props *Config

type Config struct {
	Verbose uint8
	Env     string
	// Server                string
	// FileStore             string
	// URI                   []string
	// SubtitleStore         string
	// AudioStore            string
	Listen string
	// APIKey                string
	// Polling               int64
	// ProcessLimit          int
	// Disp                  bool
	// Priority              []string
	// Process               []string
	// Ffmpeg                string
	// FFProbe               string
	// ThumbDir              string
	// UploadDir             string
	// MediaReelUrl          string
	// TelegramBot           string
	// TelegramChat          string
	// Magick                string
	// BaseFolderImagesPng   string
	// Port                  string
	// Code                  string
	// ContentType           string
	// DlFolder              bool
	// DownloadBytes         int64
	// Ext                   string
	// Filename              string
	// ID                    string
	// InputtedMimeType      []string
	// Kind                  string
	// Notcreatetopdirectory bool
	// OverWrite             bool
	// Resumabledownload     string
	// SearchID              string
	// ShowFileInf           bool
	// Size                  int64
	// Skip                  bool
	// SkipError             bool
	// URL                   string
	// WorkDir               string
	// DownloadSuccessful    bool
}

func Parse(location string) bool {
	if location == "" {
		location = utils.GetEnv("CONFIG", "conf/")
		location = filepath.Join(location, "conf.yaml")
	}

	viper.SetConfigFile(location)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file %s read failed. %v\n", location, err)
	}
	err = viper.GetViper().Unmarshal(&Props)
	if err != nil {
		fmt.Printf("sfu config file %s loaded failed. %v\n", location, err)
	}
	UpdateLogLevel()

	// utils.CheckDirExists(Props.FileStore)
	// utils.CheckDirExists(Props.SubtitleStore)
	// utils.CheckDirExists(Props.AudioStore)
	return true
}

func UpdateLogLevel() {
	if Props.Verbose == 0 {
		golog.SetLevel("debug")
	} else if Props.Verbose == 1 {
		golog.SetLevel("info")
	} else if Props.Verbose == 2 {
		golog.SetLevel("warn")
	}
}

// func NotifyEvent(msg string) {
// 	utils.NotifyTelegram(fmt.Sprintf("Msg:%s Server:%s", msg, Props.Server), Props.TelegramBot, Props.TelegramChat)
// }
