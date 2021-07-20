package main

import "fmt"

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

func main() {
	fmt.Println("IoT Sensor Analysis")
}
