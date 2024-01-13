# 初始化

```shell
yarn create vite
yarn add  -D tailwindcss postcss autoprefixer
npx tailwindcss init -p
```

可以将 `postcss.config.js`的配置，写在 `vite.config.js`中:

```js
// postcss.config.js
export default {
  plugins: {
    tailwindcss: {},
    autoprefixer: {},
  },
}

// vite.config.js
import autoprefixer from "autoprefixer";
import tailwindcss from "tailwindcss";

export default defineConfig({
  plugins: [react()],
  css: {
    postcss: {
      plugins: [tailwindcss(), autoprefixer()],
    },
  },
});
```
