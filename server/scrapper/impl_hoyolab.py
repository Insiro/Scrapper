import json
from typing import Literal
from urllib.parse import SplitResult
from uuid import uuid4

import requests
from bs4 import BeautifulSoup

from server.utils import download_image
from server.utils.constant import USER_AGENT

from ..domain.PageType import PageType
from .AbsScrapper import AbsScrapper


class ImplHoyolab(AbsScrapper):
    def __init__(self) -> None:
        super().__init__()
        self.pageType = PageType.hoyolab

    def gen_args(self, url: SplitResult):
        args = super().gen_args(url)
        args.key = url.path.split("article")[-1][1:]
        return args

    async def scrap(self, args):
        apiPath = f"https://bbs-api-os.hoyolab.com/community/post/wapi/getPostFull?post_id={args.key}"
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
        return self._create_scrap(userData["uid"], userData["nickname"], args.key, f"{title}\n{content}", fname_list)


class ImplHoyoLink(ImplHoyolab):
    def preprocess_url(self, url: str):
        res = requests.get(url, headers={"User-Agent": USER_AGENT})
        meta_list = BeautifulSoup(res.content, features="html.parser").find_all("meta")
        item = None
        for it in meta_list:
            if it.get("content", "").startswith("0;url="):
                item = it
                break

        url = item.get("content")[7:-1]
        return super().preprocess_url(url)

    def gen_args(self, url) -> any:
        args = super().gen_args(url)
        args.key = url.fragment.split("article")[1].split("/")[1]
        return args


if __name__ == "__main__":
    tweet_url = input("Enter the Twitter URL: ")

    downloader = ImplHoyolab()
    result = downloader.scrap(tweet_url)
    print(result)
