from os import makedirs
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

from server.database import engine, Base
from server.scrapRouter import router
from server.utils.config import config

makedirs(config.media, exist_ok=True)
makedirs(config.export, exist_ok=True)


app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["http://localhost:5173"],  # React 앱의 URL
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)



# Create the database tables
Base.metadata.create_all(bind=engine)

# Include the router
app.include_router(router, prefix="/api")


@app.get("/")
def root():
    return {"message": "downloader server"}


print("downloader start!!")
