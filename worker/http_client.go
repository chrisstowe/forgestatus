package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/chrisstowe/forgestatus/common"
)

var client = http.Client{Timeout: time.Second}

func getStatus(taskType common.TaskType, workerID int) (string, error) {
	url := fmt.Sprintf(
		"http://forgestatus-worker-%d-service-%s/%s",
		workerID,
		common.EnvConfig.Env,
		taskType,
	)

	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func getStatusForAllWorkers(t common.TaskType) []string {
	results := make([]string, common.EnvConfig.WorkerCount)
	for id := 1; id <= common.EnvConfig.WorkerCount; id++ {
		var status string
		var err error

		// If this is the current worker, then no request is necessary.
		if id == common.EnvConfig.WorkerID {
			status = "22.2"
		} else {
			// Failed requests are simply empty values.
			status, err = getStatus(t, id)
			if err != nil {
				status = ""
			}
		}

		results = append(results, status)
	}

	return results
}
