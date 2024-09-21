from bs4 import BeautifulSoup
from playwright.async_api import async_playwright

IMAGE_EXTENSIONS = (".jpg", ".jpeg", ".png", ".gif", ".webp", ".svg")


# 네트워크 요청을 필터링하는 함수
def filter_requests(block_urls: list[str] = None):
    # 요청 URL이 이미지 파일 확장자 중 하나로 끝나는지 확인
    async def filter(route, request):
        req: str = request.url
        if any(req.endswith(ext) for ext in IMAGE_EXTENSIONS):
            return await route.abort()

        if block_urls is not None and any(req.startswith(url) for url in block_urls):
            return await route.abort()

        await route.continue_()

    return filter


async def load_soup(url, block_urls: list[str] = None):
    async with async_playwright() as p:
        browser = await p.chromium.launch(headless=False, slow_mo=0, args=["--disable-software-rasterizer"])
        context = await browser.new_context()
        page = await context.new_page()
        await page.route("**/*", filter_requests(block_urls))

        await page.goto(url)
        await page.wait_for_load_state("networkidle")

        content = await page.content()
        close = browser.close()
        soup = BeautifulSoup(content, features="html.parser")
        await close
        return soup
