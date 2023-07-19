package tallyManager

import (
	"encoding/json"
	"io/ioutil"
	"logger"
	"os"
)

func ChangeCam(camId int, layerId int) {
	file, err := os.Open(filename)
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
		tally.Layer1 = camId
	} else if layerId == 2 {
		tally.Layer2 = camId
	} else if layerId == 3 {
		tally.Layer3 = camId
	} else {
		logger.ErrorLogger.Println("Error changing cam, layerId not valid")
	}

	tallyJSON, err = json.Marshal(tally)
	if err != nil {
		logger.ErrorLogger.Println("Error marshalling tally")
	}

	err = ioutil.WriteFile(filename, tallyJSON, 0644)
	if err != nil {
		logger.ErrorLogger.Println("Error writing tally to file")
	}
}

func RemoveCamIdOfAllLayer(CamId int, LayerIdNotTouch int) {
	file, err := os.Open(filename)
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

	if LayerIdNotTouch != 1 {
		if tally.Layer1 == CamId {
			ChangeCam(0, 1)
		}
	} else if LayerIdNotTouch != 2 {
		if tally.Layer2 == CamId {
			ChangeCam(0, 2)
		}
	}
	if LayerIdNotTouch != 3 {
		if tally.Layer3 == CamId {
			ChangeCam(0, 3)
		}
	}
}
