import { useEffect, useState } from "react";

type RequestMethod = "POST" | "GET";

/**
 * The `usePostRequest` hook can be used in ReactComponents to send data to an api and use the response data.
 *
 * `usePostRequest` is a generic function. It needs to know the type of the
 * response data. Pass it in angled brackets like this:
 *
 * ```
 * const {data} = usePostRequest<User>("/api/auth/login");
 * ```
 *
 * If no data type is given, it is assumed that the response data is an empty object
 *
 * It returns information about the request in an object (use object destructuring!)
 * https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Destructuring_assignment#object_destructuring
 *
 * Information returned:
 * - `error`: Request failed?
 * - `errorMessage`
 * - `loading`: request in progress?
 * - `data`: return data, might be null
 * - `send`: call this function to send the request. Pass the req body as parameter
 *
 * Example:
 *
 * ```
 * // somewhere inside a React Component
 * // only destructure the vars you actually need!
 * const {loading, error, errorMessage, data, send} = usePostRequest<User>("/api/auth/login");
 * ```
 */
export function usePostRequest<T = Record<string, never>>(url: string) {
	return useRequest<T>(url, "POST", {});
}

interface useGetOptions extends Options {
	dispatchImmediately?: boolean;
}

/**
 * The `useGetRequest` hook can be used in ReactComponents to retrieve data from an api and use the response data.
 *
 * `useGetRequest` is a generic function. It needs to know the type of the
 * response data. Pass it in angled brackets like this:
 *
 * ```
 * const {data} = useGetRequest<User>("/api/user");
 * ```
 *
 * If no data type is given, it is assumed that the response data is an empty object
 *
 * The request is sent immediately the first time the hook is called.
 *
 * It returns information about the request in an object (use object destructuring!)
 * https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Destructuring_assignment#object_destructuring
 *
 * Information returned:
 * - `error`: Request failed?
 * - `errorMessage`
 * - `loading`: request in progress?
 * - `data`: return data, might be null
 * - `send`: call this function to send the request.
 *
 * Example:
 *
 * ```
 * // somewhere inside a React Component
 * // only destructure the vars you actually need!
 * const {loading, error, errorMessage, data} = useGetRequest<User>("/api/user");
 * ```
 */
export function useGetRequest<T>(url: string, options: useGetOptions = {}) {
	const { send, ...values } = useRequest<T>(url, "GET", {
		caching: options.caching ?? false,
	});

	/* pass an empty array of dependencies to ensure that 
    the request is only send once, when the component mounts
    */
	useEffect(() => {
		if (options.dispatchImmediately !== false) {
			send();
		}
		// eslint-disable-next-line react-hooks/exhaustive-deps
	}, []);

	return { ...values, send };
}

interface Options {
	caching?: boolean;
}

function useRequest<T = {}>(
	url: string,
	method: RequestMethod,
	options: Options,
) {
	const [loading, setLoading] = useState(false);
	const [error, setError] = useState(false);
	const [errorMessage, setErrorMessage] = useState("");
	const [data, setData] = useState<T | undefined>(undefined);

	function send(body: Record<string, any> | null = null) {
		setLoading(true);
		/*
        A second error var has to be used to store whether the request failed or not.
        If the information is saved in the state "error" it wont have updated once the response body
        is parsed and the errorMessage wont be extracted
        */
		let reqError = false;
		/* 
        Chained callback functions must be used here instead of await.
        Otherwise the states (loading, error, data, etc..) wouldn't update correctly
        */

		if (url.startsWith("/")) {
			url = `${process.env.APP_URL}${url}`;
		}

		if (options.caching) {
			const cachedReq = sessionStorage.getItem(createHash(url, method));

			if (cachedReq) {
				setError(false);
				setData(JSON.parse(cachedReq));
				setLoading(false);
				return;
			}
		}

		fetch(url, {
			headers: {
				Accept: "application/json",
				"Content-Type": "application/json; charset=utf-8",
			},
			credentials: "same-origin",
			method: method,
			body: body ? JSON.stringify(body) : null,
		})
			.then((res) => {
				setError(!res.ok);
				reqError = !res.ok;
				return res.json();
			})
			.then((data) => {
				if (reqError) {
					/*
                    Surrounded with a try catch expr since it's not sure
                    if this value exists on the response data
                    */
					try {
						setErrorMessage(data.error);
					} catch (e) {}
				} else if (options.caching && data) {
					sessionStorage.setItem(
						createHash(url, method),
						JSON.stringify(data),
					);
				}

				setData(data);
				setLoading(false);
			});
	}

	return { loading, error, errorMessage, data, send };
}

function createHash(url: string, method: RequestMethod): string {
	return `${url}${method}`;
}
