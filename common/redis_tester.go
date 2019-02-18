package common

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/go-redis/redis"
)

const (
	// TestRedisURL is the locations for the redis test database.
	TestRedisURL     = "localhost:6379"
	redisImage       = "redis:5.0-alpine"
	redisPortDefault = "6379"
)

// RedisTester is a client for a redis test database.
type RedisTester struct {
	Client      *redis.Client
	containerID string
}

// NewRedisTester creates and starts a redis test database.
func NewRedisTester() (*RedisTester, error) {
	portMap := redisPortDefault + ":" + redisPortDefault
	dockerArgs := []string{"run", "-d", "-p", portMap, redisImage}

	out, err := exec.Command("docker", dockerArgs...).CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("docker run: %v: %s", err, out)
	}

	containerID := strings.TrimSpace(string(out))
	client := redis.NewClient(&redis.Options{
		Addr: TestRedisURL,
	})

	tester := &RedisTester{
		Client:      client,
		containerID: containerID,
	}

	return tester, nil
}

// Close stops the redis test database.
func (tester *RedisTester) Close() error {
	out, err := exec.Command("docker", "rm", "-f", tester.containerID).CombinedOutput()
	if err != nil {
		return fmt.Errorf("docker rm: %v: %s", err, out)
	}

	return nil
}
