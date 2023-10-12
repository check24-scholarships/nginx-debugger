import { useState } from "react";
import { Box, Button, Flex, Stack, Text } from "@chakra-ui/react";

interface Props<T> {
	options: { value: T; name: string }[];
	defaultOption: T;
	onSelect: (item: T) => void;
}

export default function Select<T>({
	options,
	defaultOption,
	onSelect,
}: Props<T>) {
	const [open, setOpen] = useState(false);
	const [selected, setSelected] = useState<T>(defaultOption);

	const select = (newSelected: T) => {
		setSelected(newSelected);
		onSelect(newSelected);
	};
	const toggleOpen = () => {
		setOpen(!open);
	};

	return (
		<Flex className="relative">
			<Button onClick={toggleOpen}>
				<Stack direction="row" alignItems="center" gap={1}>
					<Text>
						{
							options.find((option) => option.value === selected)
								?.name
						}
					</Text>
				</Stack>
				{open && (
					<Stack
						position="absolute"
						zIndex={10}
						boxShadow="lg"
						top="100%"
						width="max-content"
						align="flex-start"
						borderRadius="lg"
						backgroundColor="white"
						padding={2}
						spacing={2}
						marginTop={2}>
						{options.map(({ value, name }) => (
							<Box
								key={`${value}`}
								padding={1.5}
								borderRadius="md"
								width="full"
								textAlign="start"
								color="black"
								onClick={() => select(value)}
								_hover={{
									backgroundColor: "brand.500",
									color: "white",
								}}>
								<Text>{name}</Text>
							</Box>
						))}
					</Stack>
				)}
			</Button>
		</Flex>
	);
}
