package parser_factory

import (
	"fmt"
)

type WarehouseInfo struct {
	IsValid       int
	Region        string
	Location      string
	Class         string
	Square        string
	FloorInfo     string
	FireSystem    string
	Industry      string
	ServiceClass  string
	ServiceRegion string
}

func (whi *WarehouseInfo) String() string {
	return fmt.Sprintf("region: %s | location: %s | class: %s | square: %s | floorInfo: %s | fireSystem: %s | serviceClass: %s | serviceRegion: %s", whi.Region, whi.Location, whi.Class, whi.Square, whi.FloorInfo, whi.FireSystem, whi.ServiceClass, whi.ServiceRegion)
}

func CreateWarehouseInfo() *WarehouseInfo {
	wi := &WarehouseInfo{
		IsValid:    0,
		Region:     "",
		Location:   "",
		Class:      "",
		Square:     "",
		FloorInfo:  "",
		FireSystem: "",
	}
	return wi
}
