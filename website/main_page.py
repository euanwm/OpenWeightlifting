"""main flask page shit"""
from enum import Enum
from flask import Flask, request, jsonify
from database_handler.api_machine import GoRESTYourself

app = Flask(__name__)
api_function = GoRESTYourself()


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
    """landing page"""
    return f"Add in the main page shit here"


@app.route("/api/lifter_totals", methods=[HTTP.GET, HTTP.POST])
def api_lifter_totals():
    """post total bitch"""
    match request.method:
        case HTTP.GET:
            return jsonify(api_function.lifter_totals())
        case HTTP.POST:
            # todo: build the payload parser thingy
            return jsonify(api_function.lifter_totals())
        case _:
            return jsonify({"error": "not a valid method"})


@app.route("/api/lifter/<name>", methods=[HTTP.GET])
def api_single_lifter(name):
    """lifter performance history"""
    return jsonify(api_function.lifter_lookup(name))


@app.route("/api/lookup/<name>", methods=[HTTP.GET])
def api_search_lifters(name):
    """will be used in livesearch"""
    return jsonify(api_function.lifter_suggest(name)) if len(name) >= 2 else jsonify([{"name": None, "country": None}])


if __name__ == '__main__':
    app.run()
