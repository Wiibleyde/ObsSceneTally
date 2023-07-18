package tallyManager

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"log"
)

func ChangeCam(camId int, layerId int) {
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
		tally.Layer1 = camId
	} else if layerId == 2 {
		tally.Layer2 = camId
	} else if layerId == 3 {
		tally.Layer3 = camId
	} else {
		log.Fatal("layerId must be 1, 2, or 3")
	}

	tallyJSON, err = json.Marshal(tally)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("data/tally.json", tallyJSON, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func RemoveCamIdOfAllLayer(CamId int, LayerIdNotTouch int) {
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
