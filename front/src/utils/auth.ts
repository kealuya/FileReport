const TokenKey = "authorized-token";

type paramsMapType = {
    name: string;
    expires: number;
    accessToken: string;
};
import Cookies from 'js-cookie';
// 获取token
export function getToken() {
    // 此处与TokenKey相同，此写法解决初始化时Cookies中不存在TokenKey报错
    return Cookies.get("authorized-token");
}

// 设置token以及过期时间（cookies ）
// 后端需要将用户信息和token以及过期时间都返回给前端，过期时间主要用于刷新token
export function setToken(data: any) {
    const {accessToken, expires, name} = data;
    // 提取关键信息进行存储
    const paramsMap: paramsMapType = {
        name,
        expires: 30,
        accessToken
    };
    const dataString = JSON.stringify(paramsMap);
    // useUserStoreHook().SET_TOKEN(accessToken);
    // useUserStoreHook().SET_NAME(name);
    expires > 0
        ? Cookies.set(TokenKey, dataString, {
            expires: expires / 86400000
        })
        : Cookies.set(TokenKey, dataString);
    sessionStorage.setItem(TokenKey, dataString);
}

// 删除token
export function removeToken() {
    Cookies.remove(TokenKey);
    sessionStorage.removeItem(TokenKey);
}

export const cookiesStorage: Storage = {
    key(index: number): string | null {
        return null;
    },
    length: 0,
    removeItem(key: string): void {
    },
    clear(): void {
    },
    setItem(key: string, value: string): void {
        Cookies.set(key, value, {expires: 30})
    },
    getItem(key): string {
        let s = Cookies.get(key)
        s ??= ""
        return s
    }
}
