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

export default function Home() {
	return (
		<Grid templateColumns="repeat(2, 1fr)" gap={6}>
			<GridItem>Config view</GridItem>
			<Tabs>
				<TabList>
					<Tab>Explanation</Tab>
					<Tab>Requets</Tab>
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
