package coordinate

import (
	"bufio"
	"bytes"
	"embed"
	_ "embed"
	"strings"
	"sync"

	"github.com/spf13/cast"
)

// City is the city struct
type City struct {
	SN   string
	Name string
	Lon  float64
	Lat  float64
}

var (
	cities []City
	rwLock sync.RWMutex
)

//go:embed city_coordinate.csv
var csv embed.FS

// Initiation Cities
func Initiation() {
	rwLock.Lock()
	fBytes, err := csv.ReadFile("city_coordinate.csv")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(fBytes))
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, ",")
		if len(splits) != 4 {
			continue
		}

		city := City{
			SN:   splits[0],
			Name: splits[1],
			Lon:  cast.ToFloat64(splits[2]),
			Lat:  cast.ToFloat64(splits[3]),
		}

		cities = append(cities, city)
	}
	rwLock.Unlock()
}

// GetCities returns all cities
func GetCities() []City {
	rwLock.RLock()
	defer rwLock.RUnlock()
	return cities
}
