import { RouterProvider, createBrowserRouter } from "react-router-dom";
import { routes } from "./Routes";

const router = createBrowserRouter(routes);

export const App = () => {
	return (
		<>
			<RouterProvider router={router} />
		</>
	);
};
