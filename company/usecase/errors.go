package usecase

import "errors"

//ErrorInvalidCSV is the default error for any invalid field in a csv file
var ErrorInvalidCSV = errors.New("The csv file has invalid fields")

var ErrorCouldNotParseZipCode = errors.New("Could not parse a zipCode")
