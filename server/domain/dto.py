from typing import List, Optional

from pydantic import BaseModel

from .entity import Scrap


class URLInput(BaseModel):
    url: str
    force: bool = False


class ScrapCreate(BaseModel):
    url: str
    content: str
    author_name: str
    author_tag: str
    source: str
    image_names: List[str]


class ScrapUpdate(BaseModel):
    content: str
    author_name: str
    author_tag: str
    delete_images: list[int]
    comment: Optional[str]


class ScrapResponse(BaseModel):
    id: int
    url: str
    content: Optional[str]
    author_name: str
    author_tag: str
    image_names: List[str]

    @staticmethod
    def fromScrap(scrap: Scrap):
        return ScrapResponse(
            id=scrap.id,
            url=scrap.url,
            content=scrap.content,
            author_name=scrap.author_name,
            author_tag=scrap.author_tag,
            image_names=[image.file_name for image in scrap.images],
        )
