export { NotFoundPage } from "./NotFound";
export { ScrapDetailPage } from "./ScrapDetail/ScrapDetailPage";
export { HomePage } from "./HomePage/HomePage";

import { homePageLoader } from "./HomePage/loader";
import { scrapDetailLoader } from "./ScrapDetail/loader";
import { scrapListLoader } from "./ScrapList/loader";

export const loader = { list: scrapListLoader, home: homePageLoader, scrap: scrapDetailLoader };
