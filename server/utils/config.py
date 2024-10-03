import json
from os import getcwd, getenv, path

import sqlalchemy


class Config:
    def __init__(self) -> None:
        # 기본 경로 설정
        self.storage = getenv("SCRAPER_STORAGE", "./storage")
        self.media = path.join(self.storage, "media")
        self.export = path.join(self.storage, "export")

        # 데이터베이스 설정
        self.db_driver = getenv("SCRAPER_DB_DRIVER", "sqlite")  # 기본값은 SQLite
        self.db_name = getenv("SCRAPER_DB_NAME", "scrapper.db")

        # 기본값으로 SQLite DB URL 생성, 다른 DB의 경우 별도의 형식으로 생성
        if self.db_driver == "sqlite":
            self.db_url = f"sqlite:///{ path.join(self.storage, self.db_name)}"

        else:
            self.db_username = getenv("SCRAPER_DB_USERNAME", "")  # 사용자명 (SQLite의 경우 필요하지 않음)
            self.db_password = getenv("SCRAPER_DB_PASSWORD", "")  # 비밀번호 (SQLite의 경우 필요하지 않음)
            self.db_host = getenv("SCRAPER_DB_HOST", "")  # 데이터베이스 호스트 (SQLite의 경우 필요하지 않음)
            self.db_port = getenv("SCRAPER_DB_PORT", 5432)
            self.db_url = sqlalchemy.engine.URL.create(
                drivername=self.db_driver,
                username=self.db_username,
                password=self.db_password,
                host=self.db_host,
                port=self.db_port,
                database=self.db_name,
            )

        # 기본 URL 설정
        self.base_url = getenv("SCRAPER_BASE_PATH", "/")
        print(json.dumps(self.__dict__, indent=2))


config = Config()
