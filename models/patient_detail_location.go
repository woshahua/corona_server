package models

type PatientDetailLocation struct {
	//gorm.Model
	Latitude      float64 `json: "latitude"`
	Longitude     float64 `json: "longitude"`
	GeoHash       string  `json: "geo_hash"`
	Address       string  `json: "address"`
}

func InsertPatientDetailLocation(patientDetailLocation *PatientDetailLocation) error {
	location := PatientDetailLocation{}

	var notExist = db.Where("patient_detail_location.geo_hash = ?", patientDetailLocation.GeoHash).First(&location).RecordNotFound()

	if notExist {
		err := db.Create(&patientDetailLocation).Error
		return err
	}
	return nil
}

func FindByAddress(address string) *PatientDetailLocation {
	location := PatientDetailLocation{}

	if err := db.Where("patient_detail_location.address = ?", address).First(&location).Error; err != nil {
		return nil
	}
	return &location
}