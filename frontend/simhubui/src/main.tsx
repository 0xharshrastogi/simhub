import React from "react";
import ReactDOM from "react-dom/client";
import { RealmType, SimCompanies } from "simcompanies";
import { App } from "./app/App";
import "./index.css";

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
	<React.StrictMode>
		<App />
	</React.StrictMode>
);

const realm = SimCompanies.loadRealm(RealmType.Magnates);
console.log(realm.fetchResourceList().then(console.log).catch(console.log));
