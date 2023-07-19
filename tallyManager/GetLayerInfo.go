package tallyManager

import (
	"encoding/json"
	"io/ioutil"
	"logger"
	"os"
)

func GetLayerInfo(layerId int) int {
	file, err := os.Open("data/tally.json")
	if err != nil {
		logger.ErrorLogger.Println("Error opening tally file")
	}
	defer file.Close()

	tallyJSON, err := ioutil.ReadAll(file)
	if err != nil {
		logger.ErrorLogger.Println("Error reading tally file")
	}

	var tally TallyStruct
	err = json.Unmarshal(tallyJSON, &tally)
	if err != nil {
		logger.ErrorLogger.Println("Error unmarshalling tally")
	}

	if layerId == 1 {
		return tally.Layer1
	} else if layerId == 2 {
		return tally.Layer2
	} else if layerId == 3 {
		return tally.Layer3
	} else {
		logger.ErrorLogger.Println("Error getting layer info, layerId not valid")
	}
	return 0
}
