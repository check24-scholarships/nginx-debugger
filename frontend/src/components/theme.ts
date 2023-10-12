import { extendTheme, theme, withDefaultColorScheme } from "@chakra-ui/react";

export default extendTheme(
	{
		colors: {
			brand: theme.colors.blue,
		},
	},
	withDefaultColorScheme({
		colorScheme: "brand",
	}),
);
