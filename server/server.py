#!/usr/bin/python

import os
from uuid import uuid4
from flask import Flask, request, session, g, redirect, url_for, abort, flash
from flask.ext.restful import Resource, Api
from werkzeug.utils import secure_filename
import json

app = Flask(__name__)

ALLOWED_EXTENSIONS = set(['json'])

app = Flask(__name__)

api = Api(app)

applications = {}

def addAicd(aicd):
    o = json.loads(aicd.read())
    applicationName = o['applicationName']
    if not applications.has_key(applicationName):
        applications[applicationName] = {}
    application = applications[applicationName]
    version = o['version']
    if not application.has_key(version):
        application[version] = o
    print applications

def allowed_file(filename):
    return '.' in filename and \
           filename.rsplit('.', 1)[1] in ALLOWED_EXTENSIONS

class UploadAicd(Resource):
    def post(self):
        file = request.files['aicd']
        if file and allowed_file(file.filename):
            filename = secure_filename(file.filename)
            addAicd(file)
            return { 'filename': filename }
        return { 'msg': 'Not allowed' }

class GetApplications(Resource):
    def get(self):
        return applications

api.add_resource(UploadAicd, '/uploadAicd')
api.add_resource(GetApplications, '/getApplications')

if __name__ == '__main__':
    app.run(debug=True, port=8080)


