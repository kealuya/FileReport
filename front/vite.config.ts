import {resolve} from 'path'
import {UserConfigExport, ConfigEnv, loadEnv} from "vite";
import vue from '@vitejs/plugin-vue'

import Components from 'unplugin-vue-components/vite'
import {ElementPlusResolver} from 'unplugin-vue-components/resolvers'

import Unocss from 'unocss/vite'
import {
    presetAttributify,
    presetIcons,
    presetUno,
    transformerDirectives,
    transformerVariantGroup,
} from 'unocss'
import * as path from "path";

// 当前执行node命令时文件夹的地址（工作目录）
const pathSrc = path.resolve(__dirname, 'src') ///Users/kealuya/mywork/my_git/FileReport/front/src

///Users/kealuya/mywork/my_git/FileReport/front
// 下面两个等价
const root: string = process.cwd();
const pathRoot = path.resolve(__dirname, '')

// https://vitejs.dev/config/
export default ({command, mode}: ConfigEnv): UserConfigExport => {
    /*
        serve
        development
     */
    console.log(command);// serve build
    console.log(mode);// 开发模式、生产模式

    const {
        VITE_PORT,
        VITE_PROXY_DOMAIN,// 一般都叫api，没用上
        VITE_PROXY_DOMAIN_REAL
    } = loadEnv(mode, root);

    return {
        build: {
            outDir: "../static/dist/",

        },
        resolve: {
            alias: {
                '@': resolve(process.cwd(), '/src'),
                '#': resolve(process.cwd(), '/types'),
                '~/': `${pathSrc}/`,
            }
        },
        server: {
            port: parseInt(VITE_PORT),
            open: false,
            proxy: {
                '/api': {
                    target: VITE_PROXY_DOMAIN_REAL,
                    changeOrigin: true,
                    ws: true,
                    rewrite: (path) => path.replace(new RegExp('^/api'), '')
                }
            }
        },
        css: {
            preprocessorOptions: {
                scss: {
                    additionalData: `@use "~/styles/element/index.scss" as *;`,
                },
            },
        },
        plugins: [
            vue(),
            Components({
                // allow auto load markdown components under `./src/components/`
                extensions: ['vue', 'md'],
                // allow auto import and register components used in markdown
                include: [/\.vue$/, /\.vue\?vue/, /\.md$/],
                resolvers: [
                    ElementPlusResolver({
                        importStyle: 'sass',
                    }),
                ],
                dts: 'src/components.d.ts',
            }),

            // https://github.com/antfu/unocss
            // see unocss.config.ts for config
            Unocss({
                presets: [
                    presetUno(),
                    presetAttributify(),
                    presetIcons({
                        scale: 1.2,
                        warn: true,
                    }),
                ],
                transformers: [
                    transformerDirectives(),
                    transformerVariantGroup(),
                ]
            }),
        ],
    }
}
