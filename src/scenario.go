package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Scenarios struct {
	Scenarios []Scenario `json:"scenarios"`
}

type Scenario struct {
	Duration           float64 `json:"duration"`
	Width              uint    `json:"width"`
	Dissemination_rate float64 `json:"dissemination_rate"`
	Lambda             float64 `json:"lambda"`
	Shard_capacity     float64 `json:"shard_capacity"`
	Period             float64 `json:"period"`
}

func import_scenarios() {
	jsonFile, err := os.Open("../scenarios.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened scenarios.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &g_scenario)

	fmt.Println(g_scenario)
}
