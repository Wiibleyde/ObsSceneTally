# ObsSceneTally

## Description

This is a simple script that will show to a camera if it is live or not. It will also show the layer of the scene that is live.

## Installation

1. Download the script from [here](https://github.com/Wiibleyde/ObsSceneTally/releases)
2. Extract the zip file
3. Open the folder in a terminal
4. Run `docker compose up -d`

## Usage

### To show a camera if it is live or not

1. Go to a browser and type in the ip address of the computer running the script
2. Add `/cam?camId=X` to the end of the url where X is the id of the camera you want to show

### To show the layer of the scene that is live

1. Go to a browser and type in the ip address of the computer running the script
2. Add `/api/setCam?camId=X&layerId=Y` to the end of the url where X is the id of the camera you want to show and Y is the id of the layer you want to show.

### To get the camera showing on a specific layer

1. Go to a browser and type in the ip address of the computer running the script
2. Add `/api/getLayer?layerId=X` to the end of the url where X is the id of the layer you want to show.
