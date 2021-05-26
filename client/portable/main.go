package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Settings struct {
	HostName  string `json:"hostname"`  // The host name to call home to
	GroupID   string `json:"groupID"`   // The group ID to sent with the data
	Frequency uint32 `json:"frequency"` // How many minutes between sending data
}

func getSettings() Settings {
	var objSettings Settings

	objReader := bufio.NewReader(os.Stdin)
	fmt.Print("Host URL: ")
	strlHostName, _ := objReader.ReadString('\n')
	strlHostName = strings.Replace(strlHostName, "\n", "", -1)
	// Add Host Check

	fmt.Print("Group ID: ")
	strlGroupID, _ := objReader.ReadString('\n')
	strlGroupID = strings.Replace(strlGroupID, "\n", "", -1)
	// Add Group Check

	fmt.Print("Frequency (in seconds): ")
	var intlFrequency uint32
	_, objError := fmt.Scan(&intlFrequency)

	if objError != nil {
		fmt.Println("Problem with Frequnecy input detaulting to 3200")
		intlFrequency = 3200
	}

	objSettings = Settings{
		HostName:  strlHostName,
		GroupID:   strlGroupID,
		Frequency: intlFrequency,
	}

	return objSettings
}

func writeSettings(objSettings Settings, strFilename string) {
	fileSettings, objError := json.MarshalIndent(objSettings, "", " ")
	if objError != nil {
		fmt.Println(objError.Error())
	}
	fmt.Println("JSON File: " + string(fileSettings))
	_ = ioutil.WriteFile(strFilename, fileSettings, 0644)
}

func printSettings(objSettings Settings) {
	fmt.Println("Host URL: " + objSettings.HostName)
	fmt.Println("Group ID: " + objSettings.GroupID)
	fmt.Println("Frequency (in seconds): " + fmt.Sprint(objSettings.Frequency))
}

func main() {
	var strSettingsFileName string = "settings.json"
	fmt.Println("Inventory Itemization")
	var objSettings Settings
	fileSettings, objError := ioutil.ReadFile(strSettingsFileName)

	if objError != nil {
		objSettings = getSettings()
		writeSettings(objSettings, strSettingsFileName)
	} else {
		_ = json.Unmarshal([]byte(fileSettings), &objSettings)
	}

	fmt.Println("Read the following settings: ")
	printSettings(objSettings)
}
