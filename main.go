package main

import (
	"flag"
	"fmt"
	"github.com/Knetic/govaluate"
	"log"
	"strconv"
	"strings"
)

//type taskRecord struct {
//	level string
//	mission	string
//	result float32
//	id int16
//}

//type mongoInfo struct {
//	host string
//	port string
//	db string
//	rs string
//}

//var mongoDSN mongoInfo
type task struct {
	template string
	factor   int
	divider  int
	a        int
	b        int
	c        int
	iter     int
	level    int
	result   interface{}
	mission  string
}

var taskInfo task
var parameters = make(map[string]interface{}, 8)

func init() {
	//flag.StringVar(&mongoDSN.host,"mongo.host","127.0.0.1","MongoDB host")
	//flag.StringVar(&mongoDSN.port,"mongo.port","27017","MongoDB port")
	//flag.StringVar(&mongoDSN.db,"mongo.db","kidscalc","MongoDB db name")
	//flag.StringVar(&mongoDSN.rs,"mongo.rs","rs0","MongoDB replica set")
	flag.StringVar(&taskInfo.template, "task.template", "", "Task template. Example a+m*(b-c)")
	flag.IntVar(&taskInfo.level, "task.level", 1, "Task level")
	flag.IntVar(&taskInfo.iter, "task.iter", 0, "Start task iteration number from")
	flag.IntVar(&taskInfo.factor, "task.m", 1, "m - multiplier in template")
	flag.IntVar(&taskInfo.divider, "task.d", 1, "d - divider in template")
	flag.IntVar(&taskInfo.a, "task.a", 1, "parameter a")
	flag.IntVar(&taskInfo.b, "task.b", 1, "parameter b")
	flag.IntVar(&taskInfo.c, "task.c", 1, "parameter c")
	flag.Parse()
}

func renderTemplate(params map[string]interface{}) string {
	var r *strings.Replacer
	// TODO: replace switch with if statement
	switch (params["m"] == 1) && (params["d"] == 1) {
	case true:
		r = strings.NewReplacer(
			"a", strconv.Itoa(params["a"].(int)),
			"b", strconv.Itoa(params["b"].(int)),
			"c", strconv.Itoa(params["c"].(int)),
			"m*", "",
			"*m", "",
			"/d", "",
			"d/", "",
		)
	default:
		r = strings.NewReplacer(
			"a", strconv.Itoa(params["a"].(int)),
			"b", strconv.Itoa(params["b"].(int)),
			"c", strconv.Itoa(params["c"].(int)),
			"m", strconv.Itoa(params["m"].(int)),
			"d", strconv.Itoa(params["d"].(int)),
		)
	}
	return r.Replace(taskInfo.template)
}

func main() {
	expression, err := govaluate.NewEvaluableExpression(taskInfo.template)

	if err != nil {
		log.Printf("Error: %s", err)
	}

	parameters["iter"] = taskInfo.iter
	for d := 1; d <= taskInfo.divider; d++ {
		parameters["d"] = d
		for m := 1; m <= taskInfo.factor; m++ {
			parameters["m"] = m
			for a := 1; a <= taskInfo.a; a++ {
				parameters["a"] = a
				for b := 1; b <= taskInfo.b; b++ {
					parameters["b"] = b
					for c := 1; c <= taskInfo.c; c++ {
						parameters["c"] = c
						result, err := expression.Evaluate(parameters)
						if err != nil {
							log.Printf("Error: %s", err)
						}
						taskInfo.result = result
						taskInfo.mission = renderTemplate(parameters)
						parameters["iter"] = parameters["iter"].(int) + 1
						fmt.Printf("%d\t%s=%v\n", parameters["iter"], taskInfo.mission, result)
					}
				}
			}
		}
	}
}
