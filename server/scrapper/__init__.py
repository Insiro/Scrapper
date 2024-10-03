from typing import Optional
from urllib.parse import SplitResult, urlsplit

from .AbsScrapper import AbsScrapper
from .impl_hoyolab import ImplHoyolab, ImplHoyoLink
from .impl_insta import ImplInsta
from .impl_twitter import ImplTwitter
from ..domain.PageType import PageType


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

    def preprocess_url(self, url: str):
        return self.instance.preprocess_url(url)

    def scrap(self, url: str = None):
        spited = None
        if url is None:
            spited = self.__url
        else:
            spited = self.instance.preprocess_url(url)
        args = self.instance.gen_args(spited)
        return self.instance.scrap(args)

    @property
    def url(self):
        return self.merge_url(self.__url)
