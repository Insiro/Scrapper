from typing import List
from sqlalchemy import Column, ForeignKey, Integer, String, Text
from sqlalchemy.orm import relationship, Mapped

from src.database import Base


class Image:
    pass


class Scrap(Base):
    __tablename__ = "scraps"

    id = Column(Integer, primary_key=True, index=True)
    url = Column(String, unique=True, index=True)
    source = Column(String)
    content = Column(Text, nullable=True)
    author_name = Column(String)
    author_tag = Column(String)
    images: Mapped[List["Image"]] = relationship("Image", back_populates="scrap", uselist=True)
    comment = Column(String, nullable=True)

    def __repr__(self):
        return f"<Scrap(id={self.id}, url={self.url}, author_info={self.author_name}@{self.author_tag})>"


class Image(Base):
    __tablename__ = "images"

    id = Column(Integer, primary_key=True, index=True)
    file_name = Column(String)
    scrap_id = Column(Integer, ForeignKey("scraps.id"))

    scrap: Mapped[Scrap] = relationship("Scrap", back_populates="images")
