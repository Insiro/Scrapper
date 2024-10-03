export { NotFoundPage } from "./NotFound"
export { ScrapDetailPage } from "./ScrapDetail/ScrapDetailPage"
export { HomePage } from "./HomePage/HomePage"

import { scrapListLoader } from "./HomePage/loader"
import { scrapDetailLoader } from "./ScrapDetail/loader"

export const loader = { list: scrapListLoader, scrap: scrapDetailLoader }