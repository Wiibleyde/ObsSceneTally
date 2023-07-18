package tallyManager

import (
	// json
	"encoding/json"
	"io/ioutil"
	"os"

	// log
	"log"
)

func ChangeCam(camId int, layerId int) {
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

	// change tally
	if layerId == 1 {
		tally.Layer1 = camId
	} else if layerId == 2 {
		tally.Layer2 = camId
	} else if layerId == 3 {
		tally.Layer3 = camId
	} else {
		log.Fatal("layerId must be 1, 2, or 3")
	}

	// convert tally to json
	tallyJSON, err = json.Marshal(tally)
	if err != nil {
		log.Fatal(err)
	}

	// write tally json to file
	err = ioutil.WriteFile("tally.json", tallyJSON, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func RemoveCamIdOfAllLayer(CamId int, LayerIdNotTouch int) {
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

	// change tally
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
