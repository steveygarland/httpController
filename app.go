package main

import "time"

type App struct {
	Id		int		`json:"id"`
	Command		string		`json:"command"`
	Arg1		string		`json:"arg1"`
	Arg2		string		`json:"arg2,omitempty"`
	Cwd		string		`json:"cwd"`
	State		string		`json:"state"`
	TimeStarted	time.Time	`json:"time"`
}

type Apps []App
