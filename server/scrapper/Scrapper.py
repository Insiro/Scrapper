from abc import ABCMeta, abstractmethod

from server.domain.dto import ScrapCreate


class Scrapper(metaclass=ABCMeta):

    @abstractmethod
    async def scrap(self, url: str) -> ScrapCreate:
        pass

    def preprocess_url(self, url: str) -> str:
        return url

    def extract_tags(self, content: str) -> list[str]:
        filtered = filter(lambda x: x.startswith("#"), content.replace("\n", " ").split(" "))
        tags = list([tag[1:] for tag in filtered])

        return tags
