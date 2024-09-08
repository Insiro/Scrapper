import React from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider } from "react-router-dom";

import router from "./router"; // 방금 생성한 라우터 파일을 가져옵니다.
import { LayoutProvider } from "./context/layoutProvider";

const root = ReactDOM.createRoot(document.getElementById("root") as HTMLElement);

root.render(
    <React.StrictMode>
        <LayoutProvider>
            <RouterProvider router={router} />
        </LayoutProvider>
    </React.StrictMode>
);
