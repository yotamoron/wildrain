#!/usr/bin/python

from flask import Flask, request, session, g, redirect, url_for, abort, flash
from flask.ext.restful import Resource, Api

app = Flask(__name__)
app.debug = True

api = Api(app)

class Echo(Resource):
        def get(self, msg):
                return {'msg': msg}

api.add_resource(Echo, '/echo/<string:msg>')

if __name__ == '__main__':
    app.run()


