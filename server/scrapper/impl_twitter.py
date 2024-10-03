from urllib.parse import SplitResult
from uuid import uuid4

from server.scrapper.AbsScrapper import ScrapArgs
from server.utils import download_image, load_soup

from .AbsScrapper import AbsScrapper
from ..domain.PageType import PageType


class ImplTwitter(AbsScrapper):
    def __init__(self) -> None:
        super().__init__()
        self.pageType = PageType.twitter

    def gen_args(self, url: SplitResult) -> ScrapArgs:
        args = super().gen_args(url)
        path = url.path
        if path.startswith("/"):
            path = path[1:]
        if path.endswith("/"):
            path = path[:-1]
        args.key = path

    async def scrap(self, args):
        soup = await load_soup(args.url)
        article = soup.find("article")
        author = article.find("div", {"data-testid": "User-Name"}).find_all("a")
        wrapper = article.find("div").find("div").find_all("div", recursive=False)[2]
        img_list = wrapper.find_all("img", recursive=True)

        content = wrapper.find("div").find("div").find("div")
        contentTxt: str = content.text

        name = author[0].get_text()
        tag = author[1].get_text()
        fname_list = []

        for img in img_list:
            fname = download_image(img.get("src"), f"{uuid4()}")
            if fname is not None:
                fname_list.append(fname)

        return self._create_scrap(tag, name, args.key, contentTxt, fname_list)


if __name__ == "__main__":
    tweet_url = input("Enter the Twitter URL: ")

    downloader = ImplTwitter()
    result = downloader.scrap(tweet_url)
    print(result)
