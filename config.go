package main

type Config struct {
	prev string
	next string
	step int
}

func getConfig() Config {
	return Config{
		prev: "0",
		next: "1",
		step: 0,
	}
}
