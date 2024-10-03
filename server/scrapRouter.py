from typing import List, Literal

from fastapi import APIRouter, Depends

from .domain.dto import ImageDelete, ScrapResponse, ScrapUpdate
from .scrapController import ScrapAPIController, URLInput, get_scrap_api_controller

router = APIRouter()


# Define routes
@router.post("/scraps", response_model=ScrapResponse)
async def parse_url(
    url_input: URLInput,
    controller: ScrapAPIController = Depends(get_scrap_api_controller),
):
    return await controller.parse_url(url_input)


@router.get("/scraps", response_model=dict[Literal["list", "count"], List[ScrapResponse] | int])
async def list_scraps(
    page: int = 1,
    limit: int = 20,
    controller: ScrapAPIController = Depends(get_scrap_api_controller),
):
    return await controller.list_scraps(page=page, limit=limit)


@router.get("/scraps/count", response_model=int)
async def list_scraps(
    controller: ScrapAPIController = Depends(get_scrap_api_controller),
):
    return await controller.count_scraps()


@router.delete("/scraps/{scrap_id}", response_model=ScrapResponse)
async def delete_scrap(scrap_id: int, controller: ScrapAPIController = Depends(get_scrap_api_controller)):
    return await controller.delete_scrap(scrap_id)


@router.patch("/scrap/{scrap_id}")
async def update_script(
    scrap_id: int,
    update_scrap: ScrapUpdate,
    controller: ScrapAPIController = Depends(get_scrap_api_controller),
):
    return await controller.update_scrap(scrap_id, update_scrap)


@router.post("/scraps/{scrap_id}", response_model=ScrapResponse)
async def re_scrap(scrap_id: int, controller: ScrapAPIController = Depends(get_scrap_api_controller)):
    return await controller.re_scrap(scrap_id)


@router.get("/scraps/{scrap_id}", response_model=ScrapResponse)
async def get_scrap_details(scrap_id: int, controller: ScrapAPIController = Depends(get_scrap_api_controller)):
    return await controller.get_scrap_details(scrap_id)


@router.delete("/images")
async def delete_images(input: ImageDelete, controller: ScrapAPIController = Depends(get_scrap_api_controller)):
    return await controller.delete_images(input)
