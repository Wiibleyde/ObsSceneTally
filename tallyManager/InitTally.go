package tallyManager

import (
	"encoding/json"
	"logger"

	"os"
)

const (
	filename = "data/tally.json"
)

type TallyStruct struct {
	Layer1 int
	Layer2 int
	Layer3 int
}

func InitTally() {
	if _, err := os.Stat("data"); os.IsNotExist(err) {
		err := os.Mkdir("data", 0755)
		if err != nil {
			logger.ErrorLogger.Println("Error creating data folder")
		}
	}
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			logger.ErrorLogger.Println("Error creating tally file")
		}
		defer file.Close()

		tally := TallyStruct{0, 0, 0}

		tallyJSON, err := json.Marshal(tally)
		if err != nil {
			logger.ErrorLogger.Println("Error marshalling tally")
		}

		_, err = file.Write(tallyJSON)
		if err != nil {
			logger.ErrorLogger.Println("Error writing tally to file")
		}
	}
}
