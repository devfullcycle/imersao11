import useSWR, { Fetcher, SWRConfiguration, SWRResponse } from "swr";

export type UseHttpResponse<Data = any, Error = any> = SWRResponse<Data, Error> & {
  isLoading: boolean;
};

export type UseHttpMetadata<Error = any> = Omit<UseHttpResponse<any, Error>, 'data'>;

export function useHttp<Data = any, Error = any>(
  url: string | null,
  fetcher: Fetcher<Data, any>,
  config?: SWRConfiguration<Data, Error, Fetcher<any, any>>
): UseHttpResponse<Data, Error> {
  const result = useSWR(url, fetcher, config);

  return {
    ...result,
    isLoading: !result.error && !result.data,
  };
}
