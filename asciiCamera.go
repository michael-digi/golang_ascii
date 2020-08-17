package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/olekukonko/tablewriter"
)

//Cameras is where the slice that contains all the Camera objects will go
type Cameras struct {
	Cameras []Camera `json:"camera_data"`
}

//Camera is the structure of the json camera object
type Camera struct {
	Name       string `json:"name"`
	URL        string `json:"url"`
	Resolution string `json:"resolution"`
	Location   string `json:"location"`
	Status     string `json:"status"`
}

func main() {
	// open json file
	jsonFile, err := os.Open("./cameras.json")

	if err != nil {
		fmt.Println(err)
	}

	// defer closing of file for after parsing
	defer jsonFile.Close()

	// get the jsonfile as bytes
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// make a variable cameras of type Cameras ready
	var cameras Cameras

	// Unmarshal parses the incoming byte data and stores it in value pointed at by &cameras
	json.Unmarshal(byteValue, &cameras)

	// gets an empty array that will hold arrays ready
	data := [][]string{}

	// assign cameras.Cameras to the easier to read 'cams'
	cams := cameras.Cameras

	// loop through cams, make an array containing the camera attribute values in appropriate order
	// append this array to the nested 'data' array declared
	for i := 0; i < len(cams); i++ {
		temp := []string{cams[i].Name, cams[i].URL, cams[i].Resolution, cams[i].Location, cams[i].Status}
		data = append(data, temp)
	}

	// new tablewriter instance
	table := tablewriter.NewWriter(os.Stdout)

	// string headers for tablewriter
	headers := []string{"NAME", "URL", "RESOLUTION", "LOCATION", "STATUS"}

	table.SetHeader(headers)

	// adds camera array-of-array to ascii table
	table.AppendBulk(data)

	table.Render()
}
