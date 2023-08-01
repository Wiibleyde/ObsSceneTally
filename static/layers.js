let apiLink = "/api/getLayers"

function updateCamStatus() {
// Fetch layers status and update number of the cam in every layers
fetch(apiLink)
    .then(response => response.text())
    .then(layers => {
        layers = layers.split(',')
        let layer1 = layers[0]
        let layer2 = layers[1]
        let layer3 = layers[2]
        if(layer1 == 0) {
            document.getElementById("layer1").innerHTML = "Offline"
            document.getElementById("layer1text").style = "color:red"
        } else if(layer1 == -1) {
            document.getElementById("layer1").innerHTML = "Video"
            document.getElementById("layer1text").style = "color:blue"
        } else {
            document.getElementById("layer1").innerHTML = layer1
            document.getElementById("layer1text").style = "color:green"
        }
        if(layer2 == 0) {
            document.getElementById("layer2").innerHTML = "Offline"
            document.getElementById("layer2text").style = "color:red"
        } else if(layer2 == -1) {
            document.getElementById("layer2").innerHTML = "Video"
            document.getElementById("layer2text").style = "color:blue"
        } else {
            document.getElementById("layer2").innerHTML = layer2
            document.getElementById("layer2text").style = "color:green"
        }
        if(layer3 == 0) {
            document.getElementById("layer3").innerHTML = "Offline"
            document.getElementById("layer3text").style = "color:red"
        } else if(layer3 == -1) {
            document.getElementById("layer3").innerHTML = "Video"
            document.getElementById("layer3text").style = "color:blue"
        } else {
            document.getElementById("layer3").innerHTML = layer3
            document.getElementById("layer3text").style = "color:green"
        }
    })
    .catch(error => {
        console.error('Error fetching camera status:', error);
        document.getElementById("layer1").innerHTML = "Error"
        document.getElementById("layer1text").style = "color:gray"
        document.getElementById("layer2").innerHTML = "Error"
        document.getElementById("layer2text").style = "color:gray"
        document.getElementById("layer3").innerHTML = "Error"
        document.getElementById("layer3text").style = "color:gray"
    });
}

setInterval(updateCamStatus, 500);