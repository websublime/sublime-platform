import { defineConfig } from 'astro/config';

// https://astro.build/config
export default defineConfig({
  integrations: [
    {
      name: "@sublime/globals",
      hooks: {
        'astro:config:setup': ({ injectScript }) => {
          injectScript(
            'page', 
            `
            import { bootGlobals } from "@websublime/ws-globals";
            bootGlobals({apiUrl: "localhost", env: "local"});
            `
          );
        }
      }
    }
  ]
});
