from typing import List, Literal

from fastapi import APIRouter, Depends

from server.controller.scrapController import ScrapAPIController
from server.domain.dto.scrap import ScrapResponse, ScrapUpdate, URLInput

scrap_router = APIRouter()


# Define routes
@scrap_router.post("/", response_model=ScrapResponse)
async def parse_url(
    url_input: URLInput,
    controller: ScrapAPIController = Depends(ScrapAPIController.depends),
):
    return await controller.parse_url(url_input)


@scrap_router.get("/", response_model=dict[Literal["list", "count"], List[ScrapResponse] | int])
async def list_scraps(
    page: int = 1,
    limit: int = 20,
    pined=False,
    controller: ScrapAPIController = Depends(ScrapAPIController.depends),
):
    offset = (page - 1) * limit
    return await controller.list_scraps(offset=offset, limit=limit, pined=pined)


@scrap_router.delete("/{scrap_id}", response_model=ScrapResponse)
async def delete_scrap(scrap_id: int, controller: ScrapAPIController = Depends(ScrapAPIController.depends)):
    return await controller.delete_scrap(scrap_id)


@scrap_router.patch("/{scrap_id}")
async def update_script(
    scrap_id: int,
    update_scrap: ScrapUpdate,
    controller: ScrapAPIController = Depends(ScrapAPIController.depends),
):
    return await controller.update_scrap(scrap_id, update_scrap)


@scrap_router.post("/{scrap_id}", response_model=ScrapResponse)
async def re_scrap(scrap_id: int, controller: ScrapAPIController = Depends(ScrapAPIController.depends)):
    return await controller.re_scrap(scrap_id)


@scrap_router.get("/{scrap_id}", response_model=ScrapResponse)
async def get_scrap_details(scrap_id: int, controller: ScrapAPIController = Depends(ScrapAPIController.depends)):
    return await controller.get_scrap_details(scrap_id)
