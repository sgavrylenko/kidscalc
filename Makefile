build:
	go build -o kidscalc

run:
	./kidscalc -task.template "a+m*(b-c)" -task.a 10 -task.b 10 -task.c 10