package tallyManager

import (
	"encoding/json"
	"io"
	"logger"
	"os"
)

func GetLayersInfo() []int {
	file, err := os.Open(filename)
	if err != nil {
		logger.ErrorLogger.Println("Error opening tally file")
	}
	defer file.Close()

	tallyJSON, err := io.ReadAll(file)
	if err != nil {
		logger.ErrorLogger.Println("Error reading tally file")
	}

	var tally TallyStruct
	err = json.Unmarshal(tallyJSON, &tally)
	if err != nil {
		logger.ErrorLogger.Println("Error unmarshalling tally")
	}

	return []int{tally.Layer1, tally.Layer2, tally.Layer3}
}