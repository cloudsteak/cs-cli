package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const PathToTeamsResponsibilitiesYaml = "responsibilities/teams.yaml"

func SetLogFlags() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)
}

func ComputeSeverity(cvss3Score string) string {
	score, err := strconv.ParseFloat(cvss3Score, 64)
	if err != nil {
		return "undefined"
	}
	if score >= 9.0 {
		return "critical"
	} else if score >= 7.0 {
		return "high"
	} else if score >= 4.0 {
		return "medium"
	}
	return "low"
}

func GetTargetDays(severity string) int {
	if strings.EqualFold(severity, "critical") || strings.EqualFold(severity, "high") {
		return 30
	} else if strings.EqualFold(severity, "medium") {
		return 90
	} else if strings.EqualFold(severity, "low") {
		return 120
	}
	return 30
}

func GenerateReq(token string, url string) (*http.Request, error) {
	var bearer = "Bearer " + token
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)
	return req, err
}

func SendReq(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error on response %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the response bytes %w", err)
	}
	return body, err
}

func ConvertToJSON(input any) []byte {
	outputJSON, err := json.MarshalIndent(input, "", "    ")
	if err != nil {
		log.Fatalf("Failed to convert map to the JSON bytes: %s", err)
	}
	return outputJSON
}

func WriteToFile(name string, content []byte) {
	dir := filepath.Dir(name)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create directory %s. %s", dir, err)
	}
	err = os.WriteFile(name, content, fs.ModePerm)
	if err != nil {
		log.Fatalf("Failed to write bytes to the file: %s", err)
	}
}

func RemoveDirectory(dirName string) error {
	// Delete the directory
	err := os.RemoveAll(dirName)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	return nil
}
