""" API Interface for OpenWeightlifting """
from fastapi import FastAPI

app = FastAPI()


@app.get("/api/top100/{sex}")
async def top100_sinclair(gender: str):
    """ This will return the top 100 lifters by sinclair """
    if gender == 'men':
        return gender
    elif gender == 'women':
        return gender
    elif gender is None:
        return gender
    return None
