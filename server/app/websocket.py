# coding: utf-8
import json

import applications

def uploadAicd(ws):
    message = ws.receive()
    if message is None:
        return
    aicd = json.loads(message)
    applications.addAicd(aicd)

    ws.send(json.dumps({'message': 'success'}))

def getApplications(ws):
    apps = applications.getApplications()

    ws.send(json.dumps(apps))
