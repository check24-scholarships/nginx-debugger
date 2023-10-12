import { promises as fs } from "fs";
import { NextResponse } from "next/server";

export const POST = async () => {
	const file = await fs.readFile("./explanationResponse.json", "utf8");
	const data = JSON.parse(file);

	return NextResponse.json(data);
};
