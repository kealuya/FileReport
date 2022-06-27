import path, {resolve} from 'path'
import {defineConfig} from 'vite'
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

const pathSrc = path.resolve(__dirname, 'src')
const pathRoot = path.resolve(__dirname, '')

// https://vitejs.dev/config/
export default defineConfig({


    build:{
        outDir:"../static/dist/",

    },
    resolve: {
        alias: {
            '@': resolve(process.cwd(), '/src'),
            '#': resolve(process.cwd(), '/types'),
            '~/': `${pathSrc}/`,
        }
    },
    server: {
        port: 3002,
        open: false,
        proxy: {
            '/api': {
                target: 'http://admin.xueyueob.cn/api',
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
})
