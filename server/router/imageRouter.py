from fastapi import APIRouter, Depends

from server.controller.scrapController import ScrapAPIController
from server.domain.dto.image import ImageDelete

image_router = APIRouter()


@image_router.delete("/")
async def delete_images(input: ImageDelete, controller: ScrapAPIController = Depends(ScrapAPIController.depends)):
    return await controller.delete_images(input)
