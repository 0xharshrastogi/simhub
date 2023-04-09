import classNames from "classnames";
import { FC } from "react";
import "./button.scss";

export type ButtonVariant = "default" | "secondary";

export type ButtonSize = "normal" | "large" | "small" | "tiny";

export type ButtonProps = {
	square?: boolean;
	variant?: ButtonVariant;
	size?: ButtonSize;
};

type Props = ButtonProps &
	Omit<
		React.DetailedHTMLProps<
			React.ButtonHTMLAttributes<HTMLButtonElement>,
			HTMLButtonElement
		>,
		"className"
	>;

export const Button: FC<Props> = (props) => {
	const { square, variant, size, ...rest } = props;
	const state = {
		canApplyVariant: !!variant && variant !== "default",
		canApplySize: !!size && size !== "normal",
	};

	return (
		<button
			className={classNames("button", {
				["square"]: props.square,
				[variant ?? ""]: state.canApplyVariant,
				[size ?? ""]: state.canApplySize,
			})}
			{...rest}
		>
			{props.children ?? "<NO LABEL>"}
		</button>
	);
};
