package utils

import (
	"log"
	"os"
)

var (
	InfoLog    *log.Logger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)
	ErrorLog   *log.Logger = log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLog *log.Logger = log.New(os.Stdout, "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile)
)
