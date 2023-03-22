import { useState } from "react";
import { Sidebar } from "../../components/sidebar";
import "./encyclopedia.css";

export const Encyclopedia = () => {
	const [open, setOpen] = useState(false);

	return (
		<>
			<section className="encyclopedia-container">
				<Sidebar />
				<div>content</div>
			</section>
		</>
	);
};
