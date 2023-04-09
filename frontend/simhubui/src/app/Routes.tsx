import { Navigate, RouteObject } from "react-router-dom";
import { Encyclopedia } from "../containers";

export const routes: RouteObject[] = [
	{ path: "", element: <Navigate to="encyclopedia" /> },
	{ path: "home", element: null },
	{ path: "encyclopedia", element: <Encyclopedia /> },
];
