from abc import ABCMeta, abstractmethod

from server.domain.dto import ScrapCreate


class AbsDownloader(metaclass=ABCMeta):

    @abstractmethod
    async def scrap(self, url: str) -> ScrapCreate:
        pass

    def preprocess_url(self, url: str) -> str:
        return url
