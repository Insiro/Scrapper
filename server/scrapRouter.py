from typing import List

from fastapi import APIRouter, Depends

from .domain.dto import ScrapResponse, ScrapUpdate
from .scrapController import (ScrapAPIController, URLInput,
                              get_scrap_api_controller)

router = APIRouter()


# Define routes
@router.post("/scraps", response_model=ScrapResponse)
async def parse_url(
    url_input: URLInput,
    controller: ScrapAPIController = Depends(get_scrap_api_controller),
):
    return await controller.parse_url(url_input)


@router.get("/scraps", response_model=List[ScrapResponse])
async def list_scraps(
    skip: int = 0,
    limit: int = 10,
    controller: ScrapAPIController = Depends(get_scrap_api_controller),
):
    return await controller.list_scraps(skip=skip, limit=limit)


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


@router.get("/scraps/{scrap_id}", response_model=ScrapResponse)
async def get_scrap_details(scrap_id: int, controller: ScrapAPIController = Depends(get_scrap_api_controller)):
    return await controller.get_scrap_details(scrap_id)
