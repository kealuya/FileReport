/// <reference types="vite/client" />

declare module "*.vue" {
    import {DefineComponent} from "vue";
    // eslint-disable-next-line @typescript-eslint/no-explicit-any, @typescript-eslint/ban-types
    const component: DefineComponent<{}, {}, any>;
    export default component;
}


declare interface MyFile {
    id: string
    updateDate: string
    createDate: string
    name: string
    updateUser: string
    updateContent: string
    version: string
    //
    isHandlerButtonShow?: boolean
    fileType: string
    isRelease: false
    isOwnerEdit: true
    owner: string
    ownerId: string

}

declare module 'vue-countup-v2';
