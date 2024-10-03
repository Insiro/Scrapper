import { createBrowserRouter, RouteObject } from "react-router-dom";

import Layout from "./Layout";
import { HomePage, loader, NotFoundPage, ScrapDetailPage } from "@/pages";
import Config from "@/shared/config";
import ScrapListPage from "@/pages/ScrapList/ScrapList";

// createBrowserRouter를 사용하여 라우터 구성
const router: RouteObject[] = [
    {
        path: "",
        element: <HomePage />,
        loader: loader.home,
    },
    {
        path: "scraps",
        element: <ScrapListPage />,
        loader: loader.list,
    },

    {
        path: "scraps/:scrapId",
        element: <ScrapDetailPage />,
        loader: loader.scrap,
    },
    {
        path: "*",
        element: <NotFoundPage />,
    },
];

const rootRouter = createBrowserRouter(
    [
        {
            children: router,
            element: <Layout />,
        },
    ],
    { basename: "/" + Config.basePath }
);

export default rootRouter;
