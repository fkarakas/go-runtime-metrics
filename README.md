# Goal

This repository provides a docker-compose for Go runtime metrics stack on your local computer.
It is using influxdb and grafana, it based on the project http://github.com/tevjef/go-runtime-metrics

# Quick run

## 1- Start docker-compose

Start metric stack:

```
make start
```

## 2- Insert metrics Go code

```go
import (
	metrics "github.com/tevjef/go-runtime-metrics"
)

func main() {
	err := metrics.RunCollector(metrics.DefaultConfig)
	
	if err != nil {
	   // handle error
	}
	
	...
}
```

## 3- Access dashboard 'Go runtime'

Wait for metrics collect and go to the dashboard Go runtime. If not selected, choose the right `measurement` in the dashboard top drop down list.

http://localhost:3000/d/1/go-runtime?orgId=1&from=now-15m&to=now&refresh=5s

![](/dashboard.png)

[Download Dashboard](https://grafana.net/dashboards/3242)

## 4- Stop docker-compose

When done, stop the metric stack:
```
make stop
```
