import logging
from typing import Any, Dict, Self

from fastapi import Depends, HTTPException
from sqlalchemy.orm import Session

from server.database import get_db
from server.domain.dto.exporter import CreateExporter, SelectExporter, UpdateExporter
from server.repository import ExporterRepository


class ExporterController:
    def __init__(
        self,
        repo: ExporterRepository,
        logger: logging.Logger,
        cache: Dict[str, Any],
    ):
        self.repo = repo
        self.logger = logger
        self.cache = cache

    async def create_exporter(self, create: CreateExporter):
        exporter = self.repo.create_exporter(create)
        return exporter

    async def update_exporter(self, update: UpdateExporter):
        exporter = self.repo.get_exporter(update.id)
        if exporter.__len__ == 0:
            raise HTTPException("404")
        return self.repo.update_exporter(exporter[0], update)

    async def list_exporter(self, selection: SelectExporter):
        return self.repo.get_exporter(selection.id, selection.title)

    @staticmethod
    def depends(
        db: Session = Depends(get_db),
        logger: logging.Logger = Depends(lambda: logging.getLogger("api")),
        cache: Dict[str, Any] = Depends(lambda: {}),
    ) -> Self:
        repo = ExporterRepository(db=db)
        return ExporterController(repo=repo, logger=logger, cache=cache)
