package main

import (
	"flag"
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/lego/log"
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
	a        int
	b        int
	c        int
}

var taskInfo task

func init() {
	//flag.StringVar(&mongoDSN.host,"mongo.host","127.0.0.1","MongoDB host")
	//flag.StringVar(&mongoDSN.port,"mongo.port","27017","MongoDB port")
	//flag.StringVar(&mongoDSN.db,"mongo.db","kidscalc","MongoDB db name")
	//flag.StringVar(&mongoDSN.rs,"mongo.rs","rs0","MongoDB replica set")
	flag.StringVar(&taskInfo.template, "task.template", "", "Task template")
	flag.IntVar(&taskInfo.factor, "task.factor", 1, "m - multiplier in template")
	flag.IntVar(&taskInfo.a, "task.a", 1, "parameter a")
	flag.IntVar(&taskInfo.b, "task.b", 1, "parameter b")
	flag.IntVar(&taskInfo.c, "task.c", 1, "parameter c")
	flag.Parse()
}

func renderTemplate(params map[string]interface{}) string {
	var r *strings.Replacer
	if taskInfo.factor == 1 {
		r = strings.NewReplacer("a", strconv.Itoa(params["a"].(int)),
			"m*", "",
			"b", strconv.Itoa(params["b"].(int)),
			"c", strconv.Itoa(params["c"].(int)))
	} else {
		r = strings.NewReplacer("a", strconv.Itoa(params["a"].(int)),
			"m", strconv.Itoa(params["m"].(int)),
			"b", strconv.Itoa(params["b"].(int)),
			"c", strconv.Itoa(params["c"].(int)))
	}
	//fmt.Printf("a=%v, b=%v, c=%v\n",strconv.Itoa(params["a"].(int)),strconv.Itoa(params["b"].(int)),strconv.Itoa(params["c"].(int)))
	return r.Replace(taskInfo.template)
}

func main() {
	expression, err := govaluate.NewEvaluableExpression(taskInfo.template)

	if err != nil {
		log.Printf("Error: %s", err)
	}

	parameters := make(map[string]interface{}, 8)
	parameters["iter"] = 0
	for parameters["m"] = range make([]int, taskInfo.factor) {
		for parameters["a"] = range make([]int, taskInfo.a) {
			for parameters["b"] = range make([]int, taskInfo.b) {
				for parameters["c"] = range make([]int, taskInfo.c) {
					result, err := expression.Evaluate(parameters)
					if err != nil {
						log.Printf("Error: %s", err)
					}
					parameters["iter"] = parameters["iter"].(int) + 1
					fmt.Printf("%d\t%s=%v\n", parameters["iter"], renderTemplate(parameters), result)
				}
			}
		}
	}
}
