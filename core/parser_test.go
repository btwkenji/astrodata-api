package core

import (
	"reflect"
	"testing"

	"github.com/nezutero/astrodata-api/vars"
)

func TestDstParser(t *testing.T) {
	data, err := Parser(vars.CurrentMonthDataUrl)

	if err != nil {
		t.Errorf("[!] expected Parser(%s) to be successful, but got an error: %v", vars.Dst7DaysDataUrl, err)
	}

	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() != reflect.Slice || dataValue.Type().Elem().Kind() != reflect.Float64 {
		t.Errorf("[!] expected data to be a slice of float64, but got %T", data)
	}
}


func TestPlasmaTemperatureParser (t *testing.T) {
	data, err := ParsePlasmaTemperatureData(vars.PlasmaTemperature7DaysDataUrl)

	if err != nil {
		t.Errorf("[!] expected Parser(%s) to be successful, but got an error: %v", vars.Dst7DaysDataUrl, err)
	}

	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() != reflect.Slice || dataValue.Type().Elem().Kind() != reflect.Float64 {
		t.Errorf("[!] expected data to be a slice of float64, but got %T", data)
	}
}

func Test7DaysDstParser (t *testing.T) {
	data, err := ParseDst7Days(vars.Dst7DaysDataUrl)

	if err != nil {
		t.Errorf("[!] expected Parser(%s) to be successful, but got an error: %v", vars.Dst7DaysDataUrl, err)
	}

	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() != reflect.Slice || dataValue.Type().Elem().Kind() != reflect.Float64 {
		t.Errorf("[!] expected data to be a slice of float64, but got %T", data)
	}
}


func TestBzParser (t *testing.T) {
	data, err := ParseBzGsmData(vars.RealTimeSolarWindDataSevenDays)

	if err != nil {
		t.Errorf("[!] expected Parser(%s) to be successful, but got an error: %v", vars.Dst7DaysDataUrl, err)
	}

	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() != reflect.Slice || dataValue.Type().Elem().Kind() != reflect.Float64 {
		t.Errorf("[!] expected data to be a slice of float64, but got %T", data)
	}
}
