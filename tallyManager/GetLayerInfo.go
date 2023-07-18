package tallyManager

import (
	// json
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func GetLayerInfo(layerId int) int {
	// open tally.json
	file, err := os.Open("tally.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read tally.json
	tallyJSON, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// convert tally.json to tally struct
	var tally TallyStruct
	err = json.Unmarshal(tallyJSON, &tally)
	if err != nil {
		log.Fatal(err)
	}

	// get layer info
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
