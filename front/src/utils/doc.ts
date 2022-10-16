import {http} from "~/http";

export const callNewDoc = async (file: DocFile): Promise<HttpResponse> => {
    // 后台请求验证码,验证码无需返回，是需要在后台缓存中等待提交
    return await http.request<HttpResponse>("post", "/newDoc", {data: file})
}

export const callDocFileList = async (p: PagingInfo): Promise<HttpResponse> => {
    // 后台请求验证码,验证码无需返回，是需要在后台缓存中等待提交
    return await http.request<HttpResponse>("post", "/getDocFileList", {data: p})
}
