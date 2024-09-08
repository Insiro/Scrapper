import logging
from typing import Any, Dict

from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session

from .database import get_db
from .domain.dto import ScrapResponse, ScrapUpdate, URLInput
from .downloader import Scrapper  # Scrapper 클래스 임포트
from .repository import *
from .utils.config import Config

router = APIRouter()


class ScrapAPIController:
    def __init__(
        self,
        repo: ScrapRepository,
        img_repo: ImageRepository,
        scrapper: Scrapper,
        logger: logging.Logger,
        cache: Dict[str, Any],
    ):
        self.repo = repo
        self.img_repo = img_repo
        self.scrapper = scrapper
        self.logger = logger
        self.cache = cache

    async def parse_url(self, url_input: URLInput):
        self.logger.info(f"Parsing URL: {url_input.url}")

        scrapper = self.scrapper.getInstance(url_input.url)
        url = scrapper.preprocess_url(url_input.url)

        if (existing_scrap := self.repo.get_scrap_by_url(url)) is not None:
            if not url_input.force:
                return ScrapResponse.fromScrap(existing_scrap)
            existing_scrap.images
            # If scrap exists and force is False, return existing scrap

        # Scrapper를 사용하여 데이터 가져오기
        data = await scrapper.scrap(url)
        scrap = self.repo.create_scrap(data)
        return ScrapResponse.fromScrap(scrap)

    async def list_scraps(self, skip: int = 0, limit: int = 10):
        scraps = self.repo.get_scraps(skip=skip, limit=limit)
        return [ScrapResponse.fromScrap(scrap) for scrap in scraps]

    async def delete_scrap(self, scrap_id: int):
        if (scrap := self.repo.delete_scrap(scrap_id)) is None:
            raise HTTPException(status_code=404, detail="Scrap not found")

        self.img_repo.delete_images_by_scrap(scrap)

        return scrap

    async def save_scrap(self, scrap_id: int, folder_path: str):
        if (scrap := self.repo.get_scrap(scrap_id)) is None:
            raise HTTPException(status_code=404, detail="Scrap not found")

        self.img_repo.save_images()

        return {"message": "Scrap saved successfully"}

    async def update_scrap(self, scrap_id: int, update_scrap: ScrapUpdate):
        if (scrap := self.repo.get_scrap(scrap_id)) is None:
            raise HTTPException(status_code=404, detail="Scrap not found")

        scrap = self.repo.update(scrap, update_scrap)
        self.img_repo.delete_images(update_scrap.delete_images)

        return scrap

    async def get_scrap_details(self, scrap_id: int):
        if (scrap := self.repo.get_scrap(scrap_id)) is None:
            raise HTTPException(status_code=404, detail="Scrap not found")

        return ScrapResponse.fromScrap(scrap)


# Dependency
def get_scrap_api_controller(
    db: Session = Depends(get_db),
    scrapper: Scrapper = Depends(Scrapper),
    config: Config = Depends(Config),
    logger: logging.Logger = Depends(lambda: logging.getLogger("api")),
    cache: Dict[str, Any] = Depends(lambda: {}),
) -> ScrapAPIController:
    repo = ScrapRepository(db=db)
    img_repo = ImageRepository(db=db, config=config)
    return ScrapAPIController(repo=repo, img_repo=img_repo, scrapper=scrapper, logger=logger, cache=cache)
