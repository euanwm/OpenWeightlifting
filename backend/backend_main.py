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
def default():
    """SLASH"""
    return {"error": "nothing to see here..."}


@app.route("/api/leaderboard", methods=[HTTP.GET, HTTP.POST])
def api_leaderboard():
    """Pulls the gendered leaderboards

    Example:
        GET -> resorts to default response which is male, 0 -> 99.

        POST -> payload should look like this -> {"gender: "male", "start": 100, "stop": 200}.
        The payload can be pure JSON with all double quotes or as explicit integers. It'll catch it.
    """
    match request.method:
        case HTTP.GET:
            return jsonify(api_function.lifter_totals())
        case HTTP.POST:
            try:
                gender, start, stop = request.json['gender'], request.json['start'], request.json['stop']
                return jsonify(api_function.lifter_totals(gender, int(start), int(stop)))
            except:
                return {"get": "fucked"}


@app.route("/api/lifter", methods=[HTTP.POST])
def api_single_lifter():
    """lifter performance history"""
    return jsonify(api_function.lifter_lookup(request.json))


@app.route("/api/lookup/<name>", methods=[HTTP.GET])
def api_search_lifters(name):
    """will be used in livesearch"""
    return jsonify(api_function.lifter_suggest(name)) if len(name) >= 2 else jsonify([{"name": None, "country": None}])


if __name__ == '__main__':
    app.run(host="0.0.0.0",
            port=8000)
