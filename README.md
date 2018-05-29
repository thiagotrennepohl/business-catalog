# Business-Catalog [![Build Status](https://travis-ci.org/thiagotrennepohl/business-catalog.svg?branch=feature-company-catalog)](https://travis-ci.org/thiagotrennepohl/business-catalog)

Simple api for sending data via csv file


## Building/Testing/Docker Build
----
`make all`

## First step - Populate database

_The databses test will run on localhost:27017 and for auditing porpuses it isn't removed after the tests._

_You can check all data at `mongodb://localhost:27017/yawoen`_

_If you run the command `make all` the following steps will be executed_ :<br>
1 - Create a mongodb
2 - Populate it using a csvfile at `assets/q1_catalog.csv`
3 - Run tests
4 - Build
5 - Docker build

If you want to run only the step 1 and 2 <br>
`make populate-db`

And a populated mongodb will be avaliable at localhost:27017

----
If you want to check all options you can just run `make` and a help will be shown on the console.

<br>

# API

**Update Companies**  - POST /v1/company<br><br>
Update companies look into mongoDB if a company exists, if exists the company will be updated
----

* **URL**

  _http://localhost:8080/v1/company_

* **Method:**
  
  _`POST`_

  
<!-- *  **URL Params**

   <_If URL params exist, specify them in accordance with name mentioned in URL section. Separate into optional and required. Document data constraints._> 

   **Required:**
 
   `id=[integer]`

   **Optional:**
 
   `photo_id=[alphanumeric]` -->

* **Data Params**

  _**form-data**_ <br>
  _**fieldname**_ : "data" <br>
  _**fieldvalue**_: .csv file<br>

  _**Recomended Format**_<br>
  ```
    name;addresszip;website
    company;12345;company.com
  ```

* **Required Headers**

  **Content-Type**: "multipart/form-data"

* **Success Response:**


  * **Code:** 200 <br />
    **Content:** `{ message : "ok" }`
 
* **Sample Call:**

  ```curl -X POST \
  http://localhost:8080/v1/company \
  -H 'Content-Type: multipart/form-data' \
  -F data=@/path/to/csv/file.csv```

----

**Find Company**  - GET /v1/company?name=&zip=<br><br>
Find Company will look for a matching company by given a name and zip.
----

* **URL**

  _http://localhost:8080/v1/company_

* **Method:**
  
  _`GET`_

  
*  **URL Params**

   **Required:**
 
   `name=[string]`

   **Optional:**
 
   `zip=[integer]`


* **Success Response:**


  * **Code:** 200 <br />
    **Content:**
     ```
    [
      {
          "ID": "5b0c74e2c8a66af33d035896",
          "Name": "TOLA SALES GROUP",
          "AddressZip": 78229,
          "Website": "http://repsources.com"
      }
    ]
    ```
 
* **Sample Call:**

  `curl -X GET 'http://localhost:8080/v1/company?name=t&zip=78229'`



## Lib for reading data
---

https://github.com/thiagotrennepohl/sdr


## Todo
----
- [] Use Goroutines for parsing the csvfile
- [] Use Goroutines for bulk insertion
- [] Split the csv file into the number of workers(go routines) specified
- [] Cover error returns
- [] Mock mongodb so we don't need a real mongodb anymore
- [] Improve documentation
- [] Include Lint into the build (Gometalinter)
- [] Include Pprof analisys into the build
- [] Improve application design
- [] Create Tests for SDR
