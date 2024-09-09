from typing import List, Optional

from sqlalchemy.orm import Session

from ..domain.dto import ScrapCreate, ScrapUpdate
from ..domain.entity import Scrap


class ScrapRepository:
    def __init__(self, db: Session):
        self.db = db

    def create_scrap(self, scrap_data: ScrapCreate) -> Scrap:
        db_scrap = Scrap(
            url=scrap_data.url,
            author_name=scrap_data.author_name,
            author_tag=scrap_data.author_tag,
            source=scrap_data.source,
            content=scrap_data.content,
            comment =scrap_data.comment
        )
        self.db.add(db_scrap)
        self.db.commit()
        self.db.refresh(db_scrap)

        self.db.commit()
        return db_scrap

    def update_scrap(self, scrap: Scrap, scrap_data: ScrapUpdate) -> Scrap:
        if scrap_data.author_name is not None:
            scrap.author_name = scrap_data.author_name
        if scrap_data.author_tag is not None:
            scrap.author_tag = scrap_data.author_tag
        if scrap_data.content is not None:
            scrap.content = scrap_data.content
        if scrap_data.comment is not None:
            scrap.comment = scrap_data.comment

        self.db.commit()
        self.db.refresh(scrap)
        return scrap

    def get_scrap_by_url(self, url: str) -> Optional[Scrap]:
        return self.db.query(Scrap).filter(Scrap.url == url).first()

    def get_scraps(self, skip: int = 0, limit: int = 10) -> List[Scrap]:
        return self.db.query(Scrap).offset(skip).limit(limit).all()

    def get_scrap(self, scrap_id: int) -> Optional[Scrap]:
        return self.db.query(Scrap).filter(Scrap.id == scrap_id).first()

    def delete_scrap(self, scrap_id: int) -> Optional[Scrap]:
        scrap = self.db.query(Scrap).filter(Scrap.id == scrap_id).first()
        if scrap is None:
            return None

        for image in scrap.images:
            self.db.delete(image)
        self.db.delete(scrap)
        self.db.commit()
        return scrap
