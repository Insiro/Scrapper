from enum import Enum


class PageType(Enum):
    twitter = "twitter"
    hoyolab = "hoyolab"
    hoyoLink = "hoyoLink"
    instagram = "instagram"

    @staticmethod
    def from_str(host_name: str):
        match host_name:
            case "x.com" | "twitter.com":
                return PageType.twitter
            case "www.hoyolab.com":
                return PageType.hoyolab
            case "hoyo.link":
                return PageType.hoyoLink
            case "www.instagram.com":
                return PageType.instagram

    def url(self, source_id: str):
        match self:
            case PageType.twitter:
                return "https://x.com/" + source_id
            case PageType.hoyolab:
                return "https://www.hoyolab.com/article/" + source_id
            case PageType.instagram:
                return "https://www.instagram.com/p/" + source_id
