"use client";

import { Text } from "@chakra-ui/react";

interface Props {
	explanation?: ExplanationResponse;
}

export default function ExplanationTab({ explanation }: Props) {
	const lines = explanation?.Explanation;
	if (!lines) return <div>No explanation fetched</div>;

	return Object.keys(lines).map((lineNumber) => {
		const lineExplanation = lines[lineNumber];

		return <Text key={lineNumber}>{lineExplanation}</Text>;
	});
}
