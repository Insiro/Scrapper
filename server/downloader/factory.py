from .AbsDownloader import AbsDownloader
from .twitter import TwitterDownloader


class Scrapper(AbsDownloader):
    def getInstance(self, url: str = None, type: str = None) -> AbsDownloader:
        return TwitterDownloader()

    def scrap(self, url):
        return self.getInstance(url).scrap(url)
