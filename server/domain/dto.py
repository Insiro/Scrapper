from typing import List, Optional

from pydantic import BaseModel

from .entity import Scrap


class URLInput(BaseModel):
    url: str


class ScrapCreate(BaseModel):
    url: str
    content: str
    author_name: str
    author_tag: str
    source: str
    comment: str = None
    image_names: List[str]


class ScrapUpdate(BaseModel):
    content: str
    author_name: str
    author_tag: str
    delete_images: list[int]
    new_images: list[int]
    comment: Optional[str]


class ImageDelete(BaseModel):
    images: List[int]


class ImageResponse(BaseModel):
    id: int
    file_name: str


class ScrapResponse(BaseModel):
    id: int
    source: str
    url: str
    content: Optional[str]
    author_name: str
    author_tag: str
    comment: Optional[str]
    images: List[ImageResponse]

    @staticmethod
    def fromScrap(scrap: Scrap):
        return ScrapResponse(
            id=scrap.id,
            url=scrap.url,
            content=scrap.content,
            comment=scrap.comment,
            author_name=scrap.author_name,
            author_tag=scrap.author_tag,
            images=[ImageResponse(id=image.id, file_name=image.file_name) for image in scrap.images],
            source=scrap.source,
        )
