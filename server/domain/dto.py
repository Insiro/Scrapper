from typing import List, Optional

from pydantic import BaseModel

from .entity import Scrap, Tags


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
    url: str
    source: str
    image_names: List[str]


class ScrapUpdate(ScrapModifier):
    delete_images: list[int]
    new_images: list[int]


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
    pin: bool
    tags: list[str] = []

    @staticmethod
    def fromScrap(scrap: Scrap, tags: list[Tags] = None):
        if tags is None:
            tags = []
        return ScrapResponse(
            id=scrap.id,
            url=scrap.url,
            content=scrap.content,
            comment=scrap.comment,
            author_name=scrap.author_name,
            author_tag=scrap.author_tag,
            images=[ImageResponse(id=image.id, file_name=image.file_name) for image in scrap.images],
            source=scrap.source,
            pin=scrap.pin,
            tags=[tag.name for tag in tags],
        )

    def setTag(self, tags: list[Tags]):
        self.tags.extend([tag.name for tag in tags])
