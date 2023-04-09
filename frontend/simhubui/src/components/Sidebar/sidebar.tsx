import { FC } from "react";
import { Link } from "react-router-dom";
import "./sidebar.scss";

const sidebarItems = [
	{
		label: "Home",
		icon: "bx bxs-home",
		to: "/home",
	},
	{
		label: "Encyclopedia",
		icon: "bx bxs-book",
		to: "/encyclopedia",
	},
	{
		label: "Game",
		icon: "bx bx-link-alt",
		to: "https://www.simcompanies.com/landscape/",
	},
];

export const Sidebar: FC<{}> = () => {
	return (
		<aside className="sidebar">
			<header>
				<Link to="/">
					<img
						src="https://d1fxy698ilbz6u.cloudfront.net/static/images/favico/favicon.e6cec2f1f142.ico"
						alt="icon"
					/>
				</Link>
			</header>
			<div className="content">
				<SidebarContent />
			</div>
		</aside>
	);
};

const SidebarContent = () => (
	<ul>
		{sidebarItems.map((item) => (
			<li key={item.label}>
				<Link className="content-item" to={item.to}>
					<span className="content-item-icon">
						<i className={item.icon}></i>
					</span>
					<span className="content-item-label">{item.label}</span>
				</Link>
			</li>
		))}
	</ul>
);
