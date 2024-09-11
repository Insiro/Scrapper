import os
from os import getenv, path, getcwd


class Config:
    def __init__(self) -> None:
        # 기본 경로 설정
        self.storage = getenv("SCRAPER_STORAGE", path.join(getcwd(), "storage"))
        self.media = path.join(self.storage, "media")
        self.export = path.join(self.storage, "export")

        # 데이터베이스 설정
        self.db_driver = getenv("SCRAPER_DB_DRIVER", "sqlite")  # 기본값은 SQLite
        self.db_username = getenv(
            "SCRAPER_DB_USERNAME", ""
        )  # 사용자명 (SQLite의 경우 필요하지 않음)
        self.db_password = getenv(
            "SCRAPER_DB_PASSWORD", ""
        )  # 비밀번호 (SQLite의 경우 필요하지 않음)
        self.db_host = getenv(
            "SCRAPER_DB_HOST", ""
        )  # 데이터베이스 호스트 (SQLite의 경우 필요하지 않음)
        self.db_port = getenv("SCRAPER_DB_PORT", 5432)
        self.db_name = getenv("SCRAPER_DB_NAME", path.join(self.media, "scrapper.db"))

        # 기본값으로 SQLite DB URL 생성, 다른 DB의 경우 별도의 형식으로 생성
        if self.db_driver == "sqlite":
            self.db_url = f"sqlite:///{self.db_name}"
        else:
            self.db_url = f"{self.db_driver}://{self.db_username}:{self.db_password}@{self.db_host}:{self.db_port}/{self.db_name}"

        # 기본 URL 설정
        self.base_url = getenv("SCRAPER_BASE_PATH", "/")


config = Config()
import json

print(json.dumps(config.__dict__, indent=2))
