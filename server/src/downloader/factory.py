from .AbsDownloader import AbsDownloader
from .twitter import TwitterDownloader


class Scrapper(AbsDownloader):
    def getInstance(self, url: str) -> AbsDownloader:
        return TwitterDownloader()

    def scrap(self, url):
        return self.getInstance(url).scrap(url)
