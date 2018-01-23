package main

import (
	"os"
	"time"

	client "github.com/influxdata/influxdb/client/v2"
)

var influxDBURL = os.Getenv("INFLUXURL")

const influxDBDB = "kiosk"

func addPoint(api APIData) error {
	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: influxDBURL,
	})
	if err != nil {
		return err
	}

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "kiosk",
		Precision: "s",
	})
	if err != nil {
		return err
	}

	// Create a point and add to batch
	tags := map[string]string{"setup": api.Setup}
	fields := map[string]interface{}{}

	for _, dataPoint := range api.Data {
		fields[dataPoint.Sensor] = dataPoint.Data
	}

	pt, err := client.NewPoint(api.Setup, tags, fields, time.Now())
	if err != nil {
		return err
	}
	bp.AddPoint(pt)

	return c.Write(bp)
}
