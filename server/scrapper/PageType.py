from enum import Enum


class PageType(Enum):
    twitter = "twitter"
    hoyolab = "hoyolab"
    instagram = "instagram"

    @staticmethod
    def from_str(host_name: str):
        match host_name:
            case "x.com" | "twitter.com":
                return PageType.twitter
            case "www.hoyolab.com":
                return PageType.hoyolab
            case "www.instagram.com":
                return PageType.instagram
