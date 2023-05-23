package app

import (
	"net/http"
	
	"github.com/kataras/golog"
)

// error handling done in controllers whcih was not done before
func PostAPI(w http.ResponseWriter, req *http.Request) APIResponse {
	golog.Info("almost there... ")

	// Created a channel to receive the result from the goroutine
	resultCh := make(chan APIResponse)

	// goroutine to execute the PostData function
	go func() {
		data, err := PostData(w, req)
		if err != nil {
			golog.Error(err)
			resultCh <- APIResponse{Data: nil, Status: false, Error: 1, Message: "Failed to Create Data"}
			return
		}
		resultCh <- APIResponse{Data: data, Status: true, Error: 0, Message: "Data Created Successfully"}
	}()

	// Wait for the goroutine to complete and receive the result
	result := <-resultCh

	return result
}

func GetAPI(w http.ResponseWriter, req *http.Request) APIResponse {

	golog.Info("almost there... ")
	data, err := GetData(w, req)
	if err != nil {
		golog.Error(err)
		return APIResponse{Data: nil, Status: false, Error: 1, Message: "Failed to Retrieve Data"}
	}

	return APIResponse{Data: data, Status: true, Error: 0, Message: "Data Retrieved Successfully"}
}
func PutAPI(w http.ResponseWriter, req *http.Request) APIResponse {

	golog.Info("almost there... ")
	err := UpdateData(w, req)
	if err != nil {
		golog.Error(err)
		return APIResponse{Data: nil, Status: false, Error: 1, Message: "Failed to Retrieve Data"}
	}

	return APIResponse{Data: UpdateData(w, req), Status: true, Error: 0, Message: "Data Retrieved Successfully"}
}

// func PostAPI(w http.ResponseWriter, req *http.Request) APIResponse {
// 	golog.Info("almost there... ")

// 	err := PostData(w, req)
// 	if err != nil {
// 		golog.Error(err)
// 		return APIResponse{Data: nil, Status: false, Error: 1, Message: "Failed to Create Data"}
// 	}
// 	return APIResponse{Data: PostData(w, req), Status: true, Error: 0, Message: "Data Created Successfully"}
// }

func DeleteAPI(w http.ResponseWriter, req *http.Request) APIResponse {

	golog.Info("almost there... ")
	err := DeleteData(w, req)
	if err != nil {
		golog.Error(err)
		return APIResponse{Data: nil, Status: false, Error: 1, Message: "Failed to delete requested data"}
	}

	return APIResponse{Data: DeleteData(w, req), Status: true, Error: 0, Message: "requested data deleted successfully"}
}
