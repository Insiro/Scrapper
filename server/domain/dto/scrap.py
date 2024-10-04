from typing import List, Optional

from pydantic import BaseModel

from server.domain.dto.image import ImageResponse

from ..entity import Scrap, Tags
from ..PageType import PageType


class URLInput(BaseModel):
    url: str


class ScrapModifier(BaseModel):
    content: str
    author_name: str
    author_tag: str
    pin: Optional[bool] = None
    comment: Optional[str] = None
    tags: Optional[List[str]] = None


class ScrapCreate(ScrapModifier):
    source_key: str
    source: str
    image_names: List[str]


class ScrapUpdate(ScrapModifier):
    delete_images: list[int]
    new_images: list[int]


class ScrapResponse(BaseModel):
    id: int
    source: str
    source_id: str
    url: str
    content: Optional[str]
    author_name: str
    author_tag: str
    comment: Optional[str]
    images: List[ImageResponse]
    pin: bool
    tags: list[str] = []

    @staticmethod
    def fromScrap(scrap: Scrap, tags: list[Tags] = None):
        if tags is None:
            tags = []

        return ScrapResponse(
            id=scrap.id,
            source_id=scrap.source_id,
            content=scrap.content,
            comment=scrap.comment,
            url=PageType[scrap.source].url(scrap.source_id),
            author_name=scrap.author_name,
            author_tag=scrap.author_tag,
            images=[ImageResponse(id=image.id, file_name=image.file_name) for image in scrap.images],
            source=scrap.source,
            pin=scrap.pin,
            tags=[tag.name for tag in tags],
        )

    def setTag(self, tags: list[Tags]):
        self.tags.extend([tag.name for tag in tags])
