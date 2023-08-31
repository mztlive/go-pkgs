package coordinate

import (
	"sort"
	"sync"

	"github.com/mztlive/go-pkgs/math"
	"github.com/samber/lo"
)

type CityDistance struct {
	City     *City
	Distance float64
}

// NearestCity returns the city closest to the input latitude and longitude
func NearestCity(lon, lat float64) City {
	// 将cities拆成10个数组
	cities := GetCities()
	chunks := lo.Chunk(cities, 70)
	wg := sync.WaitGroup{}
	wg.Add(len(chunks))
	distances := make([]*CityDistance, 0)
	results := make(chan *CityDistance, len(chunks))

	for _, chunk := range chunks {
		go func(items []City) {
			defer wg.Done()
			results <- calcDistance(lon, lat, items)
		}(chunk)
	}
	wg.Wait()
	close(results)

	for result := range results {
		distances = append(distances, result)
	}

	// 再次排序取最小值
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})

	return *distances[0].City
}

// calcDistance returns the city closest to the input latitude and longitude
func calcDistance(lon, lat float64, cities []City) *CityDistance {

	distances := make([]CityDistance, len(cities))
	for i, city := range cities {
		distance := math.Distance(lon, lat, city.Lon, city.Lat)
		distances[i] = CityDistance{
			City:     &cities[i],
			Distance: distance,
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})

	return &distances[0]
}
