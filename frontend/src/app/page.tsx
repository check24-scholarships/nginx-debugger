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
import { usePostRequest } from "@lib/request/clientRequest";
import ExplanationTab from "@components/organisms/ExplanationTab";
import { useEffect } from "react";

export default function Home() {
	const { data, send } = usePostRequest<ExplanationResponse>(
		"http://localhost:9000/analyze",
	);

	useEffect(() => {
		send({ config: "some data" });
	}, []);

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
					<TabPanel>
						<ExplanationTab explanation={data} />
					</TabPanel>
					<TabPanel>
						<RequestTab />
					</TabPanel>
				</TabPanels>
			</Tabs>
		</Grid>
	);
}
