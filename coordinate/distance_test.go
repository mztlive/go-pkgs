package coordinate

import (
	"testing"
)

func TestNearestCity(t *testing.T) {
	lon, lat := 113.43441, 23.135518
	Initiation()
	city := NearestCity(lon, lat)
	if city.Name != "广州市" {
		t.Errorf("NearestCity() = %v, want %v", city.Name, "广州市")
	}
}

func BenchmarkNearestCity(b *testing.B) {
	lon, lat := 113.43441, 23.135518
	Initiation()
	for i := 0; i < b.N; i++ {
		NearestCity(lat, lon)
	}
}

func TestCalcDistance(t *testing.T) {
	lon, lat := 113.43441, 23.135518
	Initiation()
	city := calcDistance(lon, lat, GetCities())
	if city.City.Name != "广州市" {
		t.Errorf("calcDistance() = %v, want %v", city.City.Name, "广州市")
	}
}

func BenchmarkCalcDistance(b *testing.B) {
	lon, lat := 113.43441, 23.135518
	Initiation()
	cities := GetCities()
	for i := 0; i < b.N; i++ {
		calcDistance(lat, lon, cities)
	}
}
