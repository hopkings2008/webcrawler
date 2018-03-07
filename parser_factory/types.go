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
	// region|location|class|squery|floorInfo|fireSytem|serviceClass|serviceRegion
	return fmt.Sprintf("%s | %s | %s | %s | %s | %s | %s | %s", whi.Region, whi.Location, whi.Class, whi.Square, whi.FloorInfo, whi.FireSystem, whi.ServiceClass, whi.ServiceRegion)
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
