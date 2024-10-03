from bs4 import BeautifulSoup
from playwright.async_api import async_playwright

IMAGE_EXTENSIONS = (".jpg", ".jpeg", ".png", ".gif", ".webp", ".svg")


async def filter_requests(route, request):
    req: str = request.url
    if any(req.endswith(ext) for ext in IMAGE_EXTENSIONS):
        return await route.abort()

    await route.continue_()


async def load_soup(url, block_urls: list[str] = None):
    async with async_playwright() as p:
        browser = await p.chromium.launch(headless=False, slow_mo=0, args=["--disable-software-rasterizer"])
        context = await browser.new_context(bypass_csp=True)
        page = await context.new_page()
        await page.set_viewport_size({"width": 2560, "height": 1440})
        await page.set_extra_http_headers(
            {
                "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"
            }
        )

        if block_urls:
            await page.route(block_urls + "*", filter_requests)
        await page.route("**/*", filter_requests)

        await page.goto(url)
        await page.wait_for_load_state("networkidle")

        content = await page.content()

        close = browser.close()
        soup = BeautifulSoup(content, features="html.parser")
        await close
        return soup
