import { Navigate, RouteObject } from "react-router-dom";
import { Encyclopedia } from "../containers/encyclopedia";

export const routes: RouteObject[] = [
	{ path: "", element: <Navigate to="encyclopedia" /> },
	{ path: "encyclopedia", element: <Encyclopedia /> },
];
