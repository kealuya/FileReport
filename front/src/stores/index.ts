import {defineStore} from 'pinia'
// useStore 可以是 useUser、useCart 之类的任何东西
// 第一个参数是应用程序中 store 的唯一 id
export const useUserStore = defineStore('user', {
    state: () => {
        return {
            phoneNumber: "",
            userName: "",
            token: "",// 后台生成token
        }
    },
    getters: {
        test(state) {
            return false
        },
    },
    actions: {
        loginSuccess(user: User) {
            this.token = user.token
            this.phoneNumber = user.phoneNumber
            this.userName = user.userName
        },

        logout() {
            this.$state = {
                phoneNumber: "",
                userName: "",
                token: "",
            }
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
