"use client";
import {
	Box,
	Stack,
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
import React, { useEffect } from "react";

export default function Home() {
	const { data, send } =
		usePostRequest<ExplanationResponse>("api/mock/analyze");

	useEffect(() => {
		send({ config: "some data" });
	}, []);

	console.log(data);

	return (
		<Stack direction="row" gap={6} height="100vh">
			<Box width="50%">
				<Editor onSave={(v) => console.log(v)}></Editor>
			</Box>
			<Tabs variant="solid-rounded">
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
		</Stack>
	);
}
