import os
from os import getenv, path, getcwd

class Config:
    # 기본 경로 설정
    storage = getenv("SCRAPER_STORAGE", path.join(getcwd(), "storage"))
    media = path.join(storage, "media")
    export = path.join(storage, "export")
    
    # 데이터베이스 설정
    db_driver = getenv("SCRAPER_DB_DRIVER", "sqlite")  # 기본값은 SQLite
    db_username = getenv("SCRAPER_DB_USERNAME", "")  # 사용자명 (SQLite의 경우 필요하지 않음)
    db_password = getenv("SCRAPER_DB_PASSWORD", "")  # 비밀번호 (SQLite의 경우 필요하지 않음)
    db_host = getenv("SCRAPER_DB_HOST", "")  # 데이터베이스 호스트 (SQLite의 경우 필요하지 않음)
    db_name = getenv("SCRAPER_DB_NAME", path.join(storage, "scrapper.db"))

    # 기본값으로 SQLite DB URL 생성, 다른 DB의 경우 별도의 형식으로 생성
    if db_driver == "sqlite":
        db_url = f"sqlite:///{db_name}"
    else:
        db_url = f"{db_driver}://{db_username}:{db_password}@{db_host}/{db_name}"
    
    # 기본 URL 설정
    base_url = getenv("SCRAPER_BASE_URL", "/")

config = Config()