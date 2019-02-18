# ForgeStatus

[![CircleCI](https://circleci.com/gh/chrisstowe/forgestatus.svg?style=svg)](https://circleci.com/gh/chrisstowe/forgestatus) [![Go Report Card](https://goreportcard.com/badge/github.com/chrisstowe/forgestatus)](https://goreportcard.com/report/github.com/chrisstowe/forgestatus) [![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/chrisstowe/forgestatus)

## Description

A distributed system status and metric checker. 🔎

A visual dashboard for this service can be found at [forgestatus-dashboard](https://github.com/chrisstowe/forgestatus-dashboard).

### What this project does

- A server queues up tasks for checking system metrics (memory, cpu, etc.).

- Several workers move tasks from this queue into their own queues.

- Each task describes a system metric that a worker needs to query all other workers for.

- The result is queued up once a worker has aggregated all the necessary metrics.

- Workers are also responsible for responding to all other workers requests for info.

- A client can query for metric results at any time.

#### Extras

- Automated CI/CD using [CircleCI](https://circleci.com/) (including gates and efficient artifact caching).
- Full [Kubernetes (GKE)](https://cloud.google.com/kubernetes-engine/) deployment for two environments.
- [Stackdriver](https://cloud.google.com/stackdriver/) alerts/monitoring of system and site health.
- Slack notifications for releases and alerts.

#### Live Builds

These are fully automated and monitored instances.

- [forgestatus.com](http://forgestatus.com)

- [dev.forgestatus.com](http://dev.forgestatus.com)

#### Raw API Data

Used for debugging while work is done on [forgestatus-dashboard](https://github.com/chrisstowe/forgestatus-dashboard).

- [forgestatus.com/api/status](http://forgestatus.com/api/status)

- [dev.forgestatus.com/api/status](http://dev.forgestatus.com/api/status)

## Building

### Docker

This is for demo purposes.

1 server, 3 workers, and 1 redis instance will be started.

Output for task scheduling and processing is printed to the console.

```
$ docker-compose up
```

### Go and Make

This is for local development purposes.

This requires a working [go environment](https://golang.org/doc/code.html)
A locally running instance of redis is also required.

```
$ make
$ REDIS_URL=localhost:6379 server
$ REDIS_URL=localhost:6379 worker
```

### CircleCI

Make a pull request and CircleCI will automatically build, test, and deploy your app to dev.forgestatus.com.

# Questions

## Unit/Integration tests?

There are both unit and integration tests.

```
$ make unit-test
$ make integration-test
```

There is a CI gate that prevents code from being merged to master with failing tests.

Failing tests also trigger a slack alert at [forgestatus.slack.com](forgestatus.slack.com).

As I learn more about the language, more tests will be added.

## Are there any shortcomings of the code?

An intentional shortcoming is that each worker is uniquely identified with a k8s deployment/service.
This was required because each worker needs the capability to query all other workers.
In a real system, there would only be one deployment/service for a single type of worker.

The reliable queue pattern was followed from [redis.io](https://redis.io/commands/rpoplpush).
The only missing piece is the process to go through and re-queue stale work in the pending queues.

Workers only queue up one task at a time.
Workers could potentially grab a handful of tasks and execute them in parallel.
Also, each worker sequentially queries each other worker when given a command (which could easily be done asynchronously).

There definitely needs to be more tests.
Significant time was spent on playing with golang, automation, and deployment

Automation of infrastructure resources could have been done through [Terraform](https://www.terraform.io/).
Currently, all resource commands are saved as script files in [infra](infra).

## How might this project be scaled?

This could be scaled by removing the unique worker k8s deployments/services.

The addition of [Istio](https://istio.io/) as the Ingress would make service management cleaner.

There could also be a hierarchy of task queues.
This would allow a subset of tasks to be completed in parallel.
Concurrency would be managed by parent processes/tasks.
Something like an [event aggregator](https://martinfowler.com/eaaDev/EventAggregator.html) could be used for state management.

## How might one approach doing sequential versus parallel tasks?

### Sequential

This is already supported by having the server only schedule one task at a time.
If another process needed to manage the task result, then a field for identifying dependent tasks could be added.
The process could look up completed tasks and only work on the next task when all dependencies have been met.

### Parallel

Tasks are currently identified with a unique ID.
If the server needs a task to be done in parallel, then multiple tasks could be sent with associated IDs.
If the tasks needed to be aggregated, then something like a part number or aggregate ID could be added to each task.
