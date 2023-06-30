-- script for OBS Studio to control tally lights via api 
-- The script will see when a source named "CAM X" (X is the id of the cam) is set visible and send a request to the api to set the tally light on : apiUrl/setCamVisible?camId=X
-- The script will see when a source named "CAM X" (X is the id of the cam) is set invisible and send a request to the api to set the tally light off : apiUrl/setCamInvisible?camId=X

local apiUrl = "http://localhost:5000"

-- Function to send a request to the api
function sendRequest(url)
    local http = require("socket.http")
    local response = http.request(url)
    return response
end

-- Function to get the id of the cam from the source name
function getCamId(sourceName)
    local camId = string.match(sourceName, "CAM (%d+)")
    return camId
end

-- Function to get the source name from the id of the cam
function getSourceName(camId)
    local sourceName = "CAM " .. camId
    return sourceName
end

-- Function to look at received events in the websocket
function on_event(message)
    -- print("Got message: " .. message)
    local event = json.decode(message)
    if event["update-type"] == "SwitchScenes" then
        local sceneName = event["scene-name"]
        print("Switched to scene: " .. sceneName)
        -- If the scene is "Cameras" we set all the tally lights to off
        if sceneName == "Cameras" then
            for i = 1, 4 do
                local camId = i
                local sourceName = getSourceName(camId)
                local url = apiUrl .. "/setCamInvisible?camId=" .. camId
                local response = sendRequest(url)
                print("Set " .. sourceName .. " invisible")
            end
        end
    elseif event["update-type"] == "SceneItemVisibilityChanged" then
        local sourceName = event["scene-name"]
        local sourceVisible = event["item-visible"]
        local camId = getCamId(sourceName)
        local url = ""
        if sourceVisible then
            url = apiUrl .. "/setCamVisible?camId=" .. camId
        else
            url = apiUrl .. "/setCamInvisible?camId=" .. camId
        end
        local response = sendRequest(url)
        print("Set " .. sourceName .. " visible: " .. tostring(sourceVisible))
    end
end

