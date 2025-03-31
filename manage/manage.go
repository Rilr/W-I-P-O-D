package manage

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/Rilr/W-I-P-O-D/bundle_app"
	"github.com/joho/godotenv"
)

func GetTimeSheets(memberID string) TimeSheet {
	// Load the environment variables from the bundled pax.env content
	envContent := bundle_app.ResourceManageEnv.StaticContent
	// Parse the environment variables from the content
	envMap, err := godotenv.Unmarshal(string(envContent))
	if err != nil {
		log.Fatal("Error loading environment variables from pax.env content", err)
	}
	// Manually set environment variables from the map
	for key, value := range envMap {
		_ = os.Setenv(key, value)
	}
	key := os.Getenv("key")
	// Create the webrequest
	baseURL := "http://na.myconnectwise.net/v4_6_release/apis/3.0/time/sheets"
	params := url.Values{}
	params.Add("conditions", "member/id="+memberID)
	params.Add("orderBy", "id desc")
	params.Add("page", "1")
	params.Add("pageSize", "12")
	// Construct the URL
	u, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
	}
	u.RawQuery = params.Encode()
	client := &http.Client{}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		fmt.Println("Error creating the webrequest", err)
	}
	req.Header.Add("clientId", "f38fba56-f78a-4aa0-a38a-ef6de71b13c3")
	req.Header.Add("Authorization", "Basic "+key)
	// Do the webrequest
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error doing the webrequest", err)
	}
	defer res.Body.Close()
	// Handle the response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading the body", err)
	}
	var timeSheets TimeSheet
	err = json.Unmarshal(body, &timeSheets)
	if err != nil {
		fmt.Println("Error decoding the data", err)
	}
	return timeSheets
}

func GetEntries(memberID string) TimeEntries {
	// Load the environment variables from the bundled pax.env content
	envContent := bundle_app.ResourceManageEnv.StaticContent
	// Parse the environment variables from the content
	envMap, err := godotenv.Unmarshal(string(envContent))
	if err != nil {
		log.Fatal("Error loading environment variables from pax.env content", err)
	}
	// Manually set environment variables from the map
	for key, value := range envMap {
		_ = os.Setenv(key, value)
	}
	key := os.Getenv("key")
	// Create the webrequest
	baseURL := "http://na.myconnectwise.net/v4_6_release/apis/3.0/time/entries"
	params := url.Values{}
	params.Add("conditions", "member/id="+memberID+" and workType/id=62")
	params.Add("orderBy", "id desc")
	params.Add("page", "1")
	params.Add("pageSize", "12")
	// Construct the URL
	u, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
	}
	u.RawQuery = params.Encode()
	client := &http.Client{}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		fmt.Println("Error creating the webrequest", err)
	}
	req.Header.Add("clientId", "f38fba56-f78a-4aa0-a38a-ef6de71b13c3")
	req.Header.Add("Authorization", "Basic "+key)
	// Do the webrequest
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error doing the webrequest", err)
	}
	defer res.Body.Close()
	// Handle the response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading the body", err)
	}
	var timeEntries TimeEntries
	err = json.Unmarshal(body, &timeEntries)
	if err != nil {
		fmt.Println("Error decoding the data", err)
	}
	return timeEntries
}
