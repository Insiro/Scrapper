import json
from urllib.parse import urlsplit, urlunsplit
from uuid import uuid4

import requests

from server.domain.dto import ScrapCreate
from server.utils import download_image

from .PageType import PageType
from .Scrapper import Scrapper


class ImplHoyolab(Scrapper):
    async def scrap(self, url):
        post_id = urlsplit(url).path.split("/")[-1]

        apiPath = f"https://bbs-api-os.hoyolab.com/community/post/wapi/getPostFull?post_id={post_id}"
        req = requests.get(
            apiPath,
            headers={
                "x-rpc-language": "ko-kr",
                "x-rpc-timezone": "Asia/Seoul",
                "x-rpc-show-translated": "true",
            },
        )
        data = req.json()
        post = data["data"]["post"]

        userData = post["user"]
        postData = post["post"]

        title = postData["subject"]

        contentData = json.loads(postData["content"])
        img_list = contentData["imgs"]
        content = contentData["describe"]

        fname_list = []
        for img in img_list:
            fname = download_image(img, f"{uuid4()}")
            if fname is not None:
                fname_list.append(fname)
        print(f"{title}\n{content}")
        return ScrapCreate(
            author_name=userData["nickname"],
            author_tag=userData["uid"],
            url=url,
            source=PageType.hoyolab,
            image_names=fname_list,
            content=f"{title}\n{content}",
        )

    def preprocess_url(self, url: str) -> str:
        split_url = urlsplit(url)
        return urlunsplit((split_url.scheme, split_url.netloc, split_url.path, "", split_url.fragment))


if __name__ == "__main__":
    tweet_url = input("Enter the Twitter URL: ")

    downloader = ImplHoyolab()
    result = downloader.scrap(tweet_url)
    print(result)
