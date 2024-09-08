from urllib.parse import urlsplit, urlunsplit
from uuid import uuid4

from src.domain.dto import ScrapCreate
from src.utils.bsLoader import loadSoup
from src.utils.saveImg import download_image

from .AbsDownloader import AbsDownloader


class TwitterDownloader(AbsDownloader):
    async def scrap(self, url):
        split_url = urlsplit(url)
        url = urlunsplit((split_url.scheme, split_url.netloc, split_url.path, "", split_url.fragment))
        soup = await loadSoup(url)
        article = soup.find("article")
        author = article.find("div", {"data-testid": "User-Name"}).find_all("a")
        content = article.find("div").find("div").find_all("div", recursive=False)[2]
        img_list = content.find_all("img", recursive=True)

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
            image_names=fname_list,
            content=content.get_text(),
        )

    def preprocess_url(self, url: str) -> str:
        # Twitter의 경우 쿼리 파라미터를 제거합니다.
        split_url = urlsplit(url)
        return urlunsplit((split_url.scheme, split_url.netloc, split_url.path, "", split_url.fragment))


if __name__ == "__main__":
    tweet_url = input("Enter the Twitter URL: ")

    downloader = TwitterDownloader()
    result = downloader.scrap(tweet_url)
    print(result)
