# Goal

This repository provides a docker-compose for Go runtime metrics stack on your local computer.
It is using influxdb and grafana, it based on the project github.com/tevjef/go-runtime-metrics

# Quick run

## 1- start docker-compose

Start metric stack:

```
make start
```

## 2- insert metrics code

```
import (
	metrics "github.com/tevjef/go-runtime-metrics"
)

func main() {
	err := metrics.RunCollector(metrics.DefaultConfig)
	
	if err != nil {
	   // handle error
	}
}
```

## 3- access dashboard 'Go runtime'

wait for metrics and go to the dashboard Go runtime. If not selected, choose the right `measurement` in the dashboard top drop down list.

http://localhost:3000/d/1/go-runtime?orgId=1&from=now-15m&to=now&refresh=5s

![](/dashboard.png)

[Download Dashboard](https://grafana.net/dashboards/3242)

## 4- stop docker-compose

When done, stop the metric stack:
```
make stop
```
