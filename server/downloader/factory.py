from enum import Enum
from typing import Literal, Optional
from urllib.parse import urlsplit

from .AbsDownloader import AbsDownloader
from .hoyolab import HoyolabDownloader
from .twitter import TwitterDownloader


class PageType(Enum):
    twitter = "twitter"
    hoyolab = "hoyolab"


class Scrapper(AbsDownloader):
    def getInstance(
        self,
        url: str = None,
        type: Optional[PageType] = None,
        type_name: str | None = None,
    ) -> AbsDownloader:
        if type is None:
            if type_name is None:
                match urlsplit(url).hostname:
                    case "x.com" | "twitter.com":
                        type = PageType.twitter
                    case "www.hoyolab.com":
                        type = PageType.hoyolab
            else:
                type = PageType[type_name]
        match type:
            case PageType.twitter:
                return TwitterDownloader()
            case PageType.hoyolab:
                return HoyolabDownloader()

    def scrap(self, url):
        return self.getInstance(url).scrap(url)
