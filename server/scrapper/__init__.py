from typing import Optional
from urllib.parse import SplitResult, urlsplit

from server.domain.PageType import PageType
from .AbsScrapper import AbsScrapper
from .impl_hoyolab import ImplHoyolab, ImplHoyoLink
from .impl_insta import ImplInsta
from .impl_twitter import ImplTwitter


class Scrapper(AbsScrapper):
    __url: SplitResult
    instance: AbsScrapper

    def __init__(
        self,
        url: str = None,
        type: Optional[PageType] = None,
        type_name: str | None = None,
    ) -> AbsScrapper:
        if type is None:
            if type_name is None:
                type = PageType.from_str(urlsplit(url).hostname)
            else:
                type = PageType[type_name]
        self.pageType = type
        match type:
            case PageType.twitter:
                self.instance = ImplTwitter()
            case PageType.hoyolab:
                self.instance = ImplHoyolab()
            case PageType.hoyoLink:
                self.instance = ImplHoyoLink()
            case PageType.instagram:
                self.instance = ImplInsta()
        self.__url = self.instance.preprocess_url(url)
        self.args = self.instance.gen_args(self.__url)

    def preprocess_url(self, url: str):
        return self.instance.preprocess_url(url)

    def scrap(self, url: str = None):
        args = self.args
        if url is not None:
            spited = self.instance.preprocess_url(url)
            args = self.instance.gen_args(spited)
        return self.instance.scrap(args)

    @property
    def url(self):
        return self.merge_url(self.__url)
