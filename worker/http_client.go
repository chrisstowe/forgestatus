package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/chrisstowe/forgestatus/common"
)

var client = http.Client{Timeout: 5 * time.Second}

func getStatus(taskType common.TaskType, workerID string) (string, error) {
	url := fmt.Sprintf(
		"http://forgestatus-worker-%s-service-%s/%s",
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
