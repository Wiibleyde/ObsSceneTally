import json
import os

class CamService:
    def __init__(self,filename):
        self.filename = "data/" + filename
        if not os.path.exists("data"):
            os.makedirs("data")
            self.data = {}
            self.save()
        else:
            self.load()

    def load(self):
        try:
            with open(self.filename) as json_file:
                self.data = json.load(json_file)
        except:
            self.data = {}
            self.save()

    def save(self):
        with open(self.filename, 'w') as outfile:
            json.dump(self.data, outfile)

    def setCamVisibleLayer1(self, camId):
        self.set(f"{camId}L1", True)

    def hideAllCamLayer1ExceptCam(self, camId):
        for key in self.data.keys():
            if key != camId:
                if key.endswith("L1"):
                    self.set(key, False)

    def setCamVisibleLayer2(self, camId):
        self.set(f"{camId}L2", True)

    def hideAllCamLayer2ExceptCam(self, camId):
        for key in self.data.keys():
            if key != camId:
                if key.endswith("L2"):
                    self.set(key, False)

    def setCamVisibleLayer3(self, camId):
        self.set(f"{camId}L3", True)

    def hideAllCamLayer3ExceptCam(self, camId):
        for key in self.data.keys():
            if key != camId:
                if key.endswith("L3"):
                    self.set(key, False)
        
    def delete(self, key):
        del self.data[key]
        self.save()

    def getAll(self):
        return self.data
    
    def clear(self):
        self.data = {}
        self.save()

    def set(self, key, value):
        self.data[key] = value
        self.save()

    def get(self, key):
        if key in self.data:
            return self.data[key]
        else:
            return False

    def getKeys(self):
        return self.data.keys()
    