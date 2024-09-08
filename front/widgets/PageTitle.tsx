import React, { Fragment } from "react";
import { Link, useLocation } from "react-router-dom";
import { Card } from "./Common/Card";
import { useTitleContext } from "@/entities/title";
import { color } from "@/shared/constant";

// 스타일 정의
const titleStyle = {
    fontSize: "1.5rem",
    fontWeight: "bold",
    marginBottom: "1rem",
};

const linkStyle = {
    textDecoration: "none",
    color: color.blue,
    fontSize: "0.9rem",
    marginRight: "0.5rem",
};

const separatorStyle = {
    color: color.gray2,
    margin: "0 0.5rem",
    fontSize: "0.9rem",
};

const PageTitle: React.FC = () => {
    const { pageTitle } = useTitleContext();
    const location = useLocation();
    const pathnames = location.pathname.split("/").filter((x) => x);

    // 경로에 따라 자동으로 타이틀 생성
    const breadcrumbs = pathnames.map((value, index) => {
        const to = `/${pathnames.slice(0, index + 1).join("/")}`;

        return index === pathnames.length - 1 ? (
            <Fragment key={index}>
                <span style={separatorStyle}>/</span>
                <span key={to} style={{ fontSize: "0.9rem", color: color.gray1 }}>
                    {value}
                </span>
            </Fragment>
        ) : (
            <Fragment key={index}>
                <span style={separatorStyle}>/</span>
                <Link key={to} to={to} style={linkStyle}>
                    {value.charAt(0).toUpperCase() + value.slice(1)}
                </Link>
            </Fragment>
        );
    });

    return (
        <Card style={{}}>
            <h1
                style={{
                    fontSize: "1.8rem",
                    fontWeight: "bold",
                    marginBottom: "1rem",
                }}
            >
                {pageTitle}
            </h1>
            <h1 style={titleStyle}>
                {/* "Home" 링크를 항상 첫 부분에 추가 */}
                <Link to="/" style={linkStyle}>
                    Home
                </Link>
                {breadcrumbs}
            </h1>
        </Card>
    );
};

export default PageTitle;
