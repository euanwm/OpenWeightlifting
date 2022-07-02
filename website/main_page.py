"""main flask page shit"""
from flask import Flask
from database_handler.api_machine import GoRESTYourself

app = Flask(__name__)
api_function = GoRESTYourself()
"""
GET, PUT, DELETE, and POST.
GET rectifies and recovers resources
PUT updates the current data
DELETE eliminates data
POST delivers new and unique data to the server.
"""


@app.route("/", methods=["GET"])
def index():
    """landing page"""
    return f"Add in the main page shit here"


@app.route("/api/lifter_totals/<gender>&start=<start>&stop=<stop>", methods=["GET"])
def api_lifter_totals(gender="male", start=0, stop=100):
    """post total bitch"""


@app.route("/api/lifter/<name>", methods=["GET"])
def api_single_lifter(name):
    """lifter performance history"""
    return api_function.lifter_lookup(name)

if __name__ == '__main__':
    app.run()