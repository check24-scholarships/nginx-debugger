"use client";

import { useEffect, useRef, useState } from "react";
import { EditorView, basicSetup } from "codemirror";
import { StreamLanguage } from "@codemirror/language";
import { nginx } from "@codemirror/legacy-modes/mode/nginx";
import { Container, Divider, Flex, Heading } from "@chakra-ui/react";

interface EditorProps {
	onSave(value: string): void;
}

export default function Editor(props: EditorProps) {
	const containerRef = useRef<HTMLDivElement>(null);

	const [view, setView] = useState<EditorView>();

	useEffect(() => {
		if (!containerRef.current) return;

		const theme = EditorView.theme({
			"&": {
				position: "absolute !important",
				width: "100%",
				height: "100%",
			},
		});

		const view = new EditorView({
			extensions: [basicSetup, StreamLanguage.define(nginx), theme],
			parent: containerRef.current,
		});

		setView(view);

		return () => {
			view.destroy();
			setView(undefined);
		};
	}, [containerRef]);

	const onSave = () => {
		if (view) props.onSave(view.state.doc.toString());
	};

	return (
		<Flex direction="column" height="100%">
			<Flex>
				<Heading>Editor</Heading>
				<Flex flex={1} justifyContent="flex-end">
					<div onClick={onSave}>s</div>
				</Flex>
			</Flex>
			<Divider />
			<Container
				flex={1}
				ref={containerRef}
				position="relative"></Container>
		</Flex>
	);
}
