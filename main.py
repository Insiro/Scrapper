from os import makedirs, path
from fastapi import FastAPI
from fastapi.staticfiles import StaticFiles
from fastapi.middleware.cors import CORSMiddleware

from server.database import engine, Base
from server.scrapRouter import router
from server.utils.config import config

makedirs(config.media, exist_ok=True)
makedirs(config.export, exist_ok=True)


app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["http://localhost:5173"],  # React_vite URL for Test
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

base_url = config.base_url

# Create the database tables
Base.metadata.create_all(bind=engine)

# Include the router
app.mount(path.join(base_url, "/media"), StaticFiles(directory=config.media), name="media")
app.include_router(router, prefix=path.join(base_url,"/api"))


@app.get("/")
def root():
    return {"message": "Scraper server"}


print("SCRAPER start!!")
