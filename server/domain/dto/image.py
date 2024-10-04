from typing import List
from pydantic import BaseModel


class ImageDelete(BaseModel):
    images: List[int]


class ImageResponse(BaseModel):
    id: int
    file_name: str

