from src.downloader.AbsDownloader import AbsDownloader
from src.utils.saveImg import download_image
from src.utils.bsLoader import loadSoup


class TwitterDownloader(AbsDownloader):
    def scrap(self, tweet_url):
        soup = loadSoup(tweet_url)
        article = soup.find("article")
        author = article.find("div", {"data-testid": "User-Name"}).find_all("a")
        content = article.find("div").find("div").find_all("div", recursive=False)[2]
        img_list = content.find_all("img", recursive=True)

        name = author[0].get_text()
        tag = author[1].get_text()
        fname_list = []

        for i, img in enumerate(img_list):
            fname = download_image(img.get("src"), "media", f"image_{i}.jpg")
            if fname is not None:
                fname_list.append(fname)
        return {"name": name, "tag": tag, "files": fname_list}


if __name__ == "__main__":
    tweet_url = input("Enter the Twitter URL: ")

    downloader = TwitterDownloader()
    result = downloader.scrap(tweet_url)
    print(result)
