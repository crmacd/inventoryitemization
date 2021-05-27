package main

import "time"

type Item interface {
	getManufacture() string
	getModel() string
	getOSVersion() string
	getOSStatus() string
	getSerialNumber() string
	getLatitude() float64
	getLongitude() float64
	getLocationName() string
	getReportInTimestamp() time.Time
	getPhysicalCondition() string
	getIPAddress() string
	getMACAddress() string // (Manufacture)
}

type ItemDetails struct {
	Manufacture string
	Model string
	OSVersion string
	OSStatus string
	SerialNumber string
	Latitude float64
	Longitude float64
	LocationName string
	ReportInTimeStamp time.Time
	PhysicalCondition string
	IPAddress string
	MACAddress string
}

func (objItem ItemDetails) getReportInTimestamp() time.Time {
	objItem.ReportInTimeStamp := time.Now()
	return objItem.ReportInTimeStamp

}
