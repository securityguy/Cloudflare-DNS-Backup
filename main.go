//
// Backup DNS information from Cloudflare
//
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"Cloudflare-DNS-Backup/clog"
)

var (
	programName    = "Cloudflare-DNS-Backup"
	programVersion = "0.0.1"
)

type QueryResult map[string]interface{}

func main() {

	// Check number of arguments
	if len(os.Args) < 3 {
		bail()
	}

	// Sanity check
	directory := os.Args[1]
	token := os.Args[2]
	if len(directory) < 1 || len(token) < 10 {
		bail()
	}

	clog.Logf("info", "%s %s starting", programName, programVersion)

	success := backup(directory, token)
	if success {
		clog.Log("info", "Cloudflare DNS backup completed")
	} else {
		clog.Log("error", "Cloudflare DNS backup reported errors")
	}
}

func bail() {
	fmt.Printf("use: %s <directory> <token>\n", os.Args[0])
	os.Exit(1)
}

func backup(directory string, token string) bool {

	// Assume success
	success := true

	// QueryResult structure allows arbitrary JSON
	var result = QueryResult{}

	// Get list of Cloudflare zones
	body, err := cfGet(token, "zones")
	if err != nil {
		clog.Logf("Error obtaining Cloudflare Zone list: ", err.Error())
		return false
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		clog.Logf("Error obtaining Cloudflare Zone list: ", err.Error())
		return false
	}

	// Make sure we received a result
	if _, ok := result["result"]; !ok {
		clog.Log("error", "Cloudflare zone list did not contain any results")
		return false
	}

	// Iterate over list of zones to extract ID from each
	for _, r := range result["result"].([]interface{}) {
		zone := r.(map[string]interface{})
		if _, ok := zone["id"]; ok {
			id := zone["id"].(string)
			if !backupZone(directory, token, id) {
				// Flag failure, but continue - partial backup is better than nothing
				success = false
			}
		} else {
			clog.Log("error", "Cloudflare zone list is missing an ID field")
			return false
		}
	}
	return success
}

func backupZone(directory string, token string, id string) bool {

	// Create query
	query := "zones/" + id + "/dns_records/export"

	// Query Cloudflare API
	result, err := cfGet(token, query)
	if err != nil {
		clog.Logf("error", "Unable to export DNS records for zone id %s", id)
		return false
	}

	// Write to file
	fileName := directory + string(os.PathSeparator) + id + ".txt"
	err = ioutil.WriteFile(fileName, result, 0644)
	if err != nil {
		clog.Logf("error", "Unable to write zone data to %s", fileName)
		return false
	}
	clog.Logf("info", "Successfully wrote zone data to %s", fileName)
	return true
}

func cfGet(token string, query string) ([]byte, error) {

	// Create and configure HTTP client
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.cloudflare.com/client/v4/"+query, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// Perform request
	response, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	//goland:noinspection GoUnhandledErrorResult
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
