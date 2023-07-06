from flask import Flask, render_template, request, json
import flask.cli
import logging

from services.camService import CamService

app = Flask(__name__)

@app.route('/cam/<int:camId>')
def cam(camId):
    print("Call to /cam/" + str(camId) + " received")
    return render_template('cam.html', camId=camId)

@app.route('/getCamStatus', methods=['GET'])
def getCamStatus():
    print("Call to /getCamStatus received with camId=" + request.args.get('camId'))
    camId = request.args.get('camId')
    layer = request.args.get('layer')
    if layer == "1":
        return {'status': camServiceObj.get(f"{camId}L1")}
    elif layer == "2":
        return {'status': camServiceObj.get(f"{camId}L2")}
    elif layer == "3":
        return {'status': camServiceObj.get(f"{camId}L3")}
    else:
        return {'status': 'ERROR'}


@app.route('/setCamVisible', methods=['GET'])
def setCamVisible():
    print("Call to /setCamVisible received with camId=" + request.args.get('camId'))
    camId = request.args.get('camId')
    layer = request.args.get('layer')
    if layer == "1":
        camServiceObj.hideAllCamLayer1ExceptCam(camId)
        camServiceObj.setCamVisibleLayer1(camId)
    elif layer == "2":
        camServiceObj.hideAllCamLayer2ExceptCam(camId)
        camServiceObj.setCamVisibleLayer2(camId)
    elif layer == "3":
        camServiceObj.hideAllCamLayer3ExceptCam(camId)
        camServiceObj.setCamVisibleLayer3(camId)
    return {'status': 'OK'}

@app.route('/setCamInvisible', methods=['GET'])
def setCamInvisible():
    print("Call to /setCamInvisible received with camId=" + request.args.get('camId'))
    camId = request.args.get('camId')
    layer = request.args.get('layer')
    if layer == "1":
        camServiceObj.set(f"{camId}L1", False)
    elif layer == "2":
        camServiceObj.set(f"{camId}L2", False)
    elif layer == "3":
        camServiceObj.set(f"{camId}L3", False)
    return {'status': 'OK'}

@app.route('/setAllCamsInvisible', methods=['GET'])
def setAllCamsInvisible():
    print("Call to /setAllCamsInvisible received")
    camServiceObj.clear()
    return {'status': 'OK'}

if __name__ == '__main__':
    camServiceObj = CamService("cam.json")
    flaskLog = logging.getLogger('werkzeug')
    flaskLog.disabled = True
    flask.cli.show_server_banner = lambda *args: None
    app.run(debug=True, host='0.0.0.0', port=5000)