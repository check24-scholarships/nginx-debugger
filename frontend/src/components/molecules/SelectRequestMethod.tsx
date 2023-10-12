import Select from "@components/molecules/Select";

export type RequestMethod = "GET" | "POST" | "PUT" | "DELETE";

export const defaultRequestMethod: RequestMethod = "GET";

interface Option {
	name: string;
	value: RequestMethod;
}

interface Props {
	onSelect: (item: RequestMethod) => void;
}

export default function SelectRequestMethod({ onSelect }: Props) {
	const options: Option[] = [
		{ name: "GET", value: "GET" },
		{ name: "POST", value: "POST" },
		{ name: "DELETE", value: "DELETE" },
		{ name: "PUT", value: "PUT" },
	];

	return (
		<Select<RequestMethod>
			options={options}
			defaultOption={defaultRequestMethod}
			onSelect={onSelect}
		/>
	);
}
