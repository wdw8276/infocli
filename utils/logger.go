package utils

import (
	"log"
	"os"
)

var ( // global
	// Logger = log.New(os.Stdout, "", log.Ldate | log.Ltime | log.Lshortfile)
	Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
)

// // NOTE: Usage
// var (
// 	logger = utils.Logger
// )
// logger.Println("abc")
// logger.Printf("abc")
// logger.Fatalln("abc")
// logger.Fatalf("abc")
