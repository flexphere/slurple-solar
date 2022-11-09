package main

import (
	"flag"

	db "github.com/flexphere/slurple-solar/repository/mongo"
	"github.com/flexphere/slurple-solar/service/batch"

	_ "github.com/joho/godotenv/autoload"
)

var runDuration = flag.Int("d", 0, "days to fetch")

var (
	solarRepository = &db.SolarRepositoryImpl{}
	solarService    = batch.New(solarRepository)
)

func main() {
	solarRepository.Connect()
	defer solarRepository.Disconnect()

	flag.Parse()

	if *runDuration == 0 {
		solarService.Today()
	} else {
		solarService.Duration(*runDuration)
	}
}
