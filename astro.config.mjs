import { defineConfig } from 'astro/config';

// https://astro.build/config
export default defineConfig({
  integrations: [
    {
      name: "@sublime/globals",
      hooks: {
        'astro:config:setup': ({ injectScript }) => {
          const { NODE_ENV = 'production', API_URL = 'http://localhost' } = process.env;

          injectScript(
            'page', 
            `
            import { bootGlobals } from "@websublime/ws-globals";
            bootGlobals({apiUrl: ${API_URL}, env: ${NODE_ENV}});
            `
          );
        }
      }
    }
  ]
});
