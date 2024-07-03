package action

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func splitHeader(header string) ([]string, error) {
	header = strings.ReplaceAll(header, " ", "")
	if len(header) == 0 {
		return nil, fmt.Errorf("header values must be provided")
	}
	headerValues := strings.Split(header, ",")
	return headerValues, nil
}

func convertCSVtoJSON(header []string, inputPath, separator string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("error while opening file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = []rune(separator)[0] 
	reader.LazyQuotes = true

	rows, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error while reading .csv: %v", err)
	}

	if len(rows) == 0 || !equalSlice(rows[0], header) {
		return fmt.Errorf("mismatched CSV header")
	}

	var records []map[string]interface{}

	for _, row := range rows[1:] { 
		record := make(map[string]interface{})
		for i, value := range row {
			typedValue, err := detectType(value)
			if err != nil {
				return fmt.Errorf("error detecting type for value '%s': %v", value, err)
			}
			record[header[i]] = typedValue
		}
		records = append(records, record)
	}

	jsonData, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling to JSON: %v", err)
	}

	outputPath := strings.TrimSuffix(inputPath, ".csv") + ".json"

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer outputFile.Close()

	_, err = outputFile.Write(jsonData)
	if err != nil {
		return fmt.Errorf("error writing JSON data to file: %v", err)
	}

	fmt.Printf("CSV converted to JSON and saved to %s\n", outputPath)
	return nil
}

func detectType(value string) (interface{}, error) {
	if boolValue, err := strconv.ParseBool(value); err == nil {
		return boolValue, nil
	}

	if intValue, err := strconv.Atoi(value); err == nil {
		return intValue, nil
	}

	return value, nil
}

func equalSlice(a, b []string) bool {
	return slices.Equal(a, b)
}

func Convert(header, inputPath, separator string) error {
	_header, err := splitHeader(header)
	if err != nil {
        return err
    }
	convertCSVtoJSON(_header, inputPath, separator)
	return nil
}