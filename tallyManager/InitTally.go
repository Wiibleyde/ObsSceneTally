package tallyManager

import (
	"encoding/json"

	"log"

	"os"
)

type TallyStruct struct {
	Layer1 int
	Layer2 int
	Layer3 int
}

func InitTally() {
	if _, err := os.Stat("data/tally.json"); os.IsNotExist(err) {
		file, err := os.Create("data/tally.json")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		tally := TallyStruct{0, 0, 0}
		
		tallyJSON, err := json.Marshal(tally)
		if err != nil {
			log.Fatal(err)
		}

		_, err = file.Write(tallyJSON)
		if err != nil {
			log.Fatal(err)
		}
	}
}
