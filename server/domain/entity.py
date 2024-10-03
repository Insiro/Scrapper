import enum
from typing import List

from sqlalchemy import (
    Boolean,
    Column,
    Enum,
    ForeignKey,
    Integer,
    String,
    Text,
    UniqueConstraint,
)
from sqlalchemy.orm import Mapped, relationship

from server.database import Base


class Scrap(Base):
    __tablename__ = "scraps"
    __table_args__ = (UniqueConstraint("source_id", "source"),)

    id = Column(Integer, primary_key=True, index=True)
    pin = Column(Boolean, default=False, index=True)
    source_id = Column(String(255), unique=True, index=True)

    source = Column(String(50))
    content = Column(Text, nullable=True)
    author_name = Column(String(100))
    author_tag = Column(String(100))
    images: Mapped[List["Image"]] = relationship("Image", back_populates="scrap", uselist=True)
    comment = Column(String(255), nullable=True)


class Image(Base):
    __tablename__ = "images"

    id = Column(Integer, primary_key=True, index=True)
    file_name = Column(String(255))
    scrap_id = Column(Integer, ForeignKey("scraps.id"))

    scrap: Mapped[Scrap] = relationship("Scrap", back_populates="images")


class ExportMode(enum.Enum):
    FULL = ("FULL",)
    IMG = ("IMG",)
    TXT = "TXT"


class Export(Base):
    __tablename__ = "export"
    id = Column(Integer, primary_key=True, index=True)
    title = Column(String(255))
    mode: Enum = Enum(ExportMode, metadata=Base.metadata, validate_strings=True)
    name_rule = Column(String(255))


class Tags(Base):
    __tablename__ = "tags"
    __table_args__ = (UniqueConstraint("name", "scrap_id"),)

    name = Column(String(200), index=True)
    scrap_id = Column(Integer, ForeignKey("scraps.id", ondelete="CASCADE"), index=True)
    __mapper_args__ = {"primary_key": [name, scrap_id]}
