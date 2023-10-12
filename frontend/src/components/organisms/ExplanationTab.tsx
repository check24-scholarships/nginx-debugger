"use client";

interface Props {
	explanation?: ExplanationResponse;
}

export default function ExplanationTab({ explanation }: Props) {
	const lines = explanation?.Explanation;
	if (!lines) return <div>No explanation fetched</div>;
}
