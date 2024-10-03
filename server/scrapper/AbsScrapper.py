from abc import ABCMeta, abstractmethod
from urllib.parse import SplitResult, urlsplit, urlunsplit

from server.domain.dto import ScrapCreate
from server.domain.PageType import PageType


class ScrapArgs:
    url: str
    key: str

    def __init__(self, url, key=None) -> None:
        self.url = url
        self.key = key


class AbsScrapper(metaclass=ABCMeta):
    pageType: PageType = None

    @abstractmethod
    async def scrap(self, args: ScrapArgs) -> ScrapCreate:
        pass

    def preprocess_url(self, url: str):
        split = urlsplit(url)._replace(query="")
        return split

    def gen_args(self, url: SplitResult) -> ScrapArgs:
        return ScrapArgs(url=self.merge_url(url))

    def merge_url(self, url: SplitResult):
        return urlunsplit((url.scheme, url.netloc, url.path, url.query, url.fragment))

    def _create_scrap(
        self, author_tag: str, author_name: str, source_key: str, content: str, image_names: list[str]
    ) -> ScrapCreate:
        # region process author info
        if author_name.startswith("@"):
            author_name = author_name[1:]
        if author_tag.startswith("@"):
            author_tag = author_tag[1:]
        # endregion
        filtered = filter(lambda x: x.startswith("#"), content.replace("\n", " ").split(" "))
        tags = list([tag[1:] for tag in filtered])

        return ScrapCreate(
            author_name=author_name,
            author_tag=author_tag,
            source_key=source_key,
            source=self.pageType,
            image_names=image_names,
            content=content,
            tags=tags,
        )
