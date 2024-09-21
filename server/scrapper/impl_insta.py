from urllib.parse import urlsplit, urlunsplit
from uuid import uuid4

from server.domain.dto import ScrapCreate
from server.utils import download_image, load_soup

from .PageType import PageType
from .Scrapper import Scrapper


class ImplInsta(Scrapper):
    async def scrap(self, url):
        split_url = urlsplit(url)
        url = urlunsplit((split_url.scheme, split_url.netloc, split_url.path, "", split_url.fragment))
        soup = await load_soup(url)

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

        fname_list = []
        for src in img_src_list:
            fname = download_image(src, f"{uuid4()}")
            if fname is not None:
                fname_list.append(fname)

        return ScrapCreate(
            author_name=name.text,
            author_tag=use_id,
            url=url,
            source=PageType.instagram,
            image_names=fname_list,
            content=content.text,
        )

    def preprocess_url(self, url: str) -> str:
        split_url = urlsplit(url)
        return urlunsplit((split_url.scheme, split_url.netloc, split_url.path, "", split_url.fragment))


if __name__ == "__main__":
    tweet_url = input("Enter the Twitter URL: ")

    downloader = ImplInsta()
    result = downloader.scrap(tweet_url)
    print(result)
