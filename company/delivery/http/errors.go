package http

import "errors"

//ErrorNoSuchFile is the default error return when the app can't write or read a file
var ErrorNoSuchFile = errors.New("No such fle in formData")

//ErrorNoSuchDir is the default  error return when the app can't access a dir
var ErrorNoSuchDir = errors.New("No such directory named assets")
