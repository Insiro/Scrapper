import logging
from typing import Any, Dict, Self

from fastapi import Depends, HTTPException
from sqlalchemy.orm import Session

from server.database import get_db
from server.domain.dto.image import ImageDelete
from server.domain.dto.scrap import ScrapResponse, ScrapUpdate, URLInput
from server.repository import *
from server.scrapper import Scrapper
from server.utils import config


class ScrapAPIController:
    def __init__(
        self,
        repo: ScrapRepository,
        img_repo: ImageRepository,
        logger: logging.Logger,
        cache: Dict[str, Any],
    ):
        self.repo = repo
        self.img_repo = img_repo
        self.logger = logger
        self.cache = cache

    async def parse_url(self, url_input: URLInput):
        self.logger.info(f"Parsing URL: {url_input.url}")

        scrapper = Scrapper(url_input.url)

        if (existing_scrap := self.repo.get_scrap_by_url(scrapper.pageType, scrapper.args.key)) is not None:
            return ScrapResponse.fromScrap(existing_scrap)

        # Scrapper를 사용하여 데이터 가져오기
        data = await scrapper.scrap()
        scrap = self.repo.create_scrap(data)
        self.img_repo.save_images(data.image_names, scrap.id)

        tags = self.repo.get_tags(scrap_id=scrap.id)
        return ScrapResponse.fromScrap(scrap, tags)

    async def re_scrap(self, scrap_id: int):
        existing_scrap = self.repo.get_scrap(scrap_id)
        scrapper = Scrapper(existing_scrap.url, type=existing_scrap.source)
        data = await scrapper.scrap()
        scrap = self.repo.update_scrap(existing_scrap, data)

        self.img_repo.delete_images_by_scrap(scrap.id)
        self.img_repo.save_images(data.image_names, scrap.id)

        tags = self.repo.get_tags(scrap_id=scrap_id)

        return ScrapResponse.fromScrap(scrap, tags)

    async def list_scraps(self, offset=0, limit: int = 20, pined=False):
        scraps = self.repo.get_scraps(offset=offset, limit=limit, pined=pined)
        count = self.repo.count_scrap()

        scrapReturns = []
        for scrap in scraps:
            tags = self.repo.get_tags(scrap_id=scrap.id)
            scrapReturns.append(ScrapResponse.fromScrap(scrap, tags=tags))

        return {"list": scrapReturns, "count": count}

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

        tags = self.repo.get_tags(scrap_id=scrap_id)
        return ScrapResponse.fromScrap(scrap, tags)

    async def get_scrap_details(self, scrap_id: int):
        if (scrap := self.repo.get_scrap(scrap_id)) is None:
            raise HTTPException(status_code=404, detail="Scrap not found")

        tags = self.repo.get_tags(scrap_id=scrap_id)
        return ScrapResponse.fromScrap(scrap, tags)

    async def delete_image(self, image_id: int):
        if (images := self.img_repo.get_image(image_id)) is None:
            raise HTTPException(status_code=404, detail="Image not found")
        return images

    async def delete_images(self, delete_image: ImageDelete):
        result = self.img_repo.delete_images(*delete_image.images)

        return result

    @staticmethod
    def depends(
        db: Session = Depends(get_db),
        logger: logging.Logger = Depends(lambda: logging.getLogger("api")),
        cache: Dict[str, Any] = Depends(lambda: {}),
    ) -> Self:
        repo = ScrapRepository(db=db)
        img_repo = ImageRepository(db=db, config=config)
        return ScrapAPIController(repo=repo, img_repo=img_repo, logger=logger, cache=cache)
