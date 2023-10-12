"use client";

import { useRef } from "react";
import { Container, Divider, Flex, Heading } from "@chakra-ui/react";
import Monaco from "@monaco-editor/react";
import "monaco-editor-nginx";

interface EditorProps {
	onSave(value: string): void;
}

export default function Editor(props: EditorProps) {
	const editorRef = useRef(null);
	const monacoRef = useRef(null);

	const onSave = () => {
		if (!editorRef.current) return;
		props.onSave(editorRef.current.getValue());
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
			<Container flex={1}>
				<Monaco
					height="100%"
					language="nginx"
					theme="nginx-theme"
					defaultValue="test"
					onMount={(editor, monaco) => {
						editorRef.current = editor;
						monacoRef.current = monaco;
					}}
				/>
			</Container>
		</Flex>
	);
}
