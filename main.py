from flask import Flask, render_template, request, json
import flask.cli
import logging

from services.camService import CamService

app = Flask(__name__)

@app.route('/cam/<int:camId>')
def cam(camId):
    # print("Call to /cam/" + str(camId) + " received")
    return render_template('cam.html', camId=camId)

@app.route('/getCamStatus', methods=['GET'])
def getCamStatus():
    # print("Call to /getCamStatus received with camId=" + request.args.get('camId'))
    camId = request.args.get('camId')
    # json response
    return {'status': camServiceObj.get(camId)}

@app.route('/setCamVisible', methods=['GET'])
def setCamVisible():
    # print("Call to /setCamVisible received with camId=" + request.args.get('camId'))
    camId = request.args.get('camId')
    camServiceObj.set(camId, True)
    return {'status': 'OK'}

@app.route('/setCamInvisible', methods=['GET'])
def setCamInvisible():
    # print("Call to /setCamInvisible received with camId=" + request.args.get('camId'))
    camId = request.args.get('camId')
    camServiceObj.set(camId, False)
    # return response 200
    return {'status': 'OK'}

@app.route('/setAllCamsInvisible', methods=['GET'])
def setAllCamsInvisible():
    # print("Call to /setAllCamsInvisible received")
    camServiceObj.clear()
    # return response 200
    return {'status': 'OK'}

if __name__ == '__main__':
    camServiceObj = CamService("cam.json")
    flaskLog = logging.getLogger('werkzeug')
    flaskLog.disabled = True
    flask.cli.show_server_banner = lambda *args: None
    app.run(debug=True, host='0.0.0.0', port=5000)