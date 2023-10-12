package core

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

func Parser(url string) ([]float64, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP request failed with status code %d", response.StatusCode)
	}

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	var result []float64

	document.Find("pre.data").Each(func(index int, element *goquery.Selection) {
		data := element.Text()

		values := strings.Fields(data)
		for _, value := range values {
			num, err := strconv.ParseFloat(value, 64)
			if err == nil {
				if num > 0 {
					num = 0
				}
				result = append(result, num)
			}
		}
	})

	if len(result) > 30 {
		result = result[30:]
	}

	return result, nil
}

func ParseBzGsmData(url string) ([]float64, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code %d", response.StatusCode)
	}

	var rawData [][]interface{}
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&rawData); err != nil {
		return nil, err
	}

	var bzGsmValues []float64
	for _, data := range rawData {
		if len(data) >= 4 {
			if bzGsmStr, ok := data[3].(string); ok {
				bzGsm, err := strconv.ParseFloat(bzGsmStr, 64)
				if err == nil {
					bzGsmValues = append(bzGsmValues, bzGsm)
				}
			}
		}
	}

	return bzGsmValues, nil
}

func ParseDst7Days(url string) ([]float64, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code %d", response.StatusCode)
	}

	var rawData [][]interface{}
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&rawData); err != nil {
		return nil, err
	}

	var dstWeekValues []float64
	for _, data := range rawData {
		if len(data) >= 1 {
			if dstWeekValuesStr, ok := data[1].(string); ok {
				dst, err := strconv.ParseFloat(dstWeekValuesStr, 64)
				if err == nil {
					dstWeekValues = append(dstWeekValues, dst)
				}
			}
		}
	}

	return dstWeekValues, nil
}

func ParsePlasmaTemperatureData(url string) ([]float64, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status code %d", response.StatusCode)
	}

	var rawData [][]interface{}
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&rawData); err != nil {
		return nil, err
	}

	var plasmaData []float64
	for _, data := range rawData {
		if len(data) >= 3 {
			if plasmaDataStr, ok := data[2].(string); ok {
				plasma, err := strconv.ParseFloat(plasmaDataStr, 64)
				if err == nil {
					plasmaData = append(plasmaData, plasma)
				}
			}
		}
	}

	return plasmaData, nil
}
