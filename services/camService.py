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
        with open(self.filename) as json_file:
            self.data = json.load(json_file)

    def save(self):
        with open(self.filename, 'w') as outfile:
            json.dump(self.data, outfile)

    def get(self, key):
        try:
            return self.data[key]
        except KeyError:
            return False
    
    def set(self, key, value):
        self.data[key] = value
        self.save()

    def delete(self, key):
        del self.data[key]
        self.save()

    def getAll(self):
        return self.data
    
    def clear(self):
        self.data = {}
        self.save()

    def getKeys(self):
        return self.data.keys()
    