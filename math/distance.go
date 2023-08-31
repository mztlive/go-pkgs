package math

import (
	g_math "math"
)

const (
	// 地球半径
	EarthRadius = 6371
)

// Distance returns the distance between two points
func Distance(lng1, lat1, lng2, lat2 float64) float64 {
	// 将经纬度转换为弧度
	lat1Rad := degToRad(lat1)
	lng1Rad := degToRad(lng1)
	lat2Rad := degToRad(lat2)
	lng2Rad := degToRad(lng2)

	// 计算两点之间的大圆距离
	deltaLng := g_math.Abs(lng1Rad - lng2Rad)
	centralAngle := g_math.Acos(g_math.Sin(lat1Rad)*g_math.Sin(lat2Rad) + g_math.Cos(lat1Rad)*g_math.Cos(lat2Rad)*g_math.Cos(deltaLng))

	// 计算两点之间的距离
	distance := EarthRadius * centralAngle

	return distance
}

// 将角度转换为弧度
func degToRad(deg float64) float64 {
	return deg * g_math.Pi / 180
}
