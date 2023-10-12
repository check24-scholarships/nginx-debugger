import { promises as fs } from "fs";
import { NextResponse } from "next/server";

export async function POST() {
	const file = await fs.readFile(
		process.cwd() + "/src/app/api/mock/analyze/explanationResponse.json",
		"utf8",
	);
	const data = JSON.parse(file);

	return NextResponse.json(data);
}
