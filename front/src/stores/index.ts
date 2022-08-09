import {defineStore} from 'pinia'
import {cookiesStorage} from "~/utils/auth";
// useStore 可以是 useUser、useCart 之类的任何东西
// 第一个参数是应用程序中 store 的唯一 id
export const useUserStore = defineStore('user', {
    state: () => ({
        phoneNumber: "",
        userName: "",
        token: "",// 后台生成token
    }),
    getters: {
        isOwner(state) {
            return false
        },
    },
    actions: {
        getCaptcha(phoneNumber: string): boolean {
            // todo 后台请求验证码,验证码无需返回，是需要在后台缓存中等待提交

            console.log("发送captcha")
            return true
        },
        loginWithCaptche(phoneNumber: string, captcha: string) {
        },
        test(s: string) {
            this.token = s
        }
    },
    persist: {
        enabled: true,
        strategies: [
            {
                // storage: cookiesStorage,
                storage: localStorage,
            },
        ],
    },
})
