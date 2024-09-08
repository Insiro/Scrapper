from os import makedirs
from fastapi import FastAPI
from src.database import engine, Base
from src.scrapRouter import router
from src.utils.config import config

makedirs(config.media, exist_ok=True)
makedirs(config.export, exist_ok=True)


app = FastAPI()

# Create the database tables
Base.metadata.create_all(bind=engine)

# Include the router
app.include_router(router, prefix="/api")


@app.get("/")
def root():
    return {"message": "downloader server"}


print("downloader start!!")
