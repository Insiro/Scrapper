from typing import Optional

from server.domain.entity import ExportMode


class CreateExporter:
    title: str
    mode: ExportMode
    name_rule: Optional[str]


class UpdateExporter:
    id: int
    title: Optional[str]
    mode: Optional[ExportMode]
    name_rule: Optional[str]


class SelectExporter:
    id: Optional[int]
    title: Optional[str]
