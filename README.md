#### API document

- `GET /api/patient/japanese/location`

  ```json
  { 
   "code": 200,
   "msg": "OK",
   "data": [
     {"Location": "北海道",
      "PatientSum": 100
     }, ...
   ]
  }
  ```

- `GET /api/patient/news?number=5`

  ```json
  { 
   "code": 200,
   "msg": "OK",
   "data": [
     {"Title": "corona virus",
      "Url": "www.google,com"
     }, ...
   ]
  }
  ```

- `GET /api/patient/japanese/summary`

  ```json
  { 
   "code": 200,
   "msg": "OK",
   "data": {
    "PatientSum": 200,
    "NewPatient": {
    	"Sum": 10,
      "Growth": 12, 
  	},
    "DeadPatient": {
    	"Sum": 20,
      "Growth": -5, 
    }
    "UpdatedDate": "2020年3月7日"
  }
  ```

- `GET /api/patient/summary`

  ```json
  { 
   "code": 200,
   "msg": "OK",
   "data": {
    "GlobalSum": 1000000,
    "ShipSum": 1000
   }
  }
  ```

  