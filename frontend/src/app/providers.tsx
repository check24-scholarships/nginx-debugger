"use client";
import { CacheProvider } from "@chakra-ui/next-js";
import { ChakraProvider } from "@chakra-ui/react";
import React from "react";
import theme from "@components/theme";

export function Providers({ children }: { children: React.ReactNode }) {
	return (
		<CacheProvider>
			<ChakraProvider theme={theme}>{children}</ChakraProvider>
		</CacheProvider>
	);
}
