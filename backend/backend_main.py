"""main flask page shit"""
from enum import Enum
from flask import Flask, request, jsonify
from flask_cors import CORS
from api_machine import GoRESTYourself

app = Flask(__name__)
api_function = GoRESTYourself()
CORS(app, resources={r"/api/*": {"origins": "*"}})


class HTTP(str, Enum):
    """HTTP request methods"""
    GET = "GET"
    POST = "POST"
    PUT = "PUT"
    DELETE = "DELETE"
    PATCH = "PATCH"
    OPTION = "OPTION"

@app.route("/", methods=[HTTP.GET])
def index():
    """standard shit"""
    return {"nothing to see here": None}

@app.route("/api/top100", methods=[HTTP.GET])
def default_top100():
    """bog standard shit"""
    return jsonify(api_function.lifter_totals())


@app.route("/api/lifter", methods=[HTTP.POST])
def api_single_lifter():
    """lifter performance history"""
    return jsonify(api_function.lifter_lookup(request.json))


@app.route("/api/lookup/<name>", methods=[HTTP.GET])
def api_search_lifters(name):
    """will be used in livesearch"""
    return jsonify(api_function.lifter_suggest(name)) if len(name) >= 2 else jsonify([{"name": None, "country": None}])


if __name__ == '__main__':
    app.run()
