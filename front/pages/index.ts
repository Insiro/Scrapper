export { NotFoundPage } from "./NotFound"
export { ScrapDetailPage } from "./ScrapDetail/ScrapDetailPage"
export { HomePage } from "./HomePage/HomePage"

import { homePageLoader } from "./HomePage/loader"
import { scrapDetailLoader } from "./ScrapDetail/loader"

export const loader = { home: homePageLoader, scrap: scrapDetailLoader }