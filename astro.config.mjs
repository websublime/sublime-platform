import node from '@astrojs/node';
import { defineConfig } from 'astro/config';

// https://github.com/withastro/astro/tree/main/examples/ssr/src
// https://astro.build/config
export default defineConfig({
  adapter: node(),
  integrations: [
    {
      hooks: {
        'astro:config:setup': ({ injectScript }) => {
          const { API_URL = 'http://localhost', NODE_ENV = 'production' } = process.env;
          injectScript('page', `import { bootGlobals } from "@websublime/ws-globals";bootGlobals({apiUrl: "${API_URL}", env: "${NODE_ENV}"});`);
        }
      },
      name: '@sublime/globals'
    }
  ],
  vite: {
    optimizeDeps: {
      exclude: ['@websublime/ws-essential']
    }
  },
  output: 'server'
});
