from io import BytesIO
from mimetypes import guess_extension
from os import path

import requests
from PIL import Image

from .config import config


def download_image(url: str, filename: str, width_thr=200, height_thr=200):
    # 이미지 URL을 다운로드하여 파일로 저장
    response = requests.get(url)
    if response.status_code != 200:
        return None
    try:
        image = Image.open(BytesIO(response.content))
        width, height = image.size
        if width < width_thr or height < height_thr:
            return None

        content_type = response.headers.get("Content-Type", "")
        if (ext := guess_extension(content_type)) is None:  # 확장자를 결정하지 못하는 경우 기본 확장자 설
            ext = ".jpg"  # 기본값으로 '.jpg' 사용

        filename = f"{filename}{ext}"

        with open(path.join(config.media, filename), "wb") as file:
            file.write(response.content)

        return filename
    except Exception as e:
        return None
