package storage

import (
	"fmt"
	"os"
	"strconv"
)

func GetValue(filename string) float64 {
	dat, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("unable to read from storage")
		return 0
	}
	floatVal, err := strconv.ParseFloat(string(dat), 64)
	if err != nil {
		return 0
	}
	return floatVal
}

func SetValue(filename string, floatVal float64) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("%f", floatVal))
	if err != nil {
		return err
	}
	return nil
}
