package git

import (
	"fmt"
	"net/http"
)

func Jenkins(username string,password string,jenurl string,job string,key string) error {
	client := http.Client{}

	URL := jenurl+ "/job/"+ job + "/build?token=" + key
	req, err := http.NewRequest(http.MethodGet, URL, http.NoBody)
	if err != nil {
		return err
	}

	req.SetBasicAuth(username, password)

	resp,err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		return nil
	} else {
		switch resp.StatusCode {
		case http.StatusUnauthorized:
			return fmt.Errorf("authentication failed: invalid credentials")
		case http.StatusForbidden:
			return fmt.Errorf("access denied: you don't have permission to trigger this job")
		case http.StatusNotFound:
			return fmt.Errorf("job not found: check the job name and URL")
		default:
			return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}
	}
}

func JenkinsLog(username string,password string,jenurl string,job string) error {
	client := http.Client{}

	URL := jenurl+ "/job/"+ job + "/lastBuild/consoleText"
	req, err := http.NewRequest(http.MethodGet, URL, http.NoBody)
	if err != nil {
		return err
	}

	req.SetBasicAuth(username, password)

	resp,err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		return nil
	} else {
		switch resp.StatusCode {
		case http.StatusUnauthorized:
			return fmt.Errorf("authentication failed: invalid credentials")
		case http.StatusForbidden:
			return fmt.Errorf("access denied: you don't have permission to trigger this job")
		case http.StatusNotFound:
			return fmt.Errorf("job not found: check the job name and URL")
		default:
			return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}
	}
}

