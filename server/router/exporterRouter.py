from fastapi import APIRouter, Depends

from server.controller.exporterController import ExporterController
from server.domain.dto.exporter import *

exporter_router = APIRouter()


@exporter_router.delete("/")
async def delete_exporter(input: SelectExporter, controller: ExporterController = Depends(ExporterController.depends)):
    return await controller.delete_exporter(input)


@exporter_router.get("/")
async def get_exporter(input: SelectExporter, controller: ExporterController = Depends(ExporterController.depends)):
    return await controller.list_exporter(input)


@exporter_router.post("/")
async def create_exporter(input: CreateExporter, controller: ExporterController = Depends(ExporterController.depends)):
    return await controller.create_exporter(input)


@exporter_router.patch("/")
async def update_exporter(input: UpdateExporter, controller: ExporterController = Depends(ExporterController.depends)):
    return await controller.update_exporter(input)
