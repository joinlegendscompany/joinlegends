import { defineConfig } from '@kubb/core'
import { pluginClient } from '@kubb/plugin-client'
import { pluginOas } from '@kubb/plugin-oas'
import { pluginReactQuery } from '@kubb/plugin-react-query'
import { pluginTs } from '@kubb/plugin-ts'
import dotenv from 'dotenv'

dotenv.config()
const swaggerUrl = process.env.API_SWAGGER_URL || 'http://localhost:3000/swagger.json'

export default defineConfig(() => {
    return {
        root: '.',
        input: {
            path: swaggerUrl,
        },
        output: {
            path: './src/http/generated',
            clean: true,
            extension: {
                '.ts': ''
            }
        },
        plugins: [
            pluginOas({
                validate: false,
                generators: [],
            }),
            pluginTs({
                output: {
                    path: 'models.ts',
                },
                group: {
                    type: 'tag'
                },
                enumType: 'asConst', // Melhor tipagem para enums
                // dateType: 'date', // Usar Date objects
            }),
            pluginClient({
                output: {
                    path: 'client',
                },
                group: {
                    type: 'tag'
                },
                importPath: '@/http/client',
                dataReturnType: 'full',
                // pathParamsType: 'object', // Consistência com React Query
            }),
            pluginReactQuery({
                output: {
                    path: 'query',
                },
                group: {
                    type: 'tag'
                },
                // paramsType: 'object',
                suspense: false,
                // infinite: {
                //     queryParam: 'page', // Para paginação infinita
                //     initialPageParam: 1,
                //     cursorParam: undefined,
                // },
                client: {
                    importPath: '@/http/client',
                    dataReturnType: 'full',
                },
                query: {
                    importPath: '@tanstack/react-query'
                }
            }),
            // pluginZod({
            //     output: { path: 'schemas' },
            //     group: { type: 'tag' },
            // }),
        ],
    }
})
