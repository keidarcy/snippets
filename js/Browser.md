- Event and CustomEvent

```
$0.dispatchEvent(new Event('newEvent'))
$0.dispatchEvent(new CustomEvent('newEvent', { detail: 'my detail' })

```
 
 - Go to top
 
 ```
 export const toTop = () => {
  document.body.scrollTop = 0 // For Safari
  document.documentElement.scrollTop = 0
}
 ```

- Webpack setting for ts with dom

```js
// webpack.config.js
module.exports = {
  mode: 'development',
  entry: './src/index.ts',
  output: {
    path: `${__dirname}/dist`,
    filename: 'bundle.js'
  },
  devtool: 'eval-source-map', // for source map
  module: {
    rules: [
      {
        test: /\.ts$/,
        use: 'ts-loader'
      }
    ]
  },
  resolve: {
    extensions: ['.ts', '.js']
  }
};

```

- Upload image

```
const fd = new FormData()
fd.append('image', file, file.name)
const res = await axios.post(apiPath, fd)
```
