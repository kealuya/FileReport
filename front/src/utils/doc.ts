import {http} from "~/http";
import {

    PureHttpRequestConfig
} from "~/http/types";
import {useUserStore} from "~/stores";

const userStore = useUserStore()


export const callNewDoc = async (file: DocFile): Promise<HttpResponse> => {
    //
    return await http.request<HttpResponse>("post", "/newDoc", {data: file})
}

export const callUpdateDoc = async (file: DocFile): Promise<HttpResponse> => {
    //
    return await http.request<HttpResponse>("post", "/updateDoc", {data: file}, {},)
}

export const callDocFileList = async (p: PagingInfo): Promise<HttpResponse> => {
    //
    return await http.request<HttpResponse>("post", "/getDocFileList", {
        data: p,
        beforeRequestCallback: (request: PureHttpRequestConfig) => {
            request.headers = {
                token: userStore.token
            }
        }
    })
}

export const callAuthorityDoc = async (file: DocFile): Promise<HttpResponse> => {
    return await http.request<HttpResponse>("post", "/updateAuthorityDoc", {data: file})
}
