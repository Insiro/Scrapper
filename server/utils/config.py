from os import getcwd, getenv, path


class Config:
    storage = getenv("STORAGE", path.join(getcwd(), "storage"))
    media = path.join(storage, "media")
    export = path.join(storage, "export")
    db_url = getenv(
        "DATABASE_URL", f"sqlite:///{path.join(storage, 'test.db')}"
    )  # 기본값으로 SQLite


config = Config()
