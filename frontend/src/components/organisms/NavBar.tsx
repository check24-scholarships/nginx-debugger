import { Stack, Text } from "@chakra-ui/react";

export default function NavBar() {
	return (
		<Stack direction="row" width="full" p={4}>
			<Text fontWeight="bold">Nginx Debugger</Text>
		</Stack>
	);
}
