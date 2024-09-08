from bs4 import BeautifulSoup
from playwright.async_api import async_playwright

IMAGE_EXTENSIONS = (".jpg", ".jpeg", ".png", ".gif", ".webp", ".svg")


# 네트워크 요청을 필터링하는 함수
def filter_requests(route, request):
    # 요청 URL이 이미지 파일 확장자 중 하나로 끝나는지 확인
    if any(request.url.endswith(ext) for ext in IMAGE_EXTENSIONS):
        route.abort()  # 이미지 요청을 취소
    else:
        route.continue_()  # 나머지 요청은 계속 진행


async def load_soup(url):
    async with async_playwright() as p:
        browser = await p.chromium.launch(headless=False, slow_mo=0, args=["--disable-software-rasterizer"])
        context = await browser.new_context()
        page = await context.new_page()
        page.route("**/*", filter_requests)

        await page.goto(url)
        await page.wait_for_load_state("networkidle")

        content = await page.content()
        soup = BeautifulSoup(content)
        browser.close()
        return soup
