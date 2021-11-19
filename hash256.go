package main

import (
	"encoding/hex"
	"fmt"
	"github.com/minio/highwayhash"
	"io"
	"os"
)

func Hash256(key []byte, path string) (string, error) {
	hash, err := highwayhash.New(key)
	if err != nil {
		fmt.Printf("Failed to create HighwayHash instance: %v", err) // add error handling
		return "", err
	}

	file, err := os.Open(path) // specify your file here
	if err != nil {
		fmt.Printf("Failed to open the file: %v", err) // add error handling
		return "", err
	}
	defer file.Close()

	if _, err = io.Copy(hash, file); err != nil {
		fmt.Printf("Failed to read from file: %v", err) // add error handling
		return "", err
	}

	checksum := hash.Sum(nil)
	return hex.EncodeToString(checksum), nil
}
