from typing import Optional
from urllib.parse import urlsplit

from .impl_hoyolab import ImplHoyolab
from .impl_insta import ImplInsta
from .impl_twitter import ImplTwitter
from .PageType import PageType
from .Scrapper import Scrapper


class ScrapperFactory(Scrapper):
    def getInstance(
        self,
        url: str = None,
        type: Optional[PageType] = None,
        type_name: str | None = None,
    ) -> Scrapper:
        if type is None:
            if type_name is None:
                type = PageType.from_str(urlsplit(url).hostname)
            else:
                type = PageType[type_name]

        match type:
            case PageType.twitter:
                return ImplTwitter()
            case PageType.hoyolab:
                return ImplHoyolab()
            case PageType.instagram:
                return ImplInsta()

    def scrap(self, url):
        return self.getInstance(url).scrap(url)
