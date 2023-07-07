import json

class CamService:
    def __init__(self):
        self.data = {}

    def setCamVisibleLayer1(self, camId):
        self.set(f"{camId}L1", True)

    def hideAllCamLayer1ExceptCam(self, camId):
        for key in list(self.data.keys()):
            if key != camId and key.endswith("L1"):
                del self.data[key]

    def setCamVisibleLayer2(self, camId):
        self.set(f"{camId}L2", True)

    def hideAllCamLayer2ExceptCam(self, camId):
        for key in list(self.data.keys()):
            if key != camId and key.endswith("L2"):
                del self.data[key]

    def setCamVisibleLayer3(self, camId):
        self.set(f"{camId}L3", True)

    def hideAllCamLayer3ExceptCam(self, camId):
        for key in list(self.data.keys()):
            if key != camId and key.endswith("L3"):
                del self.data[key]

    def delete(self, key):
        if key in self.data:
            del self.data[key]

    def getAll(self):
        return self.data

    def clear(self):
        self.data = {}

    def set(self, key, value):
        self.data[key] = value

    def get(self, key):
        return self.data.get(key, False)

    def getKeys(self):
        return self.data.keys()
