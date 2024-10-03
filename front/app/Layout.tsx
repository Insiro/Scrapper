import React, { CSSProperties } from "react";
import { Outlet } from "react-router-dom";

import PageTitle from "@/widgets/PageTitle";
import { useTitleContext } from "@/entities/title";

// 스타일 정의
const layoutStyles: Record<string, CSSProperties> = {
    container: {
        display: "flex",
        flexDirection: "column" as const,
        minHeight: "100vh",
    },
    content: {
        flex: 1,
        padding: "3rem",
    },
};

const Layout: React.FC = () => {
    const { setPageTitle } = useTitleContext();
    return (
        <div style={layoutStyles.container}>
            <main style={layoutStyles.content}>
                <PageTitle />
                <Outlet context={{ setPageTitle }} />
            </main>
        </div>
    );
};

export default Layout;
