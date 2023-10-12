"use client";
import {
	Button,
	Stack,
	Tab,
	TabList,
	TabPanel,
	TabPanels,
	Tabs,
} from "@chakra-ui/react";
import SelectRequestMethod from "@components/molecules/SelectRequestMethod";

export default function RequestTab() {
	return (
		<Stack direction="column">
			<Stack direction="row" justifyContent="space-between">
				<SelectRequestMethod onSelect={() => {}} />
				<Button>SEND</Button>
			</Stack>
			<Tabs>
				<TabList>
					<Tab>Headers</Tab>
					<Tab>Body</Tab>
				</TabList>
				<TabPanels>
					<TabPanel>WIP</TabPanel>
					<TabPanel></TabPanel>
				</TabPanels>
			</Tabs>
		</Stack>
	);
}
