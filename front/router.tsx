import { createBrowserRouter } from "react-router-dom";
import HomePage from "./pages/HomePage/HomePage";
import ScrapDetailPage from "./pages/ScrapDetail/ScrapDetailPage";
import NotFoundPage from "./pages/NotFound";
import { homePageLoader } from "./pages/HomePage/loader";
import { scrapDetailLoader } from "./pages/ScrapDetail/loader";

// createBrowserRouter를 사용하여 라우터 구성
const router = createBrowserRouter([
    {
        path: "/",
        element: <HomePage />,
        loader: homePageLoader,
    },
    {
        path: "/scraps/:scrapId",
        element: <ScrapDetailPage />,
        loader: scrapDetailLoader,
    },
    {
        path: "*",
        element: <NotFoundPage />,
    },
]);

export default router;
