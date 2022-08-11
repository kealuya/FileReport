import {http} from "~/http";

export const callGetCaptcha = async (phoneNumber: string): Promise<HttpResponse> => {
    // 后台请求验证码,验证码无需返回，是需要在后台缓存中等待提交
    return await http.request<HttpResponse>("post", "/getCaptcha", {data: {phoneNumber}})
}
export const callLoginWithCaptcha = async (phoneNumber: string, captcha: string) => {

    return await http.request<HttpResponse>("post", "/loginFr", {data: {phoneNumber, captcha}})
}
