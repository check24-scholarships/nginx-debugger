"use client";
import { Flex } from "@chakra-ui/react";
import SelectRequestMethod from "@components/molecules/SelectRequestMethod";

export default function RequestTab() {
	return (
		<Flex>
			<SelectRequestMethod onSelect={() => {}} />
		</Flex>
	);
}
