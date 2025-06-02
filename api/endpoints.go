package api

import (
	"fmt"
	"net/http"
)

var (
	destFolder, newPassword string
)

func ConfigureDestFolder(response http.ResponseWriter, request *http.Request) {
	destFolder = request.FormValue("dest")
	fmt.Println("New destination folder for the save files", destFolder)
}

func ConfigurePassword(response http.ResponseWriter, request *http.Request) {
	newPassword = request.FormValue("new_pass")
	fmt.Println("New admin password: ", newPassword)
}

func Settings(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Destination Folder: ", destFolder, " admin password: ", newPassword)
}
