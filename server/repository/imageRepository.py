from os import path, remove, rename
from typing import List

from sqlalchemy.orm import Session

from server.domain.entity import Image
from server.utils.config import Config  # 모델 정의가 포함된 모듈


class ImageRepository:
    def __init__(self, db: Session, config: Config):
        self.db = db
        self.media_path = config.media

    def save_images(self, files: List[str], scrap_id: int) -> None:
        """
        파일 목록을 저장합니다.

        :param files: 저장할 파일의 경로 목록
        :param scrap_id: 해당 스크랩의 ID
        """
        images = []
        for file_name in files:
            # 데이터베이스에 파일 정보를 저장합니다.
            images.append(Image(file_name=file_name, scrap_id=scrap_id))

        self.db.add_all(images)

        self.db.commit()

    def delete_images(self, *image_id: int) -> None:
        """
        파일 목록을 삭제합니다.

        :param image_id: 이미지의 아이디
        """
        images = self.db.query(Image).filter(Image.id.in_(image_id)).all()

        for img in images:
            file = path.join(self.media_path, img.file_name)
            if path.exists(file):
                remove(file)
                # 데이터베이스에서 파일 정보를 삭제합니다.

        self.db.query(Image).filter(Image.id.in_(image_id)).delete()
        self.db.commit()
        return True

    def delete_images_by_scrap(self, scrap_id: int) -> None:
        """
        특정 스크랩 ID에 연결된 모든 이미지를 삭제합니다.

        :param scrap_id: 스크랩의 ID
        """
        images = self.db.query(Image).filter(Image.scrap_id == scrap_id).all()
        for image in images:
            file = path.join(self.media_path, image.file_name)
            print(file)
            if path.exists(file):
                remove(file)

        image_ids = [img.id for img in images]

        self.db.query(Image).filter(Image.id.in_(image_ids)).delete()
        self.db.commit()
        return True

    def export_and_delete(self, image_list: List[Image]) -> None:
        """
        지정된 이미지 ID 목록에 해당하는 이미지를 이동 후 제거 합니다.

        :param image_id_list: 이동할 이미지의 ID 목록
        """
        for image in image_list:
            file_name = image.file_name
            source = path.join(self.storage_path, file_name)
            destination = path.join(self.export_path, file_name)
            if path.exists(source):
                rename(source, destination)

        image_ids = [img.id for img in image_list]

        self.db.query(Image).filter(Image.id.in_(image_ids)).delete()
        self.db.commit()

    def get_image(self, image_name_list) -> list[Image]:
        images = self.db.query(Image).filter(Image.file_name.in_(image_name_list))
        return images
