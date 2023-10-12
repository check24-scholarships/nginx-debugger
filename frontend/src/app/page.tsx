"use client";

import {
	Grid,
	GridItem,
	Tab,
	TabList,
	TabPanel,
	TabPanels,
	Tabs,
} from "@chakra-ui/react";
import RequestTab from "@components/organisms/RequestTab";
import Editor from "../components/organisms/Editor";

export default function Home() {
	return (
		<Grid templateColumns="repeat(2, 1fr)" gap={6} height="100vh">
			<GridItem>
				<Editor onSave={(v) => console.log(v)}></Editor>
			</GridItem>
			<Tabs>
				<TabList>
					<Tab>Explanation</Tab>
					<Tab>Requests</Tab>
				</TabList>
				<TabPanels>
					<TabPanel>WIP</TabPanel>
					<TabPanel>
						<RequestTab />
					</TabPanel>
				</TabPanels>
			</Tabs>
		</Grid>
	);
}
