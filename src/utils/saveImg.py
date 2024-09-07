from io import BytesIO
from PIL import Image
import os
import requests


def download_image(
    url: str, foldername: str, filename: str, width_thresh=200, height_thresh=200
):
    # 이미지 URL을 다운로드하여 파일로 저장
    response = requests.get(url)
    if response.status_code != 200:
        return None
    try:
        image = Image.open(BytesIO(response.content))
        width, height = image.size
        if width < width_thresh or height < height_thresh:
            return None

        folder = os.path.join(os.getcwd(), foldername)
        fname = os.path.join(folder, filename)
        with open(fname, "wb") as file:
            file.write(response.content)
        return fname
    except:
        return None
