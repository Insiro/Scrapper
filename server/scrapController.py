import logging
from typing import Any, Dict

from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session

from .database import get_db
from .domain.dto import ImageDelete, ScrapResponse, ScrapUpdate, URLInput
from .repository import *
from .scrapper import ScrapperFactory  # Scrapper 클래스 임포트
from .utils import config

router = APIRouter()


class ScrapAPIController:
    def __init__(
        self,
        repo: ScrapRepository,
        img_repo: ImageRepository,
        scrapper: ScrapperFactory,
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
            return ScrapResponse.fromScrap(existing_scrap)

        # Scrapper를 사용하여 데이터 가져오기
        data = await scrapper.scrap(url)
        scrap = self.repo.create_scrap(data)
        self.img_repo.save_images(data.image_names, scrap.id)

        return ScrapResponse.fromScrap(scrap)

    async def re_scrap(self, scrap_id: int):
        existing_scrap = self.repo.get_scrap(scrap_id)
        scrapper = self.scrapper.getInstance(type_name=existing_scrap.source)
        url = scrapper.preprocess_url(existing_scrap.url)
        data = await scrapper.scrap(url)
        scrap = self.repo.update_scrap(existing_scrap, data)

        self.img_repo.delete_images_by_scrap(scrap.id)
        self.img_repo.save_images(data.image_names, scrap.id)

        return ScrapResponse.fromScrap(scrap)

    async def list_scraps(self, page: int = 1, limit: int = 20):
        scraps = self.repo.get_scraps(page=page, limit=limit)
        count = self.repo.count_scrap()
        return {"list": [ScrapResponse.fromScrap(scrap) for scrap in scraps], "count": count}

    async def delete_scrap(self, scrap_id: int):
        if (scrap := self.repo.delete_scrap(scrap_id)) is None:
            raise HTTPException(status_code=404, detail="Scrap not found")

        self.img_repo.delete_images_by_scrap(scrap)

        return scrap

    async def update_scrap(self, scrap_id: int, update_scrap: ScrapUpdate):
        if (scrap := self.repo.get_scrap(scrap_id)) is None:
            raise HTTPException(status_code=404, detail="Scrap not found")

        scrap = self.repo.update(scrap, update_scrap)
        self.img_repo.delete_images(*update_scrap.delete_images)

        return scrap

    async def get_scrap_details(self, scrap_id: int):
        if (scrap := self.repo.get_scrap(scrap_id)) is None:
            raise HTTPException(status_code=404, detail="Scrap not found")

        return ScrapResponse.fromScrap(scrap)

    async def delete_image(self, image_id: int):
        if (images := self.img_repo.get_image(image_id)) is None:
            raise HTTPException(status_code=404, detail="Image not found")
        return images

    async def delete_images(self, delete_image: ImageDelete):
        result = self.img_repo.delete_images(*delete_image.images)

        return result


# Dependency
def get_scrap_api_controller(
    db: Session = Depends(get_db),
    scrapper: ScrapperFactory = Depends(ScrapperFactory),
    logger: logging.Logger = Depends(lambda: logging.getLogger("api")),
    cache: Dict[str, Any] = Depends(lambda: {}),
) -> ScrapAPIController:
    repo = ScrapRepository(db=db)
    img_repo = ImageRepository(db=db, config=config)
    return ScrapAPIController(repo=repo, img_repo=img_repo, scrapper=scrapper, logger=logger, cache=cache)
