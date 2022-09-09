"""main flask page shit"""
from pydantic import BaseModel
from api_machine import GoRESTYourself
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()
api_function = GoRESTYourself()

origins = ['https://owl-production-backend.herokuapp.com/',
           'http://www.openweightlifting.org',
           'https://www.openweightlifting.org']

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"]
)


class SingleLifterData(BaseModel):
    """Data required to pull a single lifters history"""
    name: str
    country: str


class LeaderboardPost(BaseModel):
    """Shit for the main leaderboard"""
    gender: str
    start: int
    stop: int


@app.get("/api/")
async def default():
    """SLASH"""
    return {"error": "nothing to see here..."}


@app.get("/api/leaderboard")
def api_get_leaderboard():
    """Pulls the default leaderboard

    Example:
        GET -> resorts to default response which is male, 0 -> 99.
    """
    return api_function.lifter_totals()


@app.post("/api/leaderboard")
async def api_post_leaderboard(payload: LeaderboardPost):
    """Specific leaderboard and of a certain size

    Example:
        POST -> Payload will use the LeaderboardPost pydantic shit.
        The payload can be pure JSON with all double quotes or as explicit integers. It'll catch it.
        """
    request = payload.dict()
    try:
        gender, start, stop = request['gender'], request['start'], request['stop']
        return api_function.lifter_totals(gender, int(start), int(stop))
    except:  # YOU'RE NOT MY SUPERVISOR
        return {"get": "fucked"}


@app.post("/api/lifter/")
async def api_single_lifter(lifter: SingleLifterData):
    """lifter performance history"""
    lookup_results = api_function.lifter_lookup(lifter.dict())
    return lookup_results


@app.get("/api/lookup/{name}")
async def api_search_lifters(name: str):
    """will be used in livesearch"""
    return api_function.lifter_suggest(name) if len(name) >= 2 else [{"name": None, "country": None}]
