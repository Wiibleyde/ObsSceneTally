package tallyManager

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func GetLayerInfo(layerId int) int {
	file, err := os.Open("data/tally.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	tallyJSON, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var tally TallyStruct
	err = json.Unmarshal(tallyJSON, &tally)
	if err != nil {
		log.Fatal(err)
	}

	if layerId == 1 {
		return tally.Layer1
	} else if layerId == 2 {
		return tally.Layer2
	} else if layerId == 3 {
		return tally.Layer3
	} else {
		log.Fatal("layerId must be 1, 2, or 3")
	}
	return 0
}
