package pageType

import "errors"

type PageType string

const (
    Twitter   PageType = "twitter"
    HoyoLab   PageType = "hoyolab"
    HoyoLink  PageType = "hoyolink"
    Instagram PageType = "instagram"
    Null      PageType = ""
)

func FromHost(hostName string) (PageType, error) {
    switch hostName {
    case "x.com":
        fallthrough
    case "twitter.com":
        return Twitter, nil
    case "www.instagram.com":
        return Instagram, nil
    case "www.hoyolab.com":
        return HoyoLab, nil
    case "hoyo.link":
        return HoyoLink, nil
    }
    return Null, errors.New("invalid host name")
}

func (p PageType) Url(souceId string) string {
    switch p {
    case Twitter:
        return "https://x.com/" + souceId
    case HoyoLab:
        return "https://www.hoyolab.com/article/" + souceId
    case Instagram:
        return "https://www.instagram.com/p/" + souceId
    default:
        return ""
    }
}
