""" API Interface for OpenWeightlifting """
from fastapi import FastAPI
from .database_handler.handler_main import HandlerMain

app = FastAPI()
db_handler = HandlerMain()


@app.get("/api/top100/{sex}")
async def top100(sinclair_total: str, gender: str):
    """ This will return the top 100 lifters by sinclair or total """
    # todo: implement case switching with python 3.10.


@app.get("/api/update_index")
async def update_index_db():
    """ Lazy endpoint to update the hooked DBs index files """
    return db_handler.update_index()
