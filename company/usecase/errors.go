package usecase

import "errors"

//ErrorInvalidCSV is the default error for any invalid field in a csv file
var ErrorInvalidCSV = errors.New("The csv file has invalid fields")

//ErrorCouldNotParseZipCode is the default error when a zipCode isn't parseable
var ErrorCouldNotParseZipCode = errors.New("Could not parse a zipCode")

//ErrorInvalidHeaders is the default error when csv headers are invalid
var ErrorInvalidHeaders = errors.New("Invalid headers")

//ErrorTooManyHeaders is the default error when a csv file has more headers than it shoulds
var ErrorTooManyHeaders = errors.New("There are more headers than the accepted format")
