from urllib.parse import SplitResult
from uuid import uuid4

from server.scrapper.AbsScrapper import ScrapArgs
from server.utils import download_image, load_soup

from .AbsScrapper import AbsScrapper
from ..domain.PageType import PageType


class ImplInsta(AbsScrapper):
    def __init__(self) -> None:
        super().__init__()
        self.pageType = PageType.instagram

    def gen_args(self, url: SplitResult) -> ScrapArgs:
        args = super().gen_args(url)
        args.key = url.path.split("/p/")[1].split("/")[0]
        return args

    async def scrap(self, args):
        soup = await load_soup(args.url)

        meta_list = soup.find_all("meta")
        generator = filter(lambda x: x.get("property") == "instapp:owner_user_id", meta_list)
        use_id = next(generator).get("content")

        wrapper = soup.find("article").find("div").find_all("div", recursive=False)

        image_tag_list = wrapper[0].find_all("img")
        img_src_list = set([img["src"] for img in image_tag_list])

        authorTag = wrapper[1].find("header").find_all("div", recursive=False)[1].find_all("div", recursive=False)
        name = authorTag[0].find_next("div").find_next("div")

        content = wrapper[1].find("section").find_next_siblings("div")[0]
        content = content.find("ul").find("div").find("div").find("h2").next_sibling

        contentTxt = content.text

        fname_list = []
        for src in img_src_list:
            fname = download_image(src, f"{uuid4()}")
            if fname is not None:
                fname_list.append(fname)

        return self._create_scrap(use_id, name.text, args.key, contentTxt, fname_list)


if __name__ == "__main__":
    tweet_url = input("Enter the Twitter URL: ")

    downloader = ImplInsta()
    result = downloader.scrap(tweet_url)
    print(result)
