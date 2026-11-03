package cmdmanager

import (
	"errors"
	"fmt"
)

type CMDManager struct{}

func (cmdm CMDManager) ReadLines() ([]string, error) {
	var prices []string
	var price string

	fmt.Print("Type 'stop' to exit price reading.")

	for {
		fmt.Print("Price: ")
		_, err := fmt.Scan(&price)

		if err != nil {
			return []string{}, errors.New("Failed to read read value")
		}

		if price == "stop" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}

func (cmdm CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}
