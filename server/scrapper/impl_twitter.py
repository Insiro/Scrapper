from urllib.parse import urlsplit, urlunsplit
from uuid import uuid4

from server.domain.dto import ScrapCreate
from server.utils import download_image, load_soup

from .PageType import PageType
from .Scrapper import Scrapper


class ImplTwitter(Scrapper):
    async def scrap(self, url):
        split_url = urlsplit(url)
        url = urlunsplit((split_url.scheme, split_url.netloc, split_url.path, "", split_url.fragment))

        soup = await load_soup(url)
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

        return ScrapCreate(
            author_name=name,
            author_tag=tag,
            url=url,
            source=PageType.twitter,
            image_names=fname_list,
            content=contentTxt,
            tags=self.extract_tags(content.text),
        )

    def preprocess_url(self, url: str) -> str:
        split_url = urlsplit(url)
        return urlunsplit((split_url.scheme, split_url.netloc, split_url.path, "", split_url.fragment))


if __name__ == "__main__":
    tweet_url = input("Enter the Twitter URL: ")

    downloader = ImplTwitter()
    result = downloader.scrap(tweet_url)
    print(result)
