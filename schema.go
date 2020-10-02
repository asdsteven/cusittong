package main

import (
	"time"
)

var thisTerm = [2]string{
	"2000",
	"jj",
}

type reserveS struct {
	major  string
	quota  int
	enroll int
}

type rowHeadS struct {
	row      int
	code     string
	group    string
	nbr      string
	title    string
	units    string
	teachers []string
	reserves []reserveS
	rowBodys []rowBodyS
}

type rowBodyS struct {
	quota     int
	vacancy   int
	component string
	section   string
	language  string
	rowFoots  []rowFootS
	dept      string
}

type rowFootS struct {
	weekday string
	period  []time.Time
	room    string
	date    []time.Time
	add     bool
	drop    bool
}

type changeS struct {
	stamp int64
	index int
	value interface{}
}

type historyS struct {
	group    string
	title    string
	quota    int
	enroll   int
	reserves []reserveS
	changes  []changeS
}

var db map[string]*historyS
