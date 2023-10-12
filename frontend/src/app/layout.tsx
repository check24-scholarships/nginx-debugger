// app/layout.tsx
import { Providers } from "./providers";
import React from "react";
import { Box } from "@chakra-ui/react";

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
						{children}
					</Box>
				</Providers>
			</body>
		</html>
	);
}
