package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"time"
)

// Values array
type Values struct {
	Name     string  `json:"name"`
	TempMin  float32 `json:"tempMin"`
	TempMax  float32 `json:"tempMax"`
	Interval int     `json:"interval"`
	Values   []Value `json:"values"`
}

// Value struct
type Value struct {
	Message      int     `json:"messageID"`
	Temperature  float32 `json:"temperature"`
	EnqueuedTime string  `json:"enqueuedTime"`
}

// Readings struct
type reading struct {
	hour       int
	normal     float32
	outofrange float32
}

func main() {
	jsonFile, err := os.Open("data.json")
	if err != nil {
		log.Fatal("File not found")
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var v Values

	json.Unmarshal(byteValue, &v)

	tempMap := make(map[int][]float32)
	fmt.Println("Total Sensor Readings: ", len(v.Values))
	for _, e := range v.Values {
		t, err := time.Parse("2006-01-02- 15:04:05", e.EnqueuedTime)
		if err != nil {
			log.Fatal(err)
		}
		h := t.Hour()
		tempMap[h] = append(tempMap[h], e.Temperature)
	}

	var normal, outofrange float32
	var readings []reading

	for x, y := range tempMap {
		normal, outofrange = 0.0, 0.0
		for _, b := range y {
			if b >= v.TempMin && b <= v.TempMax {
				normal++
			} else {
				outofrange++
			}
		}
		read := reading{x, normal, outofrange}
		readings = append(readings, read)
	}

	sort.Slice(readings, func(i, j int) bool {
		return readings[i].hour < readings[j].hour
	})

	printTable(readings)

}

func printTable(r []reading) {
	fmt.Printf("Hour\tTotal\tNormal\tOut of Range\tPercent\n")
	for _, val := range r {
		total := val.normal + val.outofrange
		percent := val.outofrange / total * 100
		fmt.Printf("%v\t%v\t%v\t%5v\t\t%5.1f\n", val.hour, total, val.normal, val.outofrange, percent)
	}
}
