// app/layout.tsx
import { Providers } from "./providers";
import React from "react";
import { Box } from "@chakra-ui/react";
import NavBar from "@components/organisms/NavBar";

export default function RootLayout({
	children,
}: {
	children: React.ReactNode;
}) {
	return (
		<html lang="en">
			<body>
				<Providers>
					<Box height="100vh" overflow="hidden">
						<NavBar />
						<Box mb={4} h={1} backgroundColor="blue.500" />
						<Box>{children}</Box>
					</Box>
				</Providers>
			</body>
		</html>
	);
}
