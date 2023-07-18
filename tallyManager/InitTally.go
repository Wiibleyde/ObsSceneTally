package tallyManager

import (
	// json
	"encoding/json"

	// log
	"log"

	// os
	"os"
)

type TallyStruct struct {
	Layer1 int
	Layer2 int
	Layer3 int
}

// InitTally() is called from main.go
// It creates the tally jsonfile if it doesn't exist
func InitTally() {
	// check if tally.json exists
	if _, err := os.Stat("tally.json"); os.IsNotExist(err) {
		// create tally.json
		file, err := os.Create("tally.json")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// create tally struct
		tally := TallyStruct{0, 0, 0}
		
		// convert tally to json
		tallyJSON, err := json.Marshal(tally)
		if err != nil {
			log.Fatal(err)
		}

		// write tally json to file
		_, err = file.Write(tallyJSON)
		if err != nil {
			log.Fatal(err)
		}
	}
}
