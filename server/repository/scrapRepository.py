from typing import List, Optional

from sqlalchemy import desc
from sqlalchemy.orm import Session

from ..domain.dto import ScrapCreate, ScrapUpdate
from ..domain.entity import Scrap, Tags


class ScrapRepository:
    def __init__(self, db: Session):
        self.db = db

    def put_tag(self, scrap_id, tagList: list[str]):
        self.db.query(Tags).filter(Tags.scrap_id == scrap_id).delete()
        tags = [Tags(scrap_id=scrap_id, name=tag) for tag in tagList]
        self.db.add_all(tags)

    def create_scrap(self, scrap_data: ScrapCreate) -> Scrap:
        db_scrap = Scrap(
            url=scrap_data.url,
            author_name=scrap_data.author_name,
            author_tag=scrap_data.author_tag,
            source=scrap_data.source,
            content=scrap_data.content,
            comment=scrap_data.comment,
        )
        self.db.add(db_scrap)
        self.db.commit()
        self.db.refresh(db_scrap)

        if scrap_data.tags is not None:
            self.put_tag(db_scrap.id, scrap_data.tags)

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
        if scrap_data.pin is not None:
            scrap.pin = scrap_data.pin

        if scrap_data.tags is not None:
            self.put_tag(scrap.id, scrap_data.tags)

        self.db.commit()
        self.db.refresh(scrap)

        if scrap_data.tags is not None:
            self.put_tag(scrap.id, scrap_data.tags)

        return scrap

    def get_scrap_by_url(self, url: str) -> Optional[Scrap]:
        return self.db.query(Scrap).filter(Scrap.url == url).first()

    def get_scraps(self, offset: int = 0, limit: int = 20, pined=False) -> List[Scrap]:
        query = self.db.query(Scrap).outerjoin(Tags)
        if pined == True:
            query = query.filter(Scrap.pin == True)

        return query.order_by(desc(Scrap.id)).offset(offset).limit(limit).all()

    def get_scrap(self, scrap_id: int) -> Optional[Scrap]:
        return self.db.query(Scrap).filter(Scrap.id == scrap_id).first()

    def count_scrap(self) -> int:
        return self.db.query(Scrap).count()

    def delete_scrap(self, scrap_id: int) -> Optional[Scrap]:
        scrap = self.db.query(Scrap).filter(Scrap.id == scrap_id).first()
        if scrap is None:
            return None

        for image in scrap.images:
            self.db.delete(image)
        self.db.delete(scrap)
        self.db.commit()
        return scrap

    def get_tags(self, scrap_id: int = None, tag_name: str = None):
        query = self.db.query(Tags)
        if scrap_id is not None:
            query = query.filter_by(scrap_id=scrap_id)
        if tag_name is not None:
            query = query.filter_by(tag_name=tag_name)
        return query.distinct(tag_name).all()
