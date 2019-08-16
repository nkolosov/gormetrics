# gormetrics

A plugin for gorm providing metrics using Prometheus.

Warning: this plugin is still in an early stage of development. APIs may change.

## Usage

```go
import "github.com/profects/gormetrics"

err := gormetrics.Register(db, <database>)
if err != nil {
	// handle the error
}
```

gormetrics does not expose the metrics endpoint using promhttp, you have to do this yourself.
Otherwise, you can use the following snippet:
```go
go func() {
    http.Handle("/metrics", promhttp.Handler())
    log.Fatal(http.ListenAndServe(":2112", nil))
}()
```

## Exported metrics

gormetrics exports the following metrics (counter vectors):
* gormetrics_all_total
* gormetrics_creates_total
* gormetrics_deletes_total
* gormetrics_queries_total
* gormetrics_updates_total

These all have the following labels:
* database: the name of the database
* driver: the driver for the database (e.g. pq)
* status: fail or success

It also export the following metrics (histogram vectors):
* gormetrics_connections_idle
* gormetrics_connections_in_use
* gormetrics_connections_open

These all have the following labels:
* database: the name of the database
* driver: the driver for the database (e.g. pq)