package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	// "io/ioutil"
	"math"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/kataras/golog"
)

var httpDefaultClient = &http.Client{Timeout: 40 * time.Second}

func PathIs(name []byte, r *http.Request) bool {
	return bytes.Equal([]byte(r.URL.Path), name)
}

func GetEnv(key, fallback string) (value string) {
	value = fallback
	if val, ok := os.LookupEnv(key); ok {
		value = val
	}
	return
}

// SplitDirAndFilename :
func SplitDirAndFilename(path string) (string, string) {
	lastSlash := strings.LastIndex(path, "/")
	if lastSlash == -1 {
		return "", path
	}
	return path[:lastSlash], path[lastSlash+1:]
}

func SendHTTPRequest(reqhosts []string, seed int64, url string, payloadBytes []byte, headers map[string]string, formData map[string]string, method string, httpClient *http.Client, retry int, ssl bool, maxRetry int) ([]byte, map[string][]string, int, error) { //response in bytes,res headers, error, shouldRetry
	hosts := append([]string{}, reqhosts...)
	if (len(hosts) == 0 || hosts[0] == "") && seed != -1 {
		return nil, nil, -1, fmt.Errorf("no host provided for %s", url)
	}
	if retry > maxRetry {
		return nil, nil, -2, fmt.Errorf("max retry limit of %d exceeded for url %s", retry, url)
	}
	var urlString, hostAddress, protocol string
	randomHostIndex := 0
	if ssl {
		protocol = "https"
	} else {
		protocol = "http"
	}
	if seed != -1 {
		hostAddress, randomHostIndex = GetRandomHostAddress(hosts, seed)
		urlString = protocol + "://" + path.Join(hostAddress, url)
	} else {
		urlString = url
	}

	var body io.Reader
	if payloadBytes != nil {
		body = bytes.NewReader(payloadBytes)
	} else if formData != nil {
		payload := &bytes.Buffer{}
		writer := multipart.NewWriter(payload)
		for k, v := range formData {
			writer.WriteField(k, v)
		}
		if headers == nil {
			headers = map[string]string{}
		}
		headers["Content-Type"] = writer.FormDataContentType()
		writer.Close()
		body = payload
	} else {
		body = nil
	}

	req, err := http.NewRequest(method, urlString, body)
	if err != nil {
		golog.Infof("request object could not be created for making a curl hit: %s", err)
		return nil, nil, -1, fmt.Errorf("request object could not be created for making a curl hit: %s", err)
	}

	var hasContentType = false
	for k, v := range headers {
		req.Header.Set(k, v)
		if k == "Content-Type" {
			hasContentType = true
		}
	}
	if !hasContentType {
		req.Header.Set("Content-Type", "application/json")
	}

	var client *http.Client
	if httpClient != nil {
		client = httpClient
	} else {
		client = httpDefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		if retry == maxRetry {
			golog.Errorf("Error in making http request on %s on try count =%d, error string - %s", urlString, retry, err.Error())
			return nil, nil, 0, fmt.Errorf("max retry http request failed %s", err)
		}
		golog.Infof("Error in making http request on %s on try count =%d, error string - %s", urlString, retry, err.Error())
		// golog.Errorf("Error in making http request on %s on try count =%d, error string - %s", urlString, retry, err.Error())
		time.Sleep((time.Duration(math.Pow(2, float64(retry)))) * time.Second)
		if len(hosts) > 1 {
			seed = seed + int64(math.Pow10(retry))
			hosts[randomHostIndex] = hosts[len(hosts)-1]
			hosts = hosts[:len(hosts)-1]
		}
		return SendHTTPRequest(hosts, seed, url, payloadBytes, headers, formData, method, httpClient, retry+1, ssl, maxRetry)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return respBody, resp.Header, resp.StatusCode, fmt.Errorf("http request failed with status: %d, body: %s", resp.StatusCode, respBody)
	}
	if err != nil {
		if retry == maxRetry {
			golog.Errorf("Error read body on %s on try count =%d, error string - %s", urlString, retry, err.Error())
			return nil, nil, 0, fmt.Errorf("max retry read body failed %s", err)
		}
		golog.Infof("Error read body on %s on try count =%d, error string - %s", urlString, retry, err.Error())
		// golog.Errorf("Error in making http request on %s on try count =%d, error string - %s", urlString, retry, err.Error())
		time.Sleep((time.Duration(math.Pow(2, float64(retry)))) * time.Second)
		if len(hosts) > 1 {
			seed = seed + int64(math.Pow10(retry))
			hosts[randomHostIndex] = hosts[len(hosts)-1]
			hosts = hosts[:len(hosts)-1]
		}
		return SendHTTPRequest(hosts, seed, url, payloadBytes, headers, formData, method, httpClient, retry+1, ssl, maxRetry)
	}
	return respBody, resp.Header, resp.StatusCode, err
}

func GetRandomHostAddress(hostsArray []string, seed int64) (string, int) {
	if len(hostsArray) == 1 {
		return hostsArray[0], 0
	}
	rand.Seed(seed)
	randomIndex := rand.Intn(len(hostsArray))
	randomHost := hostsArray[randomIndex]
	return randomHost, randomIndex //returning random host
}
func CheckDirExists(dirPath string) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			golog.Error(err)
		}
	}
}
func SendJSONResponse(w http.ResponseWriter, data interface{}, message string, errMessage string, statusCode int, errorCode int) {
	response := map[string]interface{}{
		"data":    data,
		"message": message,
	}
	if errMessage != "" {
		response["error"] = errMessage
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}
func SetCorsHeaders(w *http.ResponseWriter, r *http.Request, allowCredentials bool) {
	// Set the Access-Control-Allow-Origin header to the value of the Origin header
	origin := r.Header.Get("Origin")
	if origin != "" {
		(*w).Header().Set("Access-Control-Allow-Origin", origin)
	}

	// Set the Access-Control-Allow-Methods header to allow GET, POST, OPTIONS, PUT, and DELETE
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")

	// Set the Access-Control-Allow-Headers header to allow Content-Type, Authorization, and any headers specified in the request
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	requestHeaders := r.Header.Get("Access-Control-Request-Headers")
	if requestHeaders != "" {
		(*w).Header().Set("Access-Control-Allow-Headers", requestHeaders)
	}

	// Set the Access-Control-Allow-Credentials header if specified
	if allowCredentials {
		(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	}
}
