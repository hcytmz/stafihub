export interface LedgerEraExchangeRate {
    denom?: string;
    /** @format int64 */
    era?: number;
    value?: string;
}
export interface LedgerExchangeRate {
    denom?: string;
    value?: string;
}
export declare type LedgerMsgAddNewPoolResponse = object;
export declare type LedgerMsgClearCurrentEraSnapShotsResponse = object;
export declare type LedgerMsgRemovePoolResponse = object;
export declare type LedgerMsgSetChainBondingDurationResponse = object;
export declare type LedgerMsgSetCommissionResponse = object;
export declare type LedgerMsgSetEraUnbondLimitResponse = object;
export declare type LedgerMsgSetInitBondResponse = object;
export declare type LedgerMsgSetLeastBondResponse = object;
export declare type LedgerMsgSetPoolDetailResponse = object;
export declare type LedgerMsgSetReceiverResponse = object;
export interface LedgerQueryEraExchangeRatesByDenomResponse {
    eraExchangeRates?: LedgerEraExchangeRate[];
}
export interface LedgerQueryExchangeRateAllResponse {
    exchangeRates?: LedgerExchangeRate[];
}
export interface LedgerQueryGetEraExchangeRateResponse {
    eraExchangeRate?: LedgerEraExchangeRate;
}
export interface LedgerQueryGetExchangeRateResponse {
    exchangeRate?: LedgerExchangeRate;
}
export interface ProtobufAny {
    "@type"?: string;
}
export interface RpcStatus {
    /** @format int32 */
    code?: number;
    message?: string;
    details?: ProtobufAny[];
}
export declare type QueryParamsType = Record<string | number, any>;
export declare type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;
export interface FullRequestParams extends Omit<RequestInit, "body"> {
    /** set parameter to `true` for call `securityWorker` for this request */
    secure?: boolean;
    /** request path */
    path: string;
    /** content type of request body */
    type?: ContentType;
    /** query params */
    query?: QueryParamsType;
    /** format of response (i.e. response.json() -> format: "json") */
    format?: keyof Omit<Body, "body" | "bodyUsed">;
    /** request body */
    body?: unknown;
    /** base url */
    baseUrl?: string;
    /** request cancellation token */
    cancelToken?: CancelToken;
}
export declare type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;
export interface ApiConfig<SecurityDataType = unknown> {
    baseUrl?: string;
    baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
    securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}
export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
    data: D;
    error: E;
}
declare type CancelToken = Symbol | string | number;
export declare enum ContentType {
    Json = "application/json",
    FormData = "multipart/form-data",
    UrlEncoded = "application/x-www-form-urlencoded"
}
export declare class HttpClient<SecurityDataType = unknown> {
    baseUrl: string;
    private securityData;
    private securityWorker;
    private abortControllers;
    private baseApiParams;
    constructor(apiConfig?: ApiConfig<SecurityDataType>);
    setSecurityData: (data: SecurityDataType) => void;
    private addQueryParam;
    protected toQueryString(rawQuery?: QueryParamsType): string;
    protected addQueryParams(rawQuery?: QueryParamsType): string;
    private contentFormatters;
    private mergeRequestParams;
    private createAbortSignal;
    abortRequest: (cancelToken: CancelToken) => void;
    request: <T = any, E = any>({ body, secure, path, type, query, format, baseUrl, cancelToken, ...params }: FullRequestParams) => Promise<HttpResponse<T, E>>;
}
/**
 * @title ledger/genesis.proto
 * @version version not set
 */
export declare class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
    /**
     * No description
     *
     * @tags Query
     * @name QueryGetEraExchangeRate
     * @summary Queries a list of getEraExchangeRate items.
     * @request GET:/stafiprotocol/stafihub/ledger/EraExchangeRate/{denom}/{era}
     */
    queryGetEraExchangeRate: (denom: string, era: number, params?: RequestParams) => Promise<HttpResponse<LedgerQueryGetEraExchangeRateResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryEraExchangeRatesByDenom
     * @summary Queries a list of eraExchangeRatesByDenom items.
     * @request GET:/stafiprotocol/stafihub/ledger/eraExchangeRatesByDenom/{denom}
     */
    queryEraExchangeRatesByDenom: (denom: string, params?: RequestParams) => Promise<HttpResponse<LedgerQueryEraExchangeRatesByDenomResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryExchangeRateAll
     * @summary Queries a list of exchangeRateAll items.
     * @request GET:/stafiprotocol/stafihub/ledger/exchangeRateAll
     */
    queryExchangeRateAll: (params?: RequestParams) => Promise<HttpResponse<LedgerQueryExchangeRateAllResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryGetExchangeRate
     * @summary Queries a list of getExchangeRate items.
     * @request GET:/stafiprotocol/stafihub/ledger/exchangerate/{denom}
     */
    queryGetExchangeRate: (denom: string, params?: RequestParams) => Promise<HttpResponse<LedgerQueryGetExchangeRateResponse, RpcStatus>>;
}
export {};
