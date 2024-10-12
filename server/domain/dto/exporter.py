from typing import Optional

from pydantic import BaseModel

from server.domain.entity import ExportMode


class CreateExporter(BaseModel):
    title: str
    mode: ExportMode
    name_rule: Optional[str]


class UpdateExporter(BaseModel):
    id: int
    title: Optional[str]
    mode: Optional[ExportMode]
    name_rule: Optional[str]


class SelectExporter(BaseModel):
    id: Optional[int]
    title: Optional[str]
