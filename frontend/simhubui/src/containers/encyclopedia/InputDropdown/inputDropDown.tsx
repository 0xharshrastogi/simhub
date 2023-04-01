import classNames from "classnames";
import { FC, ReactNode, useState } from "react";
import "./inputDropDown.css";
import { resources } from "./resources";

type InputProps = React.DetailedHTMLProps<
	React.InputHTMLAttributes<HTMLInputElement>,
	HTMLInputElement
> & {
	isActive: boolean;
};

interface OptionsProps {
	isActive: boolean;

	children: ReactNode;
}

interface OptionProps<T> {
	children: string;

	value: T;

	onSelect?: (value: T) => void;
}

export const InputDropdown = () => {
	const [value, setValue] = useState("");
	const [isOptionsOpen, setIsOptionsOpen] = useState(false);

	const items = resources.filter((r) =>
		r.image.toString().toLowerCase().includes(value.toLowerCase())
	);

	return (
		<>
			<InputBox
				isActive={isOptionsOpen}
				type="text"
				placeholder="any value"
				value={value}
				onChange={(e) => setValue(e.target.value)}
				onFocus={() => setIsOptionsOpen(true)}
				onBlur={() => setIsOptionsOpen(false)}
			/>
			<Options isActive={isOptionsOpen}>
				{items.map((item) => (
					<Option
						key={item.kind}
						value={item}
						onSelect={(item) => setValue(item.image.toString())}
					>
						{item.image.toString()}
					</Option>
				))}
			</Options>
		</>
	);
};

export const InputBox: FC<InputProps> = (props) => {
	const { isActive, ...inputElmProps } = props;
	return (
		<input
			className={classNames("input-bx", { open: props.isActive })}
			type="text"
			{...inputElmProps}
		/>
	);
};

const Options: FC<OptionsProps> = (props) => {
	const { children } = props;
	return (
		<ul className={classNames("input-options", { open: props.isActive })}>
			{children}
		</ul>
	);
};

const Option: <T>(props: OptionProps<T>) => JSX.Element = (props) => {
	const { value } = props;
	return <li onClick={() => props.onSelect?.(value)}>{props.children}</li>;
};
