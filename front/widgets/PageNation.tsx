import { CSSProperties, FC, PropsWithChildren, useMemo } from "react";

const PageItemStyle: CSSProperties = {
    listStyle: "none",
    display: "inline-block",
    margin: "1rem",
    background: "rgba(0,0,0,0)",
    border: "0px",
};
interface ItemProps {
    current?: boolean;
    txt?: string | number;
    onClick: () => unknown;
}
const PageItem: FC<PropsWithChildren<ItemProps>> = ({ children, onClick, txt, current = false }) => (
    <button onClick={onClick} style={{ ...PageItemStyle, color: current ? "black" : "gray" }} disabled={current}>
        {txt ?? children}
    </button>
);

interface PageProps {
    totalPage: number;
    visiblePage: number;
    current: number;
    setPage: (page: number) => unknown;
}

export const PageNation: FC<PageProps> = ({ totalPage, visiblePage, current, setPage }: PageProps) => {
    const margin = useMemo(() => Math.ceil(visiblePage / 2), [visiblePage]);
    const startNum = useMemo(() => Math.max(1, current - margin), [current, margin]);

    const pageList = useMemo(() => Array.from({ length: Math.min(visiblePage, totalPage) }, (_, idx) => idx + startNum), [visiblePage, startNum, totalPage]);

    const lastPage = useMemo(() => pageList.at(-1) ?? totalPage, [pageList, totalPage]);
    const changePage = (page: number) => page !== current && setPage(page);

    return (
        <div style={{ justifyContent: "1", maxWidth: "100%" }}>
            {1 < startNum && <PageItem onClick={() => changePage(1)} txt="..." />}
            {pageList.map((it) => (
                <PageItem key={it} onClick={() => changePage(it)} current={it == current} txt={it} />
            ))}
            {totalPage > lastPage && <PageItem onClick={() => changePage(totalPage)} txt="..." />}
        </div>
    );
};
