from fastapi import APIRouter

from .exporterRouter import exporter_router
from .imageRouter import image_router
from .scrapRouter import scrap_router

router = APIRouter()

router.include_router(image_router, prefix="/images")
router.include_router(scrap_router, prefix="/scraps")
router.include_router(exporter_router, prefix="/exporter")
