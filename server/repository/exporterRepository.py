from typing import List

from sqlalchemy.orm import Session

from server.domain.dto.exporter import CreateExporter, UpdateExporter
from server.domain.entity import Exporter


class ExporterRepository:
    def __init__(self, db: Session):
        self.db = db

    def create_exporter(self, create: CreateExporter) -> Exporter:
        exporter = Exporter(title=create.title, mode=create.mode)
        if create.name_rule is not None:
            exporter.name_rule = create.name_rule
        self.db.add(exporter)
        self.db.commit()
        self.db.refresh(exporter)

        return exporter

    def update_exporter(self, exporter: Exporter, update: UpdateExporter) -> Exporter:
        if update.title is not None:
            exporter.title = update.title
        if update.mode is not None:
            exporter.mode = update.mode
        if update.name_rule is not None:
            exporter.name_rule = update.name_rule

        self.db.commit()
        self.db.refresh(exporter)

        return exporter

    def get_exporter(self, exporter_id: int = None, title: str = None) -> List[Exporter]:
        query = self.db.query(Exporter)
        if exporter_id is not None:
            query = query.filter_by(id=exporter_id)
        if title is not None:
            query = query.filter_by(title=title)
        return query.all()

    def delete_exporter(self, exporter: Exporter):
        self.db.delete(exporter)

        return
