package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func GetElevation(lat, lon float64) int16 {
	var elevation int16

	lat_deg := int32(lat)
	lon_deg := int32(lon)

	lat_sec := 3600 - int64((lat-float64(lat_deg))*3600)
	lon_sec := int64((lon - float64(lon_deg)) * 3600)

	offset := (lat_sec * 3601 * 2) + (lon_sec * 2)

	fmt.Printf("offset:%d\n", offset)

	fmt.Printf("lat:%d lon:%d\n", lat_deg, lon_deg)
	fmt.Printf("lat_sec:%d lon_sec:%d\n", lat_sec, lon_sec)

	fname := fmt.Sprintf("N%02dE%03d.hgt", lat_deg, lon_deg)
	fmt.Println(fname)

	file, err := os.Open(fname)

	if err != nil {
		return -32768
	}

	defer file.Close()

	buf := make([]byte, 2)

	_, err = file.ReadAt(buf, offset)

	if err != nil {
		fmt.Println(err)
	}

	elevation = (int16(buf[1])) | (int16(buf[0]) << 8)

	//fmt.Println(buf)
	fmt.Println(elevation)

	return elevation
}

var elevation int16
var file *os.File

var dizi []byte

func init() {
	f, err := os.Open("N40E027.hgt")
	if err != nil {
		fmt.Println(err)
	}

	file = f

	dizi, _ = ioutil.ReadAll(f)
}

func GetElevationUltraFast(lat, lon float64) int16 {

	lat_deg := int32(lat)
	lon_deg := int32(lon)

	lat_sec := 3600 - int64((lat-float64(lat_deg))*3600)
	lon_sec := int64((lon - float64(lon_deg)) * 3600)

	offset := (lat_sec * 3601 * 2) + (lon_sec * 2)

	//fmt.Printf("offset:%d\n", offset)

	//fmt.Printf("lat:%d lon:%d\n", lat_deg, lon_deg)
	//fmt.Printf("lat_sec:%d lon_sec:%d\n", lat_sec, lon_sec)

	//fname := fmt.Sprintf("N%02dE%03d.hgt", lat_deg, lon_deg)
	//fmt.Println(fname)

	//elevation = (int16(dizi[offset+1])) | (int16(dizi[offset]) << 8)
	return (int16(dizi[offset+1])) | (int16(dizi[offset]) << 8)

	//fmt.Println(buf)
	//fmt.Println(elevation)

	//return elevation
}

func ElevationAt(x, y int) int16 {

	lat_sec := 3600 - y
	lon_sec := x

	offset := (lat_sec * 3601 * 2) + (lon_sec * 2)

	//elevation = (int16(dizi[offset+1])) | (int16(dizi[offset]) << 8)
	return (int16(dizi[offset+1])) | (int16(dizi[offset]) << 8)
}

func GetElevationFast(lat, lon float64) int16 {

	lat_deg := int32(lat)
	lon_deg := int32(lon)

	lat_sec := 3600 - int64((lat-float64(lat_deg))*3600)
	lon_sec := int64((lon - float64(lon_deg)) * 3600)

	offset := (lat_sec * 3601 * 2) + (lon_sec * 2)

	//fmt.Printf("offset:%d\n", offset)

	//fmt.Printf("lat:%d lon:%d\n", lat_deg, lon_deg)
	//fmt.Printf("lat_sec:%d lon_sec:%d\n", lat_sec, lon_sec)

	//fname := fmt.Sprintf("N%02dE%03d.hgt", lat_deg, lon_deg)
	//fmt.Println(fname)

	buf := make([]byte, 2)

	_, err := file.ReadAt(buf, offset)

	if err != nil {
		fmt.Println(err)
	}

	elevation = (int16(buf[1])) | (int16(buf[0]) << 8)

	//fmt.Println(buf)
	//fmt.Println(elevation)

	return elevation
}

func SpeedTest() {
	r := rand.New(rand.NewSource(99))

	start := time.Now()

	for i := 0; i < 200; i++ {
		GetElevationUltraFast(r.Float64()*40, r.Float64()*27)
	}
	fmt.Printf("Took %s\n", time.Since(start))
}

func main() {

	//fmt.Println(len(dizi))
	//SpeedTest()

	start := time.Now()
	points := Bresenham(0, 0, 3600, 3600)
	//fmt.Println(points)

	for _, p := range points {
		ElevationAt(p.X, p.Y)
	}
	fmt.Println(time.Since(start))

}
